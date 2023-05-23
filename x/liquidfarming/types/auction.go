package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewRewardsAuction creates a new RewardsAuction.
func NewRewardsAuction(
	liquidFarmId, auctionId uint64, startTime, endTime time.Time,
	status AuctionStatus) RewardsAuction {
	return RewardsAuction{
		LiquidFarmId: liquidFarmId,
		Id:           auctionId,
		StartTime:    startTime,
		EndTime:      endTime,
		Status:       status,
		WinningBid:   nil,
		Rewards:      sdk.Coins{}, // the value is determined when the auction is finished
		Fees:         sdk.Coins{}, // the value is determined when the auction is finished
	}
}

// Validate validates RewardsAuction.
func (auction RewardsAuction) Validate() error {
	if auction.LiquidFarmId == 0 {
		return fmt.Errorf("liquid farm id must not be 0")
	}
	if auction.Id == 0 {
		return fmt.Errorf("id must not be 0")
	}
	if !auction.EndTime.After(auction.StartTime) {
		return fmt.Errorf("end time must be set after the start time")
	}
	if auction.Status != AuctionStatusStarted && auction.Status != AuctionStatusFinished && auction.Status != AuctionStatusSkipped {
		return fmt.Errorf("invalid auction status")
	}
	if err := auction.Rewards.Validate(); err != nil {
		return fmt.Errorf("invalid rewards: %w", err)
	}
	if err := auction.Fees.Validate(); err != nil {
		return fmt.Errorf("invalid fees: %w", err)
	}
	return nil
}

func (auction *RewardsAuction) SetStatus(status AuctionStatus) {
	auction.Status = status
}

func (auction *RewardsAuction) SetWinningBid(winningBid *Bid) {
	auction.WinningBid = winningBid
}

func (auction *RewardsAuction) SetRewards(rewards sdk.Coins) {
	auction.Rewards = rewards
}

func (auction *RewardsAuction) SetFees(fees sdk.Coins) {
	auction.Fees = fees
}

// NewBid creates a new Bid.
func NewBid(
	liquidFarmId, auctionId uint64, bidderAddr sdk.AccAddress, share sdk.Coin) Bid {
	return Bid{
		LiquidFarmId:     liquidFarmId,
		RewardsAuctionId: auctionId,
		Bidder:           bidderAddr.String(),
		Share:            share,
	}
}

// Validate validates Bid.
func (b Bid) Validate() error {
	if b.LiquidFarmId == 0 {
		return fmt.Errorf("liquid farm id must not be 0")
	}
	if b.RewardsAuctionId == 0 {
		return fmt.Errorf("rewards auction id must not be 0")
	}
	if _, err := sdk.AccAddressFromBech32(b.Bidder); err != nil {
		return fmt.Errorf("invalid bidder address %w", err)
	}
	if err := b.Share.Validate(); err != nil {
		return fmt.Errorf("invalid share: %w", err)
	}
	if !b.Share.IsPositive() {
		return fmt.Errorf("share amount must be positive")
	}
	return nil
}
