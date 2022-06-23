package amm

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/crescent-network/crescent/types"
)

var (
	_ Pool = (*BasicPool)(nil)
	_ Pool = (*RangedPool)(nil)
)

// Pool is the interface of a pool.
type Pool interface {
	Balances() (rx, ry sdk.Int)
	SetBalances(rx, ry sdk.Int)
	PoolCoinSupply() sdk.Int
	Price() sdk.Dec
	IsDepleted() bool

	HighestBuyPrice() (sdk.Dec, bool)
	LowestSellPrice() (sdk.Dec, bool)
	BuyAmountOver(price sdk.Dec, inclusive bool) sdk.Int
	SellAmountUnder(price sdk.Dec, inclusive bool) sdk.Int
	BuyAmountTo(price sdk.Dec) sdk.Int
	SellAmountTo(price sdk.Dec) sdk.Int

	Clone() Pool
}

// BasicPool is the basic pool type.
type BasicPool struct {
	// rx and ry are the pool's reserve balance of each x/y coin.
	// In perspective of a pair, x coin is the quote coin and
	// y coin is the base coin.
	rx, ry sdk.Int
	// ps is the pool's pool coin supply.
	ps sdk.Int
}

// NewBasicPool returns a new BasicPool.
// It is OK to pass an empty sdk.Int to ps when ps is not going to be used.
func NewBasicPool(rx, ry, ps sdk.Int) *BasicPool {
	return &BasicPool{
		rx: rx,
		ry: ry,
		ps: ps,
	}
}

// Balances returns the balances of the pool.
func (pool *BasicPool) Balances() (rx, ry sdk.Int) {
	return pool.rx, pool.ry
}

func (pool *BasicPool) SetBalances(rx, ry sdk.Int) {
	pool.rx = rx
	pool.ry = ry
}

// PoolCoinSupply returns the pool coin supply.
func (pool *BasicPool) PoolCoinSupply() sdk.Int {
	return pool.ps
}

// Price returns the pool price.
func (pool *BasicPool) Price() sdk.Dec {
	if pool.rx.IsZero() || pool.ry.IsZero() {
		panic("pool price is not defined for a depleted pool")
	}
	return pool.rx.ToDec().Quo(pool.ry.ToDec())
}

// IsDepleted returns whether the pool is depleted or not.
func (pool *BasicPool) IsDepleted() bool {
	return pool.ps.IsZero() || pool.rx.IsZero() || pool.ry.IsZero()
}

// HighestBuyPrice returns the highest buy price of the pool.
func (pool *BasicPool) HighestBuyPrice() (price sdk.Dec, found bool) {
	// The highest buy price is actually a bit lower than pool price,
	// but it's not important for our matching logic.
	return pool.Price(), true
}

// LowestSellPrice returns the lowest sell price of the pool.
func (pool *BasicPool) LowestSellPrice() (price sdk.Dec, found bool) {
	// The lowest sell price is actually a bit higher than the pool price,
	// but it's not important for our matching logic.
	return pool.Price(), true
}

func (pool *BasicPool) BuyAmountOver(price sdk.Dec, _ bool) (amt sdk.Int) {
	if price.GTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	dx := pool.rx.ToDec().Sub(price.MulInt(pool.ry)).TruncateInt()
	if dx.IsZero() {
		return sdk.ZeroInt()
	}
	utils.SafeMath(func() {
		amt = pool.rx.ToDec().QuoTruncate(price).Sub(pool.ry.ToDec()).TruncateInt()
		if amt.GT(MaxCoinAmount) {
			amt = MaxCoinAmount
		}
	}, func() {
		amt = MaxCoinAmount
	})
	return
}

func (pool *BasicPool) SellAmountUnder(price sdk.Dec, _ bool) sdk.Int {
	if price.LTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	return pool.ry.ToDec().Sub(pool.rx.ToDec().QuoRoundUp(price)).TruncateInt()
}

