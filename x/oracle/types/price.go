package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OraclePrices []OraclePrice

func NewPrice(usdPrice sdk.Dec, ticker string, blockNum int64) *OraclePrice {
	return &OraclePrice{
		Ticker:   ticker,
		BlockNum: blockNum,
		UsdPrice: usdPrice,
	}
}

func (p *OraclePrices) FilterByBlock(blockNum int64) *OraclePrices {
	prices := OraclePrices{}
	for _, price := range *p {
		if price.BlockNum == blockNum {
			prices = append(prices, price)
		}
	}
	return &prices
}

func (p *OraclePrices) FilterByTicker(ticker string) *OraclePrices {
	prices := OraclePrices{}
	for _, price := range *p {
		if price.Ticker == ticker {
			prices = append(prices, price)
		}
	}
	return &prices
}
