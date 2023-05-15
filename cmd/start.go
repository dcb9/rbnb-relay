package cmd

import (
	"fmt"
	"rbnb-relay/pkg/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	flagLogLevel = "log_level"
	flagHomePath = "home"

	defaultConfigPath = ".rbnb_relay"
)

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Args:  cobra.ExactArgs(0),
		Short: "Start rBNB relay",
		RunE: func(cmd *cobra.Command, args []string) error {
			homePath, err := cmd.Flags().GetString(flagHomePath)
			if err != nil {
				return err
			}
			fmt.Printf("home path: %s\n", homePath)

			logLevelStr, err := cmd.Flags().GetString(flagLogLevel)
			if err != nil {
				return err
			}
			logLevel, err := logrus.ParseLevel(logLevelStr)
			if err != nil {
				return err
			}
			fmt.Printf("home path: %s\n", homePath)
			logrus.SetLevel(logLevel)

			cfg, err := config.Load(homePath)
			if err != nil {
				return err
			}

			logrus.Info("cfg %+v", cfg)
			return nil
		},
	}

	cmd.Flags().String(flagHomePath, defaultConfigPath, "Home path")
	cmd.Flags().String(flagLogLevel, logrus.InfoLevel.String(), "The logging level (trace|debug|info|warn|error|fatal|panic)")

	return cmd
}