func (pool *BasicPool) BuyAmountTo(price sdk.Dec) (amt sdk.Int) {
	if price.GTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	sqrtRx := utils.DecApproxSqrt(pool.rx.ToDec())
	sqrtRy := utils.DecApproxSqrt(pool.ry.ToDec())
	sqrtPrice := utils.DecApproxSqrt(price)
	dx := pool.rx.ToDec().Sub(sqrtPrice.Mul(sqrtRx.Mul(sqrtRy))) // dx = rx - sqrt(P * rx * ry)
	if dx.IsZero() {
		return sdk.ZeroInt()
	}
	utils.SafeMath(func() {
		amt = dx.QuoTruncate(price).TruncateInt() // dy = dx / P
		if amt.GT(MaxCoinAmount) {
			amt = MaxCoinAmount
		}
	}, func() {
		amt = MaxCoinAmount
	})
	return
}

func (pool *BasicPool) SellAmountTo(price sdk.Dec) sdk.Int {
	if price.LTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	sqrtRx := utils.DecApproxSqrt(pool.rx.ToDec())
	sqrtRy := utils.DecApproxSqrt(pool.ry.ToDec())
	sqrtPrice := utils.DecApproxSqrt(price)
	// dy = ry - sqrt(rx * ry / P)
	return pool.ry.ToDec().Sub(sqrtRx.Mul(sqrtRy).Quo(sqrtPrice)).TruncateInt()
}

func (pool *BasicPool) Clone() Pool {
	return NewBasicPool(pool.rx, pool.ry, pool.ps)
}

type RangedPool struct {
	rx, ry             sdk.Int
	ps                 sdk.Int
	minPrice, maxPrice sdk.Dec
	transX, transY     sdk.Dec
	xComp, yComp       sdk.Dec
}

// NewRangedPool returns a new RangedPool.
func NewRangedPool(rx, ry, ps sdk.Int, minPrice, maxPrice sdk.Dec) *RangedPool {
	transX, transY := DeriveTranslation(rx, ry, minPrice, maxPrice)
	return &RangedPool{
		rx:       rx,
		ry:       ry,
		ps:       ps,
		minPrice: minPrice,
		maxPrice: maxPrice,
		transX:   transX,
		transY:   transY,
		xComp:    rx.ToDec().Add(transX),
		yComp:    ry.ToDec().Add(transY),
	}
}

// CreateRangedPool creates new RangedPool from given inputs, while validating
// the inputs and using only needed amount of x/y coins(the rest should be refunded).
func CreateRangedPool(x, y sdk.Int, minPrice, maxPrice, initialPrice sdk.Dec) (pool *RangedPool, err error) {
	if !x.IsPositive() && !y.IsPositive() {
		return nil, fmt.Errorf("either x or y must be positive")
	}
	if err := ValidateRangedPoolParams(minPrice, maxPrice, initialPrice); err != nil {
		return nil, err
	}

	// P = initialPrice, M = minPrice, L = maxPrice
	var ax, ay sdk.Int
	switch {
	case initialPrice.Equal(minPrice): // single y asset pool
		ax = zeroInt
		ay = y
	case initialPrice.Equal(maxPrice): // single x asset pool
		ax = x
		ay = zeroInt
	default: // normal pool
		sqrt := utils.DecApproxSqrt
		xDec, yDec := x.ToDec(), y.ToDec()
		sqrtP := sqrt(initialPrice) // sqrt(P)
		sqrtM := sqrt(minPrice)     // sqrt(M)
		sqrtL := sqrt(maxPrice)     // sqrt(L)
		// Assume that we can accept all x
		ax = x
		// ay = {x / (sqrt(P)-sqrt(M))} / (1/sqrt(P) - 1/sqrt(L))
		ay = xDec.Quo(sqrtP.Sub(sqrtM)).Mul(inv(sqrtP).Sub(inv(sqrtL))).Ceil().TruncateInt()
		if ay.GT(y) {
			// Accept all y
			// ax = {y / (1/sqrt(P) - 1/sqrt(L))} * (sqrt(P) - sqrt(M))
			ax = yDec.Quo(inv(sqrtP).Sub(inv(sqrtL))).Mul(sqrtP.Sub(sqrtM)).Ceil().TruncateInt()
			ay = y
		}
	}
	return NewRangedPool(ax, ay, InitialPoolCoinSupply(ax, ay), minPrice, maxPrice), nil
}

