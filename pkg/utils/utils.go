package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	bncCmnTypes "github.com/stafiprotocol/go-sdk/common/types"
)

var RetryLimit = 60
var RetryInterval = time.Second * 2

var (
	EraStateUninitialized      = uint8(0)
	EraStateNewEraExecuted     = uint8(1)
	EraStateOperateExecuted    = uint8(2)
	EraStateOperateAckExecuted = uint8(3)
)

func GetBcDelegationAddressFromBsc(addrBts []byte) bncCmnTypes.AccAddress {
	return GetStakeCAoB(addrBts, DelegateCAoBSalt)
}

func GetBcRewardAddressFromBsc(addrBts []byte) bncCmnTypes.AccAddress {
	return GetStakeCAoB(GetStakeCAoB(addrBts, DelegateCAoBSalt).Bytes(), RewardCAoBSalt)
}

func GetStakeCAoB(sourceAddr []byte, salt string) bncCmnTypes.AccAddress {
	saltBytes := []byte("Staking" + salt + "Address Anchor")
	return XOR(SumTruncated(saltBytes), sourceAddr)
}

var TruncatedSize = 20
var DelegateCAoBSalt string = "Delegate"
var RewardCAoBSalt string = "Reward"

func SumTruncated(bz []byte) []byte {
	hash := sha256.Sum256(bz)
	return hash[:TruncatedSize]
}

func XOR(a, b []byte) []byte {
	c := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] ^ b[i]
	}
	return c
}

// uint256 _era,
// address[] calldata _poolAddressList,
// uint256[] calldata _undistributedRewardList,
// uint256[] calldata _latestRewardTimestampList
func NewEraProposalID(_era *big.Int, _poolAddressList []common.Address, _undistributedRewardList, _latestRewardTimestampList []*big.Int) [32]byte {

	poolAddressBts := make([]byte, 0)
	for _, poolAddress := range _poolAddressList {
		poolAddressBts = append(poolAddressBts, common.LeftPadBytes(poolAddress.Bytes(), 32)...)
	}

	rewardBts := make([]byte, 0)
	for _, reward := range _undistributedRewardList {
		rewardBts = append(rewardBts, common.LeftPadBytes(reward.Bytes(), 32)...)
	}

	rewardTimestampBts := make([]byte, 0)
	for _, rewardTimestamp := range _latestRewardTimestampList {
		rewardTimestampBts = append(rewardTimestampBts, common.LeftPadBytes(rewardTimestamp.Bytes(), 32)...)
	}

	return crypto.Keccak256Hash([]byte("newEra"), common.LeftPadBytes(_era.Bytes(), 32), poolAddressBts, rewardBts, rewardTimestampBts)
}

// return reward decimals 8
func NewRewardOnBcAfterTimestamp(bcApiEndpoint, bscSideChainId string, pool common.Address, timestamp int64) (int64, int64, error) {
	rewardAddress := GetBcRewardAddressFromBsc(pool[:]).String()

	total, lastRewardTimestamp, err := RewardTotalTimesAndLastRewardTimestamp(bcApiEndpoint, bscSideChainId, rewardAddress)
	if err != nil {
		return 0, 0, err
	}
	logrus.Debug("RewardOnBcDuTimes", "total", total, "lastRewardTimestamp", lastRewardTimestamp, "delegator", rewardAddress)
	if timestamp >= lastRewardTimestamp {
		return 0, lastRewardTimestamp, nil
	}

	rewardSum, err := stakingRewardAfter(bcApiEndpoint, bscSideChainId, rewardAddress, total, timestamp)
	if err != nil {
		logrus.Warn("stakingReward error", "err", err)
		return 0, 0, err
	}

	return rewardSum, lastRewardTimestamp, nil

}

func RewardTotalTimesAndLastRewardTimestamp(bcApiEndpoint, bscSideChainId, delegator string) (int64, int64, error) {
	api := rewardApi(bcApiEndpoint, bscSideChainId, delegator, 1, 0)
	logrus.Debug("totalAndLastHeight rewardApi", "rewardApi", api)
	sr, err := getStakingReward(api)
	if err != nil {
		return 0, 0, err
	}

	if len(sr.RewardDetails) == 0 {
		return 0, 0, nil
	}

	rewardTime, err := time.Parse(time.RFC3339, sr.RewardDetails[0].RewardTime)
	if err != nil {
		return 0, 0, err
	}
	return sr.Total, rewardTime.Unix(), nil
}

func stakingRewardAfter(bcApiEndpoint, bscSideChainId, delegator string, total, timestamp int64) (int64, error) {
	logrus.Debug("stakingReward", "delegator", delegator, "total", total, "timestamp", timestamp)
	offset := int64(0)
	rewardSum := int64(0)

OUT:
	for i := total; i > 0; i -= 100 {
		api := rewardApi(bcApiEndpoint, bscSideChainId, delegator, 100, offset)

		sr, err := getStakingReward(api)
		if err != nil {
			return 0, errors.Wrap(err, "bnc.GetStakingReward")
		}

		for _, rd := range sr.RewardDetails {
			rewardTime, err := time.Parse(time.RFC3339, rd.RewardTime)
			if err != nil {
				return 0, err
			}

			if rewardTime.Unix() <= timestamp {
				break OUT
			}

			rewardSum += int64(rd.Reward * 1e8)
			logrus.Debug("stakingReward", "add", rd.Reward, "height", rd.Height)
		}

		offset += 100
	}

	return rewardSum, nil
}

func rewardApi(bcApiEndpoint, bscSideChainId, delegator string, limit, offset int64) string {
	return fmt.Sprintf("%s/v1/staking/chains/%s/delegators/%s/rewards?limit=%d&offset=%d", bcApiEndpoint, bscSideChainId, delegator, limit, offset)
}

type StakingReward struct {
	Total         int64          `json:"total"`
	RewardDetails []RewardDetail `json:"rewardDetails"`
}

type RewardDetail struct {
	ChainId    string  `json:"chainId"`
	Validator  string  `json:"validator"`
	ValName    string  `json:"valName"`
	Delegator  string  `json:"delegator"`
	Reward     float64 `json:"reward"`
	Height     int64   `json:"height"`
	RewardTime string  `json:"rewardTime"`
}

func getStakingReward(url string) (*StakingReward, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sr := new(StakingReward)
	if err := json.Unmarshal(body, sr); err != nil {
		return nil, err
	}

	return sr, nil
}
