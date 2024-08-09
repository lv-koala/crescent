package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/crescent-network/crescent/v5/x/oracle/types"
)

func TestVote_ParsePriceTuples(t *testing.T) {
	for _, tc := range []struct {
		name        string
		malleate    func(voteStr *string)
		expectedErr string
	}{
		{
			"valid",
			func(voteStr *string) {},
			"",
		},
		{
			"invalid format",
			func(voteStr *string) {
				*voteStr = "1:ETH:1810.4,4:1682899198:ETH:1808.5,1:1682899198:BTC:28110.45,4:1682899198:BTC:28108.25"
			},
			"invalid sender address: decoding bech32 failed: invalid separator index -1: invalid address",
		},
		{
			"empty exchange id",
			func(voteStr *string) {
				*voteStr = "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5,:1682899198:BTC:283310.45,4:1682899198:BTC:28108.25"
			},
			"invalid sender address: decoding bech32 failed: invalid separator index -1: invalid address",
		},
		{
			"invalid timestamp",
			func(voteStr *string) {
				*voteStr = "1:2682899198:ETH:1810.4,4:2682899198:ETH:1808.5,1:1682899198:BTC:28110.45,4:1682899198:BTC:28108.25"
			},
			"invalid sender address: decoding bech32 failed: invalid separator index -1: invalid address",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			voteStr := "1:1682899198:ETH:1810.4,4:1682899198:ETH:1808.5,1:1682899198:BTC:28110.45,4:1682899198:BTC:28108.25"
			tc.malleate(&voteStr)
			_, err := types.ParsePriceTuples(voteStr)
			fmt.Println(err)
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}