func ValidateRangedPoolParams(minPrice, maxPrice, initialPrice sdk.Dec) error {
	if !initialPrice.IsPositive() {
		return fmt.Errorf("initial price must be positive: %s", initialPrice)
	}
	if minPrice.LT(MinRangedPoolPrice) {
		return fmt.Errorf("min price must be greater than %s", MinRangedPoolPrice)
	}
	if !maxPrice.IsPositive() {
		return fmt.Errorf("max price must be positive: %s", maxPrice)
	}
	if maxPrice.GT(MaxRangedPoolPrice) {
		return fmt.Errorf("max price must be less than %s", MaxRangedPoolPrice)
	}
	if !maxPrice.GT(minPrice) {
		return fmt.Errorf("max price must be greater than min price")
	}
	if maxPrice.Sub(minPrice).Quo(minPrice).LT(MinRangedPoolPriceGapRatio) {
		return fmt.Errorf("min price and max price are too close")
	}
	if initialPrice.LT(minPrice) {
		return fmt.Errorf("initial price must not be less than min price")
	}
	if initialPrice.GT(maxPrice) {
		return fmt.Errorf("initial price must not be greater thn max price")
	}
	return nil
}

// Balances returns the balances of the pool.
func (pool *RangedPool) Balances() (rx, ry sdk.Int) {
	return pool.rx, pool.ry
}

func (pool *RangedPool) SetBalances(rx, ry sdk.Int) {
	pool.rx = rx
	pool.ry = ry
	pool.xComp = pool.rx.ToDec().Add(pool.transX)
	pool.yComp = pool.ry.ToDec().Add(pool.transY)
}

// PoolCoinSupply returns the pool coin supply.
func (pool *RangedPool) PoolCoinSupply() sdk.Int {
	return pool.ps
}

func (pool *RangedPool) Translation() (transX, transY sdk.Dec) {
	return DeriveTranslation(pool.rx, pool.ry, pool.minPrice, pool.maxPrice)
}

// Price returns the pool price.
func (pool *RangedPool) Price() sdk.Dec {
	if pool.rx.IsZero() && pool.ry.IsZero() {
		panic("pool price is not defined for a depleted pool")
	}
	return pool.xComp.Quo(pool.yComp) // (rx + transX) / (ry + transY)
}

// IsDepleted returns whether the pool is depleted or not.
func (pool *RangedPool) IsDepleted() bool {
	return pool.ps.IsZero() || (pool.rx.IsZero() && pool.ry.IsZero())
}

// HighestBuyPrice returns the highest buy price of the pool.
func (pool *RangedPool) HighestBuyPrice() (price sdk.Dec, found bool) {
	// The highest buy price is actually a bit lower than pool price,
	// but it's not important for our matching logic.
	return pool.Price(), true
}

// LowestSellPrice returns the lowest sell price of the pool.
func (pool *RangedPool) LowestSellPrice() (price sdk.Dec, found bool) {
	// The lowest sell price is actually a bit higher than the pool price,
	// but it's not important for our matching logic.
	return pool.Price(), true
}

