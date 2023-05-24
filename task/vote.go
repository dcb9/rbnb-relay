package task

import (
	"context"
	"fmt"
	"math/big"

	"rbnb-relay/bindings/StakePool"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (t *Task) voteHandler() error {
	latestCallOpts := bind.CallOpts{
		Pending: false,
		From:    [20]byte{},
		Context: context.Background(),
	}
	latestFilterOpts := bind.FilterOpts{Context: context.Background()}

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

	err = t.checkAndVoteNewEra(currentEra, latestEra, bondedPools, &latestCallOpts, &latestFilterOpts)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) checkAndVoteNewEra(currentEra, latestEra *big.Int, bondedPools []common.Address, latestCallOpts *bind.CallOpts, latestFilterOpts *bind.FilterOpts) error {
	// case 0: currentEra==latestEra
	// no need deal
	if currentEra.Cmp(latestEra) == 0 {
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

		// case 1: currentEra==latestEra
		// no need deal
		if requestInFly[0].Cmp(big.NewInt(0)) != 0 || requestInFly[1].Cmp(big.NewInt(0)) != 0 || requestInFly[2].Cmp(big.NewInt(0)) != 0 {
			logrus.Warnf("pool %s, request is in fly %v, will wait", pool.String(), requestInFly)
			return nil
		}

		latestRewardTimestampOnChain, err := t.contractStakeManager.LatestRewardTimestampOf(latestCallOpts, pool)
		if err != nil {
			return err
		}
		reward, lastRewardTimestamp, err := t.NewRewardOnBcAfterTimestamp(pool, latestRewardTimestampOnChain.Int64())
		if err != nil {
			return err
		}
		lastRewardTimestampBig := big.NewInt(lastRewardTimestamp)
		newRewardBig := new(big.Int).Mul(big.NewInt(reward), big.NewInt(1e10)) //decimals 8 on bc, 18 on bsc

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

	proposalId := NewEraProposalID(willUseEra, bondedPools, newRewardList, lastRewardTimestampList)
	hasVoted, err := t.contractStakeManager.HasVoted(latestCallOpts, proposalId, t.bscClient.Opts().From)
	if err != nil {
		return err
	}

	if !hasVoted {
		err = t.bscClient.LockAndUpdateOpts(t.gasLimit, big.NewInt(0))
		if err != nil {
			return err
		}
		tx, err := t.contractStakeManager.NewEra(t.bscClient.Opts(), willUseEra, bondedPools, newRewardList, lastRewardTimestampList)
		t.bscClient.UnlockOpts()
		if err != nil {
			return err
		}

		err = t.waitTxOk(tx.Hash())
		if err != nil {
			return errors.Wrap(err, "wait Tx failed")
		}
	}

	return nil
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
