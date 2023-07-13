package types_test

/*
func TestPrevoteValidate(t *testing.T) {
	for _, tc := range []struct {
		name        string
		malleate    func(*types.PricesPrevote)
		expectedErr string
	}{
		{
			"invalid feeder",
			func(pv *types.PricesPrevote) {
				pv.Feeder = ""
			},
			"",
		},
		{
			"invalid hash",
			func(pv *types.PricesPrevote) {
				pv.Hash = ""
			},
			"pool id must not be 0",
		},
		{
			"invalid height",
			func(pv *types.PricesPrevote) {
				pv.Hash = ""
			},
			"pool id must not be 0",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pv := types.PricesPrevote(
				"0cf33fb528b388660c3a42c3f3250e983395290b75fef255050fb5bc48a6025f",
				"cre1zaavvzxez0elundtn32qnk9lkm8kmcszxclz6p",
				1,
			)
			tc.malleate(&pv)
			err := pv.Validate()
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}
*/
