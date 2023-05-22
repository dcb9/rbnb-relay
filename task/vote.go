package task

import (
	"context"
	"math/big"
	"rbnb-relay/pkg/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"rbnb-relay/bindings/StakePool"
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
	// case: currentEra==latestEra
	// no need deal
	if currentEra.Cmp(latestEra) == 0 {
		return nil
	}

	// case: currentEra > latestEra && !hasVoted
	// vote newEra
	willUseEra := new(big.Int).Add(latestEra, big.NewInt(1))

	newRewardList := make([]*big.Int, 0)
	lastRewardTimestampList := make([]*big.Int, 0)
	for _, pool := range bondedPools {
		latestRewardTimestampOnChain, err := t.contractStakeManager.LatestRewardTimestampOf(latestCallOpts, pool)
		if err != nil {
			return err
		}
		reward, lastRewardTimestamp, err := t.NewRewardOnBcAfterTimestamp(pool, latestRewardTimestampOnChain.Int64())
		if err != nil {
			return err
		}
		newReward := new(big.Int).Mul(big.NewInt(reward), big.NewInt(1e10)) //decimals 8 on bc, 18 on bsc

		newRewardList = append(newRewardList, newReward)
		lastRewardTimestampList = append(lastRewardTimestampList, big.NewInt(lastRewardTimestamp))
	}

	proposalId := NewEraProposalID(willUseEra, bondedPools, newRewardList, lastRewardTimestampList)
	hasVoted, err := t.contractStakeManager.HasVoted(latestCallOpts, proposalId, t.client.Opts().From)
	if err != nil {
		return err
	}
	// check and vote
	if !hasVoted {
		err = t.client.LockAndUpdateOpts(t.gasLimit, big.NewInt(0))
		if err != nil {
			return err
		}
		tx, err := t.contractStakeManager.NewEra(t.client.Opts(), willUseEra, bondedPools, newRewardList, lastRewardTimestampList)
		t.client.UnlockOpts()
		if err != nil {
			return err
		}

		err = t.waitTxOk(tx.Hash())
		if err != nil {
			return errors.Wrap(err, "wait Tx failed")
		}
	}

	//wait until newEraExecuted
	for {
		executeNewEraIter, err := t.contractStakeManager.FilterExecuteNewEra(latestFilterOpts, []*big.Int{willUseEra})
		if err != nil {
			return err
		}

		alreadyExecutedNewEra := false
		for executeNewEraIter.Next() {
			alreadyExecutedNewEra = true
			break
		}
		if !alreadyExecutedNewEra {
			logrus.Debug("wait newEraExecuted, will retry.")
			time.Sleep(utils.RetryInterval)
			continue
		}
		break
	}
	return nil
}

func (t *Task) checkAndSettle(currentEra, latestEra *big.Int, bondedPools []common.Address, latestCallOpts *bind.CallOpts, latestFilterOpts *bind.FilterOpts) error {
	for _, pool := range bondedPools {
		// check settle
		settleIter, err := t.contractStakeManager.FilterSettle(latestFilterOpts, []*big.Int{currentEra}, []common.Address{pool})
		if err != nil {
			return err
		}

		thisEraAlreadySettle := false
		for settleIter.Next() {
			thisEraAlreadySettle = true
		}

		if !thisEraAlreadySettle {
			contractStakePool, err := stake_pool.NewStakePool(pool, t.client.Client())
			if err != nil {
				return err
			}

			requestInfly, err := contractStakePool.GetRequestInFly(latestCallOpts)
			if err != nil {
				return err
			}

			if requestInfly[0].Cmp(big.NewInt(0)) == 0 &&
				requestInfly[1].Cmp(big.NewInt(0)) == 0 &&
				requestInfly[2].Cmp(big.NewInt(0)) == 0 {

				err = t.client.LockAndUpdateOpts(t.gasLimit, big.NewInt(0))
				if err != nil {
					return err
				}
				tx, err := t.contractStakeManager.Settle(t.client.Opts(), pool)
				t.client.UnlockOpts()
				if err != nil {
					return err
				}

				err = t.waitTxOk(tx.Hash())
				if err != nil {
					return errors.Wrap(err, "wait Tx failed")
				}
			}
		}

	}
	return nil
}

// uint256 _era,
// address[] calldata _poolAddressList,
// uint256[] calldata _undistributedRewardList,
// uint256[] calldata _latestRewardTimestampList
func NewEraProposalID(_era *big.Int, _poolAddressList []common.Address, _undistributedRewardList, _latestRewardTimestampList []*big.Int) [32]byte {
	return crypto.Keccak256Hash([]byte("newEra"), common.LeftPadBytes(_era.Bytes(), 32))
}
