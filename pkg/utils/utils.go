package utils

import (
	"crypto/sha256"
	"time"

	bncCmnTypes "github.com/stafiprotocol/go-sdk/common/types"
)

var RetryLimit = 60
var RetryInterval = time.Second * 2

var (
	EraStateUninitialized      = uint8(0)
	EraStateNewEraExecuted     = uint8(1)
	EraStateOperateExecuted    = uint8(2)
	EraStateOperateAckExecuted = uint8(3)
)

func GetBcDelegationAddressFromBsc(addrBts []byte) bncCmnTypes.AccAddress {
	return GetStakeCAoB(addrBts, DelegateCAoBSalt)
}

func GetBcRewardAddressFromBsc(addrBts []byte) bncCmnTypes.AccAddress {
	return GetStakeCAoB(GetStakeCAoB(addrBts, DelegateCAoBSalt).Bytes(), RewardCAoBSalt)
}

func GetStakeCAoB(sourceAddr []byte, salt string) bncCmnTypes.AccAddress {
	saltBytes := []byte("Staking" + salt + "Address Anchor")
	return XOR(SumTruncated(saltBytes), sourceAddr)
}

var TruncatedSize = 20
var DelegateCAoBSalt string = "Delegate"
var RewardCAoBSalt string = "Reward"

func SumTruncated(bz []byte) []byte {
	hash := sha256.Sum256(bz)
	return hash[:TruncatedSize]
}

func XOR(a, b []byte) []byte {
	c := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] ^ b[i]
	}
	return c
}
