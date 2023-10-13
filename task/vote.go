package task

import (
	"context"
	"fmt"
	"math/big"
	"time"

	stake_pool "rbnb-relay/bindings/StakePool"
	"rbnb-relay/pkg/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (t *Task) voteHandler() error {
	latestCallOpts := bind.CallOpts{
		Pending: false,
		From:    [20]byte{},
		Context: context.Background(),
	}

	currentEra, err := t.contractStakeManager.CurrentEra(&latestCallOpts)
	if err != nil {
		return err
	}
	latestEra, err := t.contractStakeManager.LatestEra(&latestCallOpts)
	if err != nil {
		return err
	}
	bondedPools, err := t.contractStakeManager.GetBondedPools(&latestCallOpts)
	if err != nil {
		return err
	}

	err = t.checkAndVoteNewEra(currentEra, latestEra, bondedPools, &latestCallOpts)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) checkAndVoteNewEra(currentEra, latestEra *big.Int, bondedPools []common.Address, latestCallOpts *bind.CallOpts) error {
	// case 0: currentEra==latestEra
	// no need deal
	if currentEra.Cmp(latestEra) == 0 {
		logrus.Debug("currentEra==latestEra no need deal")
		return nil
	}

	// case 2: currentEra > latestEra && !hasVoted
	// vote newEra
	willUseEra := new(big.Int).Add(latestEra, big.NewInt(1))

	newRewardList := make([]*big.Int, 0)
	lastRewardTimestampList := make([]*big.Int, 0)
	for _, pool := range bondedPools {
		contractStakePool, err := stake_pool.NewStakePool(pool, t.bscClient.Client())
		if err != nil {
			return err
		}
		requestInFly, err := contractStakePool.GetRequestInFly(latestCallOpts)
		if err != nil {
			return err
		}

		// case 1: currentEra>latestEra && requestInFly
		// no need deal
		if requestInFly[0].Cmp(big.NewInt(0)) != 0 || requestInFly[1].Cmp(big.NewInt(0)) != 0 || requestInFly[2].Cmp(big.NewInt(0)) != 0 {
			logrus.Warnf("pool %s, request is in fly %v, will wait", pool.String(), requestInFly)
			return nil
		}

		latestRewardTimestampOnChain, err := t.contractStakeManager.LatestRewardTimestampOf(latestCallOpts, pool)
		if err != nil {
			return err
		}

		if t.isDev && latestRewardTimestampOnChain.Sign() == 0 {
			latestRewardTimestampOnChain = big.NewInt(1682739051)
		}

		// api res: "rewardTime": "2023-05-28T00:00:02.000+00:00", so we add 100s here
		eraTimestamp := (willUseEra.Int64()+t.eraOffset)*t.eraSeconds + 100
		lastRewardTimestampBig := latestRewardTimestampOnChain

		eraReward := int64(0)
		retry := 0
		for {
			if retry > 600 {
				logrus.Warnf("no reward this era: %d, startTimestamp: %d endTimestamp: %d", willUseEra.Int64(), latestRewardTimestampOnChain.Int64(), eraTimestamp)

				//eraReward is zero
				eraReward = 0
				// not update latestRewardTimestamp if no reward
				break
			}

			newReward, maxRewardTimestamp, err := utils.NewRewardOnBcDu(t.bcApiEndpoint, t.bscSideChainId, pool, latestRewardTimestampOnChain.Int64(), eraTimestamp)
			if err != nil {
				return err
			}
			if newReward <= 0 {
				time.Sleep(3 * time.Second)
				retry++
				continue
			}

			// update eraReward and latestRewardTimestamp
			eraReward = newReward
			lastRewardTimestampBig = big.NewInt(maxRewardTimestamp)

			break
		}

		newRewardBig := new(big.Int).Mul(big.NewInt(eraReward), big.NewInt(1e10)) //decimals 8 on bc, 18 on bsc

		if latestRewardTimestampOnChain.Cmp(lastRewardTimestampBig) > 0 {
			return fmt.Errorf("pool %s, latestRewardTimestamp %d is big than lastRewardTimestamp %d", pool.String(), latestRewardTimestampOnChain.Int64(), lastRewardTimestampBig.Int64())
		}
		if newRewardBig.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("pool %s, newReward %d is negtive", pool.String(), newRewardBig.Int64())
		}

		newRewardList = append(newRewardList, newRewardBig)
		lastRewardTimestampList = append(lastRewardTimestampList, lastRewardTimestampBig)
	}

	//todo check ratelimit
	logrus.WithFields(logrus.Fields{
		"era":                 willUseEra.Int64(),
		"newRewardList":       newRewardList,
		"rewardTimestampList": lastRewardTimestampList,
	}).Info("newEra")

	proposalId := utils.NewEraProposalID(willUseEra, bondedPools, newRewardList, lastRewardTimestampList)
	hasVoted, err := t.contractStakeManager.HasVoted(latestCallOpts, proposalId, t.bscClient.Opts().From)
	if err != nil {
		return err
	}

	if !hasVoted {
		// check era again
		latestEra, err := t.contractStakeManager.LatestEra(latestCallOpts)
		if err != nil {
			return err
		}
		if willUseEra.Cmp(new(big.Int).Add(latestEra, big.NewInt(1))) != 0 {
			logrus.Debugf("willUseEra: %d not match latestEra: %d, no need deal", willUseEra.Int64(), latestEra.Int64())
			return nil
		}

		// send tx
		err = t.bscClient.LockAndUpdateOpts(t.gasLimit, big.NewInt(0))
		if err != nil {
			return err
		}
		tx, err := t.contractStakeManager.NewEra(t.bscClient.Opts(), bondedPools, newRewardList, lastRewardTimestampList)
		t.bscClient.UnlockOpts()
		if err != nil {
			return err
		}

		err = t.waitTxOnChain(tx.Hash())
		if err != nil {
			return errors.Wrap(err, "waitTxOnChain failed")
		}
	}

	//wait until newEra executed
	retry := 0
	for {
		// wait 2h
		if retry > 2*60*5 {
			return fmt.Errorf("wait newEra executed failed")
		}
		latestEra, err := t.contractStakeManager.LatestEra(latestCallOpts)
		if err != nil {
			logrus.Warnf("get latestEra failed: %s", err.Error())
			time.Sleep(12 * time.Second)
			retry++
			continue
		}

		if latestEra.Cmp(willUseEra) < 0 {
			logrus.Warnf("waiting newEra executed...")
			time.Sleep(12 * time.Second)
			retry++
			continue
		}
		logrus.Info("newEra already executed success")
		break
	}

	return nil
}
