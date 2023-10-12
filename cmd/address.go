package cmd

import (
	"fmt"
	"rbnb-relay/pkg/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	bncCmnTypes "github.com/stafiprotocol/go-sdk/common/types"
)

func addressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address",
		Short: "an utility to convert address between BSC chain and BC chain",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			chainId, err := cmd.Flags().GetInt("chainId")
			if err != nil {
				return err
			}
			fmt.Printf("chainId %d\n", chainId)
			if chainId == 97 {
				bncCmnTypes.Network = bncCmnTypes.TestNetwork
			}
			return nil
		},
	}
	cmd.PersistentFlags().Int("chainId", 97, "56 mainint, 97 testnet")

	bscAddrToBCCmd := &cobra.Command{
		Use:   "tobc",
		Args:  cobra.ExactArgs(1),
		Short: "an utility to convert address between BSC chain and BC chain",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := common.HexToAddress(args[0])
			fmt.Printf("addr on BSC: %s\n", addr)
			addrOnBC, err := bncCmnTypes.AccAddressFromHex(utils.RemovePrefix(addr.Hex()))
			if err != nil {
				return err
			}
			fmt.Printf("addr on BC: %s\n", addrOnBC)
			delegatorAddress := utils.GetBcDelegationAddressFromBsc(addr[:]).String()
			fmt.Printf("delegator addr on BC: %s\n", delegatorAddress)
			rewardAddress := utils.GetBcRewardAddressFromBsc(addr[:]).String()
			fmt.Printf("reward addr on BC: %s\n", rewardAddress)
			return nil
		},
	}

	cmd.AddCommand(
		bscAddrToBCCmd,
	)

	return cmd
}
