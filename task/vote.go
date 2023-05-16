package task

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

func (t *Task) voteHandler() error {
	latestCallOpts := bind.CallOpts{
		Pending: false,
		From:    [20]byte{},
		Context: context.Background(),
	}
	bondedPools, err := t.contractStakeManager.GetBondedPools(&latestCallOpts)
	if err != nil {
		return err
	}

	currentEra, err := t.contractStakeManager.CurrentEra(&latestCallOpts)
	if err != nil {
		return err
	}
	latestEra, err := t.contractStakeManager.LatestEra(&latestCallOpts)
	if err != nil {
		return err
	}
	if currentEra.Cmp(latestEra) <= 0 {
		return nil
	}

	ok, err := t.contractStakeManager.AllPoolEraStateIs(&latestCallOpts, uint8(3), false)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	willUseEra := new(big.Int).Add(latestEra, big.NewInt(1))

	rewardList := make([]*big.Int, 0)
	timestampList := make([]*big.Int, 0)

	timestampNow := uint64(time.Now().Unix())
	for _, pool := range bondedPools {
		latestRewardTimestamp, err := t.contractStakeManager.LatestRewardTimestampOf(&latestCallOpts, pool)
		if err != nil {
			return err
		}
		reward, err := t.getRewardDuring(pool, latestRewardTimestamp.Uint64(), timestampNow)
		if err != nil {
			return err
		}
		rewardList = append(rewardList, reward)
		timestampList = append(timestampList, big.NewInt(int64(timestampNow)))
	}

	err = t.client.LockAndUpdateOpts(t.gasLimit, big.NewInt(0))
	if err != nil {
		return err
	}
	defer func() {
		t.client.UnlockOpts()
	}()

	tx, err := t.contractStakeManager.NewEra(t.client.Opts(), willUseEra, bondedPools, rewardList, timestampList)
	if err != nil {
		return err
	}
	receipt, err := t.client.TransactionReceipt(tx.Hash())
	if err != nil {
		return err
	}
	logrus.Info("receipt", receipt)

	for _, pool := range bondedPools {

		pendingDelegate, err := t.contractStakeManager.PendingDelegateOf(&latestCallOpts, pool)
		if err != nil {
			return err
		}

		pendingUndelegate, err := t.contractStakeManager.PendingUndelegateOf(&latestCallOpts, pool)
		if err != nil {
			return err
		}
		tx, err := t.oprate(willUseEra, pool, pendingDelegate, pendingUndelegate)
		if err != nil {
			return err
		}
		receipt, err := t.client.TransactionReceipt(tx.Hash())
		if err != nil {
			return err
		}
		logrus.Info("receipt", receipt)
	}

	return nil
}

func (t *Task) getRewardDuring(pool common.Address, startTimestamp, endTimestamp uint64) (*big.Int, error) {
	return nil, nil
}

func (t *Task) oprate(era *big.Int, pool common.Address, pendingDelegate, pendingUndelegate *big.Int) (*types.Transaction, error) {
	return nil, nil
}
