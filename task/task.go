package task

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"rbnb-relay/pkg/config"
	"rbnb-relay/shared"

	stake_manager "rbnb-relay/bindings/StakeManager"
	"rbnb-relay/pkg/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	bncCmnTypes "github.com/stafiprotocol/go-sdk/common/types"
)

type Task struct {
	taskTicker     int64
	stop           chan struct{}
	bscRpcEndpoint string
	bcApiEndpoint  string
	keyPair        *secp256k1.Keypair
	gasLimit       *big.Int
	maxGasPrice    *big.Int

	stakeMangerAddress common.Address

	// need init on start()
	isDev          bool
	bscSideChainId string
	eraOffset      int64
	eraSeconds     int64

	bscClient            *shared.Client
	contractStakeManager *stake_manager.StakeManager
}

func NewTask(cfg *config.Config, keyPair *secp256k1.Keypair) (*Task, error) {
	gasLimitDeci, err := decimal.NewFromString(cfg.GasLimit)
	if err != nil {
		return nil, err
	}

	if gasLimitDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("gas limit is zero")
	}
	maxGasPriceDeci, err := decimal.NewFromString(cfg.MaxGasPrice)
	if err != nil {
		return nil, err
	}
	if maxGasPriceDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("max gas price is zero")
	}

	s := &Task{
		taskTicker:         15,
		stop:               make(chan struct{}),
		bscRpcEndpoint:     cfg.BscRpcEndpoint,
		bcApiEndpoint:      cfg.BcApiEndpoint,
		keyPair:            keyPair,
		gasLimit:           gasLimitDeci.BigInt(),
		maxGasPrice:        maxGasPriceDeci.BigInt(),
		stakeMangerAddress: common.HexToAddress(cfg.StakeMangerAddress),
	}

	return s, nil
}

func (task *Task) Start() error {
	client, err := shared.NewClient(task.bscRpcEndpoint, task.keyPair, task.gasLimit, task.maxGasPrice)
	if err != nil {
		return err
	}

	task.bscClient = client

	// set domain by chainId
	chainId, err := task.bscClient.Client().ChainID(context.Background())
	if err != nil {
		return err
	}
	switch chainId.Uint64() {
	case 56:
		task.isDev = false
		task.bscSideChainId = "bsc"
	case 97:
		task.isDev = true
		task.bscSideChainId = "chapel"
		bncCmnTypes.Network = bncCmnTypes.TestNetwork
	default:
		return fmt.Errorf("unsupport chainId: %d", chainId.Int64())
	}
	logrus.WithField("isDev", task.isDev).WithField("chainId", task.bscSideChainId).Infof("current chain")

	stakeManger, err := stake_manager.NewStakeManager(task.stakeMangerAddress, task.bscClient.Client())
	if err != nil {
		return err
	}
	bondedPools, err := stakeManger.GetBondedPools(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return err
	}
	if len(bondedPools) == 0 {
		return fmt.Errorf("no bonded pools")
	}
	task.contractStakeManager = stakeManger

	eraOffset, err := stakeManger.EraOffset(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return err
	}
	task.eraOffset = eraOffset.Int64()

	eraSeconds, err := stakeManger.EraSeconds(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return err
	}
	task.eraSeconds = eraSeconds.Int64()

	// check bc endpoint
	timestamp := time.Now().Unix() - 30*24*60*60
	reward, _, err := utils.NewRewardOnBcDu(task.bcApiEndpoint, task.bscSideChainId, bondedPools[0], timestamp, time.Now().Unix())
	if err != nil {
		return fmt.Errorf("query pool: %s reward from bc api: %s failed, err: %s", bondedPools[0].String(), task.bcApiEndpoint, err.Error())
	}
	if reward == 0 {
		return fmt.Errorf("query pool: %s reward from bc api: %s failed, reward is zero", bondedPools[0].String(), task.bcApiEndpoint)
	}

	utils.SafeGoWithRestart(task.Handler)
	return nil
}

func (task *Task) Stop() {
	close(task.stop)
}

func (task *Task) Handler() {
	logrus.Info("start voteHandler")
	ticker := time.NewTicker(time.Duration(task.taskTicker) * time.Second)
	defer ticker.Stop()

	for {

		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		case <-ticker.C:
			logrus.Debug("vote handler start -----------")
			err := task.voteHandler()
			if err != nil {
				logrus.Warnf("vote handler failed, err: %s", err.Error())
				continue
			}
			logrus.Debug("vote handler end -----------")
		}
	}
}

func (task *Task) waitTxOnChain(txHash common.Hash) (err error) {
	retry := 0
	txSuccess := false
	for {
		if retry > utils.RetryLimit {
			return fmt.Errorf("waitTxOnChain %s reach retry limit", txHash.String())
		}
		_, pending, err := task.bscClient.TransactionByHash(txHash)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"hash": txHash.String(),
				"err":  err.Error(),
			}).Warn("TransactionByHash")

			time.Sleep(utils.RetryInterval)
			retry++
			continue
		} else {
			if pending {
				logrus.WithFields(logrus.Fields{
					"hash":    txHash.String(),
					"pending": pending,
				}).Warn("TransactionByHash")

				time.Sleep(utils.RetryInterval)
				retry++
				continue
			} else {
				// check status
				var receipt *types.Receipt
				subRetry := 0
				for {
					if subRetry > utils.RetryLimit {
						return fmt.Errorf("TransactionReceipt %s reach retry limit", txHash.String())
					}

					receipt, err = task.bscClient.TransactionReceipt(txHash)
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"hash": txHash.String(),
							"err":  err.Error(),
						}).Warn("tx TransactionReceipt")

						time.Sleep(utils.RetryInterval)
						subRetry++
						continue
					}
					break
				}

				if receipt.Status == 1 { //success
					txSuccess = true
				}
				break
			}
		}
	}

	logrus.WithFields(logrus.Fields{
		"tx":         txHash.String(),
		"tx success": txSuccess,
	}).Info("tx already on chain")

	return nil
}