// BuyAmountOver returns the amount of buy orders for price greater than
// or equal to given price.
func (pool *RangedPool) BuyAmountOver(price sdk.Dec, _ bool) (amt sdk.Int) {
	if price.GTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	if price.LT(pool.minPrice) {
		price = pool.minPrice
	}
	// dx = (rx + transX) - P * (ry + transY)
	dx := pool.xComp.Sub(price.Mul(pool.yComp))
	if dx.IsZero() {
		return sdk.ZeroInt()
	} else if dx.GT(pool.rx.ToDec()) {
		dx = pool.rx.ToDec()
	}
	utils.SafeMath(func() {
		amt = dx.QuoTruncate(price).TruncateInt() // dy = dx / P
		if amt.GT(MaxCoinAmount) {
			amt = MaxCoinAmount
		}
	}, func() {
		amt = MaxCoinAmount
	})
	return
}

func (pool *RangedPool) SellAmountUnder(price sdk.Dec, _ bool) (amt sdk.Int) {
	if price.LTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	if price.GT(pool.maxPrice) {
		price = pool.maxPrice
	}
	// dy = (ry + transY) - (rx + transX) / P
	amt = pool.yComp.Sub(pool.xComp.QuoRoundUp(price)).TruncateInt()
	if amt.GT(pool.ry) {
		amt = pool.ry
	}
	return
}

func (pool *RangedPool) BuyAmountTo(price sdk.Dec) (amt sdk.Int) {
	if price.GTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	if price.LT(pool.minPrice) {
		price = pool.minPrice
	}
	sqrtXComp := utils.DecApproxSqrt(pool.xComp)
	sqrtYComp := utils.DecApproxSqrt(pool.yComp)
	sqrtPrice := utils.DecApproxSqrt(price)
	// dx = rx - (sqrt(P * (rx + transX) * (ry + transY)) - transX)
	dx := pool.rx.ToDec().Sub(sqrtPrice.Mul(sqrtXComp.Mul(sqrtYComp)).Sub(pool.transX))
	if dx.IsZero() {
		return sdk.ZeroInt()
	} else if dx.GT(pool.rx.ToDec()) {
		dx = pool.rx.ToDec()
	}
	utils.SafeMath(func() {
		amt = dx.QuoTruncate(price).TruncateInt() // dy = dx / P
		if amt.GT(MaxCoinAmount) {
			amt = MaxCoinAmount
		}
	}, func() {
		amt = MaxCoinAmount
	})
	return
}

func (pool *RangedPool) SellAmountTo(price sdk.Dec) (amt sdk.Int) {
	if price.LTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	if price.GT(pool.maxPrice) {
		price = pool.maxPrice
	}
	sqrtXComp := utils.DecApproxSqrt(pool.xComp)
	sqrtYComp := utils.DecApproxSqrt(pool.yComp)
	sqrtPrice := utils.DecApproxSqrt(price)
	// dy = ry - (sqrt((x + transX) * (y + transY) / P) - b)
	amt = pool.ry.ToDec().Sub(sqrtXComp.Mul(sqrtYComp).QuoRoundUp(sqrtPrice).Sub(pool.transY)).TruncateInt()
	if amt.GT(pool.ry) {
		amt = pool.ry
	}
	return
}

func (pool *RangedPool) Clone() Pool {
	return &RangedPool{
		rx:       pool.rx,
		ry:       pool.ry,
		ps:       pool.ps,
		transX:   pool.transX,
		transY:   pool.transY,
		xComp:    pool.xComp,
		yComp:    pool.yComp,
		minPrice: pool.minPrice,
		maxPrice: pool.maxPrice,
	}
}

