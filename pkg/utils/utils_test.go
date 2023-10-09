package utils

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	bncCmnTypes "github.com/stafiprotocol/go-sdk/common/types"
)

func TestProposalId(t *testing.T) {
	proposal := NewEraProposalID(big.NewInt(456),
		[]common.Address{common.HexToAddress("0xd9145CCE52D386f254917e481eB44e9943F39138"), common.HexToAddress("0xd9145CCE52D386f254917e481eB44e9943F39138")},
		[]*big.Int{big.NewInt(1),
			big.NewInt(2)},
		[]*big.Int{big.NewInt(2),
			big.NewInt(3)})
	t.Log(hex.EncodeToString(proposal[:]))
}

func TestCAonB(t *testing.T) {
	bncCmnTypes.Network = bncCmnTypes.TestNetwork
	addr := common.HexToAddress("0x44f95eef755ed4fbdc19e3e8f617773d23e44a5b")
	expected := "tbnb1hhzxw7aak98aemqa42prxzkpqzasynfhhtgtfl"
	if got := GetBcDelegationAddressFromBsc(addr[:]).String(); got != expected {
		t.Fatalf("expected: %s got: %s", expected, got)
	}

	expected = "tbnb1u2cr5h32qq00v0k7zcpxpevyp4stjz57p0p4r2"
	if got := GetBcRewardAddressFromBsc(addr[:]).String(); got != expected {
		t.Fatalf("expected: %s got: %s", expected, got)
	}
}

func TestNewReward(t *testing.T) {
	bncCmnTypes.Network = bncCmnTypes.TestNetwork

	reward, max, err := NewRewardOnBcDu("https://testnet-api.binance.org", "chapel", common.HexToAddress("0x44f95eef755ed4fbdc19e3e8f617773d23e44a5b"), 0, 1682685051)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reward, max)

	reward, max, err = NewRewardOnBcDu("https://testnet-api.binance.org", "chapel", common.HexToAddress("0x44f95eef755ed4fbdc19e3e8f617773d23e44a5b"), 1685331051, 1685439051)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reward, max)

	reward, max, err = NewRewardOnBcDu("https://testnet-api.binance.org", "chapel", common.HexToAddress("0x44f95eef755ed4fbdc19e3e8f617773d23e44a5b"), 1685404800, 1685404800)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reward, max)

	reward, max, err = NewRewardOnBcDu("https://testnet-api.binance.org", "chapel", common.HexToAddress("0x44f95eef755ed4fbdc19e3e8f617773d23e44a5b"), 1685318300, 1685318400+100)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reward, max)
}
