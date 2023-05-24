package task

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"rbnb-relay/pkg/config"
	"rbnb-relay/shared"

	"rbnb-relay/bindings/StakeManager"
	"rbnb-relay/pkg/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
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
	default:
		return fmt.Errorf("unsupport chainId: %d", chainId.Int64())
	}

	stakeManger, err := stake_manager.NewStakeManager(task.stakeMangerAddress, task.bscClient.Client())
	if err != nil {
		return err
	}
	task.contractStakeManager = stakeManger

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
