package task

import (
	"context"
	"fmt"
	"math/big"
	"rbnb-relay/pkg/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	if currentEra.Cmp(latestEra) == 0 {
		allOperateAckExecuted, err := t.contractStakeManager.AllPoolEraStateIs(&latestCallOpts, utils.EraStateOperateAckExecuted, false)
		if err != nil {
			return err
		}
		// no need deal
		if allOperateAckExecuted {
			return nil
		}
	}

	bondedPools, err := t.contractStakeManager.GetBondedPools(&latestCallOpts)
	if err != nil {
		return err
	}

	willUseEra := new(big.Int).Add(latestEra, big.NewInt(1))

	// vote newEra
	if willUseEra.Cmp(latestEra) > 0 {

		newRewardList := make([]*big.Int, 0)
		timestampList := make([]*big.Int, 0)

		for _, pool := range bondedPools {
			latestRewardTimestampOnChain, err := t.contractStakeManager.LatestRewardTimestampOf(&latestCallOpts, pool)
			if err != nil {
				return err
			}
			reward, lastTimestamp, err := t.NewRewardOnBcAfterTimestamp(pool, latestRewardTimestampOnChain.Int64())
			if err != nil {
				return err
			}
			newReward := new(big.Int).Mul(big.NewInt(reward), big.NewInt(1e10))
			newRewardList = append(newRewardList, newReward)

			timestampList = append(timestampList, big.NewInt(lastTimestamp))
		}

		// check voted
		proposalId := NewEraProposalID(willUseEra, bondedPools, newRewardList, timestampList)
		hasVoted, err := t.contractStakeManager.HasVoted(&latestCallOpts, proposalId, t.client.Opts().From)
		if err != nil {
			return err
		}
		if !hasVoted {
			err = t.client.LockAndUpdateOpts(t.gasLimit, big.NewInt(0))
			if err != nil {
				return err
			}
			tx, err := t.contractStakeManager.NewEra(t.client.Opts(), willUseEra, bondedPools, newRewardList, timestampList)
			t.client.UnlockOpts()
			if err != nil {
				return err
			}

			receipt, err := t.client.TransactionReceipt(tx.Hash())
			if err != nil {
				return err
			}
			logrus.Info("receipt", receipt)
		}
		//todo wait newEraExecuted

	}

	executeNewEraIter, err := t.contractStakeManager.FilterExecuteNewEra(&bind.FilterOpts{}, []*big.Int{willUseEra})
	if err != nil {
		return err
	}
	executeNewEraHeight := big.NewInt(0)
	for executeNewEraIter.Next() {
		executeNewEraHeight = executeNewEraIter.Event.Block
	}

	newEraExecutedCallOpts := bind.CallOpts{
		Pending:     false,
		From:        [20]byte{},
		BlockNumber: executeNewEraHeight,
		Context:     context.Background(),
	}

	// vote operate and operate ack
	for _, pool := range bondedPools {

		poolInfo, err := t.contractStakeManager.PoolInfoOf(&latestCallOpts, pool)
		if err != nil {
			return err
		}
		if poolInfo.EraState == utils.EraStateNewEraExecuted {

			minDelegate := big.NewInt(0)

			pendingDelegate, err := t.contractStakeManager.PendingDelegateOf(&newEraExecutedCallOpts, pool)
			if err != nil {
				return err
			}

			pendingUndelegate, err := t.contractStakeManager.PendingUndelegateOf(&newEraExecutedCallOpts, pool)
			if err != nil {
				return err
			}

			switch {
			case pendingDelegate.Cmp(big.NewInt(0)) > 0 && pendingUndelegate.Cmp(big.NewInt(0)) == 0:
				if pendingDelegate.Cmp(minDelegate) < 0 {
					return fmt.Errorf("unknown pending state, pendingDelegate %s pendingUndelegate %s", pendingDelegate, pendingUndelegate)
				}
			case pendingDelegate.Cmp(big.NewInt(0)) == 0 && pendingUndelegate.Cmp(big.NewInt(0)) > 0:
			default:
				return fmt.Errorf("unknown pending state, pendingDelegate %s pendingUndelegate %s", pendingDelegate, pendingUndelegate)
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
