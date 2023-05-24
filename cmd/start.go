package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"rbnb-relay/pkg/config"
	"rbnb-relay/pkg/log"
	"rbnb-relay/pkg/utils"
	"rbnb-relay/task"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	"github.com/stafiprotocol/chainbridge/utils/keystore"
)

var (
	flagHome         = "home"
	flagBscEndpoint  = "bsc_endpoint"
	flagBcEndpoint   = "bc_endpoint"
	flagAccount      = "account"
	flagGasLimit     = "gas_limit"
	flagMaxGasPrice  = "max_gas_price"
	flagStakeManager = "stake_manager"
	flagLogLevel     = "log_level"

	defaultHomePath    = filepath.Join(os.Getenv("HOME"), ".stafi/rbnb")
	defaultBscEndpoint = "https://bsc-dataseed1.binance.org"
	defaultBcEndpoint  = "https://api.binance.org"
	defaultGasLimit    = "2000000"
	defaultMaxGasPrice = "20000000000"
	defaultStakeManger = "" //todo update address
	defaultLogLevel    = logrus.InfoLevel.String()
)

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Args:  cobra.ExactArgs(0),
		Short: "Start rBNB relay",
		RunE: func(cmd *cobra.Command, args []string) error {
			configHome, err := cmd.Flags().GetString(flagHome)
			if err != nil {
				return err
			}
			fmt.Printf("config home: %s\n", configHome)

			configBscEndpoint, err := cmd.Flags().GetString(flagBscEndpoint)
			if err != nil {
				return err
			}
			configAccount, err := cmd.Flags().GetString(flagAccount)
			if err != nil {
				return err
			}
			if !common.IsHexAddress(configAccount) {
				return fmt.Errorf("account not hex address: %s", configAccount)
			}

			configGasLimit, err := cmd.Flags().GetString(flagGasLimit)
			if err != nil {
				return err
			}

			configMaxGasPrice, err := cmd.Flags().GetString(flagMaxGasPrice)
			if err != nil {
				return err
			}
			configStakeManager, err := cmd.Flags().GetString(flagStakeManager)
			if err != nil {
				return err
			}
			if !common.IsHexAddress(configAccount) {
				return fmt.Errorf("stake manager not hex address: %s", configAccount)
			}
			// check log level
			logLevelStr, err := cmd.Flags().GetString(flagLogLevel)
			if err != nil {
				return err
			}
			logLevel, err := logrus.ParseLevel(logLevelStr)
			if err != nil {
				return err
			}
			fmt.Printf("log level: %s\n", logLevelStr)
			logrus.SetLevel(logLevel)

			logFilePath := filepath.Join(configHome, "log_data")
			keystorePath := filepath.Join(configHome, "keystore")

			cfg := config.Config{}
			cfg.BscRpcEndpoint = configBscEndpoint
			cfg.Account = configAccount
			cfg.GasLimit = configGasLimit
			cfg.MaxGasPrice = configMaxGasPrice
			cfg.StakeMangerAddress = configStakeManager

			cfg.LogFilePath = logFilePath
			cfg.KeystorePath = keystorePath

			err = log.InitLogFile(cfg.LogFilePath)
			if err != nil {
				return err
			}
			logrus.Infof("cfg %+v", cfg)

			ctx := utils.ShutdownListener()
			kpI, err := keystore.KeypairFromAddress(cfg.Account, keystore.EthChain, cfg.KeystorePath, false)
			if err != nil {
				return err
			}
			kp, ok := kpI.(*secp256k1.Keypair)
			if !ok {
				return fmt.Errorf("keypair err")
			}

			logrus.Info("voter task starting...")
			t, err := task.NewTask(&cfg, kp)
			if err != nil {
				return err
			}
			err = t.Start()
			if err != nil {
				logrus.Errorf("task start err: %s", err)
				return err
			}
			defer func() {
				logrus.Infof("shutting down task ...")
				t.Stop()
			}()

			<-ctx.Done()
			return nil

		},
	}

	cmd.Flags().String(flagHome, defaultHomePath, "Home path")
	cmd.Flags().String(flagBscEndpoint, defaultBscEndpoint, "Bsc rpc endpoint")
	cmd.Flags().String(flagBcEndpoint, defaultBcEndpoint, "Bc api endpoint")
	cmd.Flags().String(flagAccount, "", "Account hex string address")
	cmd.Flags().String(flagGasLimit, defaultGasLimit, "Gas limit")
	cmd.Flags().String(flagMaxGasPrice, defaultMaxGasPrice, "Max gas price")
	cmd.Flags().String(flagStakeManager, defaultStakeManger, "Stake manager contract address")
	cmd.Flags().String(flagLogLevel, defaultLogLevel, "The logging level (trace|debug|info|warn|error|fatal|panic)")

	return cmd
}
