package utils_test

import (
	"encoding/hex"
	"math/big"
	"rbnb-relay/pkg/utils"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	bncCmnTypes "github.com/stafiprotocol/go-sdk/common/types"
)

func TestProposalId(t *testing.T) {
	proposal := utils.NewEraProposalID(big.NewInt(456),
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
	expected := "tbnb1f6quqaz9e9f3v2cg85p83ss2cjmpzu7jmu5u4p"
	if got := utils.GetBcDelegationAddressFromBsc(addr[:]).String(); got != expected {
		t.Fatalf("expected: %s got: %s", expected, got)
	}

	expected = "tbnb1vmu2gwecm2045uccq5dtndtuptjkpykpfecrcj"
	if got := utils.GetBcRewardAddressFromBsc(addr[:]).String(); got != expected {
		t.Fatalf("expected: %s got: %s", expected, got)
	}
}

func TestValidatorAddrOnB(t *testing.T) {
	{
		valHexAddr := "0xa319b1e51f9c616ea25891805bc2f2ad15d50432"
		addr := common.HexToAddress(valHexAddr)
		valAddr := bncCmnTypes.ValAddress(addr[:])
		expect := "bva15vvmregln3skagjcjxq9hshj452a2ppjkhperr"
		if expect != valAddr.String() {
			t.Fatalf("val hex addr: %s expected %s, got %s", valHexAddr, expect, valAddr.String())
		}
	}
	{
		valHexAddr := "0xa6d02bf4b1990fe3db8c97de198372221d8c37ad"
		addr := common.HexToAddress(valHexAddr)
		valAddr := bncCmnTypes.ValAddress(addr[:])
		expect := "bva15mgzha93ny878kuvjl0pnqmjygwccdad08uecu"
		if expect != valAddr.String() {
			t.Fatalf("val hex addr: %s expected %s, got %s", valHexAddr, expect, valAddr.String())
		}
	}
}

func TestDecodeValidatorAddr(t *testing.T) {
	// get validator address on BSC from BC
	validatorAddr := "bva1pnww8kx30sz4xfcqvn8wjhrn796nf4dq77hcpa"
	addr, err := bncCmnTypes.ValAddressFromBech32(validatorAddr)
	if err != nil {
		t.Fatalf("err: %s", err.Error())
	}

	bscAddress := common.Address(addr).Hex()
	expect := "0x0cDcE3d8D17c0553270064cEe95C73F17534d5A0"
	if expect != bscAddress {
		t.Fatalf("decode validator address: %s error: %s != %s", validatorAddr, expect, bscAddress)
	}

	validatorAddr = "bva1wlxpuycavhvjzq27l6qzv07uf6cymgj7wcrf74"
	addr, err = bncCmnTypes.ValAddressFromBech32(validatorAddr)
	if err != nil {
		t.Fatalf("err: %s", err.Error())
	}

	bscAddress = common.Address(addr).Hex()
	expect = "0x77Cc1e131d65d921015EfE80263Fdc4eb04dA25e"
	if expect != bscAddress {
		t.Fatalf("decode validator address: %s error: %s != %s", validatorAddr, expect, bscAddress)
	}

	validatorAddr = "bva18wgcs0k0glcmaxreweyeydu9mudtsftcxpkt4n"
	addr, err = bncCmnTypes.ValAddressFromBech32(validatorAddr)
	if err != nil {
		t.Fatalf("err: %s", err.Error())
	}

	bscAddress = common.Address(addr).Hex()
	expect = "0x3B91883ECf47F1Be98797649923785dF1ab82578"
	if expect != bscAddress {
		t.Fatalf("decode validator address: %s error: %s != %s", validatorAddr, expect, bscAddress)
	}
}

func TestNewReward(t *testing.T) {
	bncCmnTypes.Network = bncCmnTypes.TestNetwork

	reward, max, err := utils.NewRewardOnBcDu("https://testnet-api.binance.org", "chapel", common.HexToAddress("0x44f95eef755ed4fbdc19e3e8f617773d23e44a5b"), 0, 1682685051)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reward, max)

	reward, max, err = utils.NewRewardOnBcDu("https://testnet-api.binance.org", "chapel", common.HexToAddress("0x44f95eef755ed4fbdc19e3e8f617773d23e44a5b"), 1685331051, 1685439051)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reward, max)

	reward, max, err = utils.NewRewardOnBcDu("https://testnet-api.binance.org", "chapel", common.HexToAddress("0x44f95eef755ed4fbdc19e3e8f617773d23e44a5b"), 1685404800, 1685404800)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reward, max)

	reward, max, err = utils.NewRewardOnBcDu("https://testnet-api.binance.org", "chapel", common.HexToAddress("0x44f95eef755ed4fbdc19e3e8f617773d23e44a5b"), 1685318300, 1685318400+100)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reward, max)
}