// Deposit returns accepted x and y coin amount and minted pool coin amount
// when someone deposits x and y coins.
func Deposit(rx, ry, ps, x, y sdk.Int) (ax, ay, pc sdk.Int) {
	// Calculate accepted amount and minting amount.
	// Note that we take as many coins as possible(by ceiling numbers)
	// from depositor and mint as little coins as possible.

	utils.SafeMath(func() {
		rx, ry := rx.ToDec(), ry.ToDec()
		ps := ps.ToDec()

		// pc = floor(ps * min(x / rx, y / ry))
		pc = ps.MulTruncate(sdk.MinDec(
			x.ToDec().QuoTruncate(rx),
			y.ToDec().QuoTruncate(ry),
		)).TruncateInt()

		mintProportion := pc.ToDec().Quo(ps)             // pc / ps
		ax = rx.Mul(mintProportion).Ceil().TruncateInt() // ceil(rx * mintProportion)
		ay = ry.Mul(mintProportion).Ceil().TruncateInt() // ceil(ry * mintProportion)
	}, func() {
		ax, ay, pc = sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()
	})

	return
}

// Withdraw returns withdrawn x and y coin amount when someone withdraws
// pc pool coin.
// Withdraw also takes care of the fee rate.
func Withdraw(rx, ry, ps, pc sdk.Int, feeRate sdk.Dec) (x, y sdk.Int) {
	if pc.Equal(ps) {
		// Redeeming the last pool coin - give all remaining rx and ry.
		x = rx
		y = ry
		return
	}

	utils.SafeMath(func() {
		proportion := pc.ToDec().QuoTruncate(ps.ToDec())                             // pc / ps
		multiplier := sdk.OneDec().Sub(feeRate)                                      // 1 - feeRate
		x = rx.ToDec().MulTruncate(proportion).MulTruncate(multiplier).TruncateInt() // floor(rx * proportion * multiplier)
		y = ry.ToDec().MulTruncate(proportion).MulTruncate(multiplier).TruncateInt() // floor(ry * proportion * multiplier)
	}, func() {
		x, y = sdk.ZeroInt(), sdk.ZeroInt()
	})

	return
}

func DeriveTranslation(rx, ry sdk.Int, minPrice, maxPrice sdk.Dec) (transX, transY sdk.Dec) {
	sqrt := utils.DecApproxSqrt

	// M = minPrice, L = maxPrice
	rxDec, ryDec := rx.ToDec(), ry.ToDec()
	sqrtM := sqrt(minPrice)
	sqrtL := sqrt(maxPrice)

	var sqrtP sdk.Dec
	switch {
	case rxDec.Quo(ryDec).IsZero(): // y asset single pool
		sqrtP = sqrtM
	case ryDec.Quo(rxDec).IsZero(): // x asset single pool
		sqrtP = sqrtL
	default: // normal pool
		// sqrtXOverY = sqrt(rx/ry)
		sqrtXOverY := sqrt(rxDec.Quo(ryDec))
		// alpha = sqrt(M)/sqrt(rx/ry) - sqrt(rx/ry)/sqrt(L)
		alpha := sqrtM.Quo(sqrtXOverY).Sub(sqrtXOverY.Quo(sqrtL))
		// sqrtP = sqrt(P) = {(alpha + sqrt(alpha^2 + 4)) / 2} * sqrt(rx/ry)
		sqrtP = alpha.Add(sqrt(alpha.Power(2).Add(fourDec))).QuoInt64(2).Mul(sqrtXOverY)
	}

	// sqrtML = sqrt(M) * sqrt(L)
	sqrtML := sqrtM.Mul(sqrtL)
	if sqrtP.Equal(sqrtM) {
		// transX = ryl / (1/M - 1/sqrt(M*L))
		transX = ryDec.Quo(inv(minPrice).Sub(inv(sqrtML)))
	} else {
		// transX = rx / (sqrt(P/M) - 1)
		transX = rxDec.Quo(sqrtP.Quo(sqrtM).Sub(oneDec))
	}
	if sqrtP.Equal(sqrtL) {
		// transY = rx / (L - sqrt(M*L))
		transY = rxDec.Quo(maxPrice.Sub(sqrtML))
	} else {
		// transY = ry / (sqrt(L/P) - 1)
		transY = ryDec.Quo(sqrtL.Quo(sqrtP).Sub(oneDec))
	}

	return
}

