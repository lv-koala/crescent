package types

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PricePrevotes []PricePrevote
type PriceVotes []PriceVote
type MissCounters []MissCounter
type ValidPriceTuples []ValidPriceTuple
type PriceTuples []PriceTuple

func NewPricePrevote(
	hash AggregateVoteHash,
	feeder sdk.AccAddress,
	prevoteSubmitHeight int64,
) PricePrevote {
	return PricePrevote{
		Hash:   hash.String(),
		Feeder: feeder.String(),
		Height: prevoteSubmitHeight,
	}
}

func NewPriceVote(
	prevoteSubmitHeight int64,
	tupleCount uint64,
	priceTuples PriceTuples,
	salt string,
	feeder sdk.AccAddress,
) PriceVote {
	return PriceVote{
		Height:          prevoteSubmitHeight,
		ValidTupleCount: tupleCount,
		Salt:            salt,
		Feeder:          feeder.String(),
		PriceTuple:      priceTuples,
	}
}

// NewPricesTuple creates a PricesTuple instance
func NewPriceTuple(srcId uint64, ts uint64, ticker string, Prices sdk.Dec) PriceTuple {
	return PriceTuple{
		SourceExchangeId: srcId,
		Timestamp:        ts,
		Ticker:           ticker,
		UsdPrice:         Prices,
	}
}

// check if the PriceTuples format is valid
func ValidatePriceTuplesString(tuplesStr string) bool {
	if len(tuplesStr) == 0 {
		return false
	}

	tupleStrs := strings.Split(tuplesStr, ",")
	for _, tupleStr := range tupleStrs {
		//{src_id}:{ts}:{symbol}:{usd_price}
		splitStr := strings.Split(tupleStr, ":")
		if len(splitStr) != 4 {
			return false
		}

		decCoin, err := sdk.NewDecFromStr(splitStr[3])
		if err != nil {
			return false
		}
		if !decCoin.IsPositive() {
			return false
		}

		_, err = strconv.ParseUint(splitStr[0], 10, 64)
		if err != nil {
			return false
		}

		//TODO: check timestamp format
		_, err = strconv.ParseInt(splitStr[1], 10, 64)
		if err != nil {
			return false
		}
	}

	return true

}

// ParsePricesTuples PricesTuple parser
func ParsePriceTuples(tuplesStr string) (PriceTuples, error) {
	if len(tuplesStr) == 0 {
		return nil, nil
	}

	tupleStrs := strings.Split(tuplesStr, ",")
	fmt.Println("tuples:", len(tupleStrs))
	tuples := make(PriceTuples, len(tupleStrs)) //not a empty slice

	duplicateCheckMap := make(map[string]struct{})
	//{src_id}:{ts}:{symbol}:{usd_price}
	for i, tupleStr := range tupleStrs {
		splitStr := strings.Split(tupleStr, ":")
		fmt.Println(splitStr)
		if len(splitStr) != 4 {
			fmt.Println(tupleStr, "=>", splitStr)
			return nil, fmt.Errorf("invalid exchange rate %s", tupleStr)
		}

		decCoin, err := sdk.NewDecFromStr(splitStr[3])
		if err != nil {
			return nil, err
		}
		if !decCoin.IsPositive() {
			return nil, fmt.Errorf("invalid oracle price denom %s", decCoin.String())
		}

		srcId, err := strconv.ParseUint(splitStr[0], 10, 64)
		if err != nil {
			fmt.Println("invalid srcId", srcId)
			return nil, err
		}

		t, err := strconv.ParseUint(splitStr[1], 10, 64)
		if err != nil {
			fmt.Println("invalid timestamp", t)
			return nil, err
		}
		//TODO validate timestamp

		ticker := strings.ToUpper(splitStr[2])
		tuples[i] = PriceTuple{
			SourceExchangeId: srcId,
			Timestamp:        t,
			Ticker:           ticker,
			UsdPrice:         decCoin,
		}

		dupId := ticker + splitStr[0] // symbol + src_id

		if _, ok := duplicateCheckMap[dupId]; ok {
			return nil, errors.New("duplicate denom and source")
		}
		duplicateCheckMap[dupId] = struct{}{}
	}

	return tuples, nil
}

// Len implements sort.Interface
func (pb ValidPriceTuples) Len() int {
	return len(pb)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (pb ValidPriceTuples) Less(i, j int) bool {
	if pb[i].UsdPrice.LT(pb[j].UsdPrice) {
		return true
	}
	if pb[i].UsdPrice.Equal(pb[j].UsdPrice) {
		return (pb[i].SourceExchangeId < pb[j].SourceExchangeId)
	}
	return false
}

// Swap implements sort.Interface.
func (pb ValidPriceTuples) Swap(i, j int) {
	pb[i], pb[j] = pb[j], pb[i]
}

func (pb ValidPriceTuples) Median() (sdk.Dec, error) {
	if len(pb) == 0 {
		return sdk.ZeroDec(), nil
	}

	if !sort.IsSorted(pb) {
		sort.Sort(pb)
	}

	m := len(pb) / 2
	median := pb[m].UsdPrice
	if len(pb)%2 == 0 && m > 0 {
		median = median.Add(pb[m-1].UsdPrice).QuoInt64(2)
	}

	return median, nil

}

func (pb ValidPriceTuples) StandardDeviation() (sdk.Dec, error) {
	if len(pb) == 0 {
		return sdk.ZeroDec(), nil
	}

	median, err := pb.Median()
	if err != nil {
		return sdk.ZeroDec(), err
	}

	sum := sdk.ZeroDec()
	ll := int64(len(pb))
	for _, v := range pb {
		func() {
			defer func() {
				if e := recover(); e != nil {
					ll--
				}
			}()
			deviation := v.UsdPrice.Sub(median)
			sum = sum.Add(deviation.Mul(deviation))
		}()
	}

	variance := sum.QuoInt64(ll)

	standardDeviation, err := variance.ApproxSqrt()
	if err != nil {
		return sdk.ZeroDec(), err
	}

	return standardDeviation, nil

}

func (pb ValidPriceTuples) MAD() (sdk.Dec, error) {
	if len(pb) == 0 {
		return sdk.ZeroDec(), nil
	}

	median, err := pb.Median()
	if err != nil {
		return sdk.ZeroDec(), err
	}

	absDev := make([]sdk.Dec, 0, len(pb))
	for _, v := range pb {
		func() {
			d := v.UsdPrice.Sub(median)
			if d.IsNegative() {
				d = d.Neg()
			}
			absDev = append(absDev, d)
		}()
	}

	sort.Slice(absDev, func(i, j int) bool {
		return absDev[i].LT(absDev[j])
	})

	m := len(absDev) / 2
	mad := absDev[m]
	if len(absDev)%2 == 0 && m > 0 {
		mad = mad.Add(absDev[m-1]).QuoInt64(2)
	}

	return mad, nil

}
