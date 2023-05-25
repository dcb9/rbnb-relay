package utils

import (
	"crypto/sha256"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
