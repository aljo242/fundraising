package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/fundraising module sentinel errors
var (
	ErrInvalidAuctionType          = sdkerrors.Register(ModuleName, 2, "invalid auction type")
	ErrInvalidStartPrice           = sdkerrors.Register(ModuleName, 3, "invalid start price")
	ErrInvalidVestingSchedules     = sdkerrors.Register(ModuleName, 4, "invalid vesting schedules")
	ErrInvalidAuctionStatus        = sdkerrors.Register(ModuleName, 5, "invalid auction status")
	ErrInvalidPayingCoinDenom      = sdkerrors.Register(ModuleName, 6, "invalid paying coin denom")
	ErrEmptyAllowedBidders         = sdkerrors.Register(ModuleName, 7, "empty bidders")
	ErrInvalidMaxBidAmount         = sdkerrors.Register(ModuleName, 8, "invalid maximum bid amount")
	ErrOverMaxBidAmountLimit       = sdkerrors.Register(ModuleName, 9, "over maximum bid amount limit")
	ErrInsufficientRemainingAmount = sdkerrors.Register(ModuleName, 10, "insufficient remaining amount")
	ErrNotAllowedBidder            = sdkerrors.Register(ModuleName, 11, "not allowed bidder")
)
