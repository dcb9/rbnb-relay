package task

import (
	"context"
	"fmt"
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

	currentEra, err := t.contractStakeManager.CurrentEra(&latestCallOpts)
	if err != nil {
		return err
	}
	latestEra, err := t.contractStakeManager.LatestEra(&latestCallOpts)
	if err != nil {
		return err
	}

	// case: currentEra==latestEra && allPoolIsOperateAckExecuted
	// no need deal
	if currentEra.Cmp(latestEra) == 0 {
		allPoolIsOperateAckExecuted, err := t.contractStakeManager.AllPoolEraStateIs(&latestCallOpts, currentEra, utils.EraStateOperateAckExecuted, false)
		if err != nil {
			return err
		}
		if allPoolIsOperateAckExecuted {
			return nil
		}
	}

	bondedPools, err := t.contractStakeManager.GetBondedPools(&latestCallOpts)
	if err != nil {
		return err
	}

	willUseEra := new(big.Int).Add(latestEra, big.NewInt(1))

	// case: willUseEra > latestEra && !hasVoted
	// vote newEra
	if willUseEra.Cmp(latestEra) > 0 {
		newRewardList := make([]*big.Int, 0)
		lastRewardTimestampList := make([]*big.Int, 0)

		for _, pool := range bondedPools {
			latestRewardTimestampOnChain, err := t.contractStakeManager.LatestRewardTimestampOf(&latestCallOpts, pool)
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

		// check voted
		proposalId := NewEraProposalID(willUseEra, bondedPools, newRewardList, lastRewardTimestampList)
		hasVoted, err := t.contractStakeManager.HasVoted(&latestCallOpts, proposalId, t.client.Opts().From)
		if err != nil {
			return err
		}
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
	}

	newEraExecutedCallOpts := bind.CallOpts{
		Pending: false,
		From:    [20]byte{},
		Context: context.Background(),
	}

	//wait newEraExecuted
	for {
		executeNewEraIter, err := t.contractStakeManager.FilterExecuteNewEra(&bind.FilterOpts{}, []*big.Int{willUseEra})
		if err != nil {
			return err
		}
		executeNewEraHeight := big.NewInt(0)
		for executeNewEraIter.Next() {
			executeNewEraHeight = executeNewEraIter.Event.Block
			break
		}
		if executeNewEraHeight.Cmp(big.NewInt(0)) == 0 {
			logrus.Debug("wait newEraExecuted, will retry.")
			time.Sleep(utils.RetryInterval)
			continue
		}

		newEraExecutedCallOpts.BlockNumber = executeNewEraHeight
		break
	}

	// vote operate and operate ack
	for _, pool := range bondedPools {

		poolInfo, err := t.contractStakeManager.PoolInfoOf(&latestCallOpts, pool)
		if err != nil {
			return err
		}
		// case: poolEra==WillUseEra && eraState==EraStateNewEraExecuted
		// delegate or undelegate
		if poolInfo.Era == willUseEra && poolInfo.EraState == utils.EraStateNewEraExecuted {
			contractStakePool, err := stake_pool.NewStakePool(pool, t.client.Client())
			if err != nil {
				return err
			}
			minDelegate, err := contractStakePool.GetMinDelegation(&latestCallOpts)
			if err != nil {
				return err
			}

			pendingDelegate, err := t.contractStakeManager.PendingDelegateOf(&newEraExecutedCallOpts, pool)
			if err != nil {
				return err
			}

			pendingUndelegate, err := t.contractStakeManager.PendingUndelegateOf(&newEraExecutedCallOpts, pool)
			if err != nil {
				return err
			}

			switch {
			case pendingDelegate.Cmp(minDelegate) >= 0 && pendingUndelegate.Cmp(big.NewInt(0)) == 0:
				t.contractStakeManager.GetValidatorsOf()

			case pendingDelegate.Cmp(big.NewInt(0)) == 0 && pendingUndelegate.Cmp(big.NewInt(0)) > 0:

			default:
				return fmt.Errorf("unknown pending value state, pendingDelegate %s pendingUndelegate %s", pendingDelegate, pendingUndelegate)
			}

		}
	}

	return nil
}

func (t *Task) waitUntilNewEraExecuted(callOpts *bind.CallOpts) error {

	ok, err := t.contractStakeManager.AllPoolEraStateIs(callOpts, utils.EraStateNewEraExecuted, false)
	if err != nil {
		return err
	}
	if !ok {
		return nil
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