func PoolOrders(pool Pool, orderer Orderer, lowestPrice, highestPrice sdk.Dec, tickPrec int) []Order {
	return append(
		PoolBuyOrders(pool, orderer, lowestPrice, highestPrice, tickPrec),
		PoolSellOrders(pool, orderer, lowestPrice, highestPrice, tickPrec)...)
}

func PoolBuyOrders(pool Pool, orderer Orderer, lowestPrice, highestPrice sdk.Dec, tickPrec int) (orders []Order) {
	poolPrice := pool.Price()
	if poolPrice.LTE(lowestPrice) {
		return nil
	}
	tmpPool := pool.Clone()
	placeOrder := func(price sdk.Dec, amt sdk.Int) {
		orders = append(orders, orderer.Order(Buy, price, amt))
		rx, ry := tmpPool.Balances()
		rx = rx.Sub(price.MulInt(amt).Ceil().TruncateInt()) // quote coin ceiling
		ry = ry.Add(amt)
		tmpPool.SetBalances(rx, ry)
	}
	if poolPrice.GT(highestPrice) {
		placeOrder(highestPrice, tmpPool.BuyAmountTo(highestPrice))
	}
	tick := PriceToDownTick(tmpPool.Price(), tickPrec)
	for tick.GTE(lowestPrice) {
		amt := tmpPool.BuyAmountOver(tick, true)
		if amt.LT(MinCoinAmount) {
			tick = DownTick(tick, tickPrec) // TODO: check if the tick is the lowest possible tick
			continue
		}
		placeOrder(tick, amt)
		rx, _ := tmpPool.Balances()
		if !rx.IsPositive() {
			break
		}
		tick = PriceToDownTick(tick.Mul(oneDec.Sub(poolOrderPriceGapRatio(poolPrice, tick))), tickPrec)
	}
	return orders
}

func PoolSellOrders(pool Pool, orderer Orderer, lowestPrice, highestPrice sdk.Dec, tickPrec int) (orders []Order) {
	poolPrice := pool.Price()
	if poolPrice.GTE(highestPrice) {
		return nil
	}
	tmpPool := pool.Clone()
	placeOrder := func(price sdk.Dec, amt sdk.Int) {
		orders = append(orders, orderer.Order(Sell, price, amt))
		rx, ry := tmpPool.Balances()
		rx = rx.Add(price.MulInt(amt).TruncateInt()) // quote coin truncation
		ry = ry.Sub(amt)
		tmpPool.SetBalances(rx, ry)
	}
	if poolPrice.LT(lowestPrice) {
		placeOrder(lowestPrice, tmpPool.SellAmountTo(lowestPrice))
	}
	tick := PriceToUpTick(tmpPool.Price(), tickPrec)
	for tick.LTE(highestPrice) {
		amt := tmpPool.SellAmountUnder(tick, true)
		if amt.LT(MinCoinAmount) || tick.MulInt(amt).TruncateInt().IsZero() {
			tick = UpTick(tick, tickPrec)
			continue
		}
		placeOrder(tick, amt)
		_, ry := tmpPool.Balances()
		if !ry.GT(MinCoinAmount) {
			break
		}
		tick = PriceToUpTick(tick.Mul(oneDec.Add(poolOrderPriceGapRatio(poolPrice, tick))), tickPrec)
	}
	return orders
}

// InitialPoolCoinSupply returns ideal initial pool coin minting amount.
func InitialPoolCoinSupply(x, y sdk.Int) sdk.Int {
	cx := len(x.BigInt().Text(10)) - 1 // characteristic of x
	cy := len(y.BigInt().Text(10)) - 1 // characteristic of y
	c := ((cx + 1) + (cy + 1) + 1) / 2 // ceil(((cx + 1) + (cy + 1)) / 2)
	res := big.NewInt(10)
	res.Exp(res, big.NewInt(int64(c)), nil) // 10^c
	return sdk.NewIntFromBigInt(res)
}
