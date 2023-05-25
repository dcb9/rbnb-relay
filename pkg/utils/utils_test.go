package utils_test

import (
	"encoding/hex"
	"math/big"
	"rbnb-relay/pkg/utils"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
