package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rbnb-relay/pkg/utils"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// return reward decimals 8
func (c *Task) NewRewardOnBcAfterTimestamp(pool common.Address, timestamp int64) (int64, int64, error) {
	rewardAddress := utils.GetBcRewardAddressFromBsc(pool[:]).String()

	total, lastRewardTimestamp, err := c.RewardTotalTimesAndLastRewardTimestamp(rewardAddress)
	if err != nil {
		return 0, 0, err
	}
	logrus.Debug("RewardOnBcDuTimes", "total", total, "lastRewardTimestamp", lastRewardTimestamp, "delegator", rewardAddress)
	if timestamp >= lastRewardTimestamp {
		return 0, 0, nil
	}

	rewardSum, err := c.stakingRewardAfter(rewardAddress, total, timestamp)
	if err != nil {
		logrus.Warn("stakingReward error", "err", err)
		return 0, 0, err
	}

	return rewardSum, lastRewardTimestamp, nil

}

func (c *Task) RewardTotalTimesAndLastRewardTimestamp(delegator string) (int64, int64, error) {
	api := c.rewardApi(delegator, 1, 0)
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

func (c *Task) stakingRewardAfter(delegator string, total, timestamp int64) (int64, error) {
	logrus.Debug("stakingReward", "delegator", delegator, "total", total, "timestamp", timestamp)
	offset := int64(0)
	rewardSum := int64(0)

OUT:
	for i := total; i > 0; i -= 100 {
		api := c.rewardApi(delegator, 100, offset)

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

func (c *Task) rewardApi(delegator string, limit, offset int64) string {
	return fmt.Sprintf("%s/v1/staking/chains/%s/delegators/%s/rewards?limit=%d&offset=%d", c.bcApiEndpoint, c.bscSideChainId, delegator, limit, offset)
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
