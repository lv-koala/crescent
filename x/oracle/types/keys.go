package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName is the name of the oracle module
	ModuleName = "oracle"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	TStoreKey = "transient_oracle"

	// RouterKey is the message route for oracle module
	RouterKey = ModuleName

	// QuerierRoute is the query router key for the oracle module
	QuerierRoute = ModuleName
)

// KVStore key prefixes
var (
	//KeyPrefix   = []byte{0xe0}
	PrevoteKeyPrefix     = []byte{0x81} //0X81 + addr
	VoteKeyPrefix        = []byte{0x82} //0x82 + addr
	MissCounterKeyPrefix = []byte{0x83}
	PriceKeyPrefix       = []byte{0x84}
	ValidVotePrefix      = []byte{0x85} // 0x85 + symbol
	LsvFeederPrefix      = []byte{0x86} // 0x86 + lsv address
	//LatestRewardKeyPrefix = []byte{0xe5}
)

func GetPrevoteKey(addr sdk.AccAddress) []byte {
	return append(PrevoteKeyPrefix, address.MustLengthPrefix(addr)...)
}

func GetPrevoteKeyWithHeight(addr sdk.AccAddress, height uint64) []byte {
	return append(append(PrevoteKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(height)...)
}

func GetMissCounterKey(addr sdk.AccAddress) []byte {
	return append(MissCounterKeyPrefix, address.MustLengthPrefix(addr)...)

}

func GetVoteKey(addr sdk.AccAddress, height uint64) []byte {
	return append(append(VoteKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(height)...)
}

func GetPriceKey(symbol string) []byte {
	return append(PriceKeyPrefix, []byte(symbol)...)
}

func GetLsvFeederKey(lsvAddr string) []byte {
	return append(LsvFeederPrefix, []byte(lsvAddr)...)
}

func GetValidVoteKey(symbol string) []byte {
	return append(ValidVotePrefix, LengthPrefixString(symbol)...)
}

// LengthPrefixString returns length-prefixed bytes representation
// of a string.
func LengthPrefixString(s string) []byte {
	bz := []byte(s)
	bzLen := len(bz)
	if bzLen == 0 {
		return bz
	}
	return append([]byte{byte(bzLen)}, bz...)
}
