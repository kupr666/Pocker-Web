package main

import (
	"errors"
	"math"

	tablev1 "github.com/kupr666/Pocker-Web/gen/go/table/v1"
	tablesv1 "github.com/kupr666/Pocker-Web/gen/go/tables/v1"
)

const (
	minSeats uint32 = 2
	maxSeats uint32 = 9

	minStartingStackBB uint64 = 20
	maxStartingStackBB uint64 = 200

	minTimeout     uint32 = 15
	defaultTimeout uint32 = 30
	maxTimeout     uint32 = 60
)

func prepareRules(input *tablesv1.TableRulesInput) (*tablev1.TableRules, error) {

	// values by default
	rules := &tablev1.TableRules{
		MaxSeats:           6,
		StartingStack:      1000,
		SmallBlind:         5,
		BigBlind:           10,
		TurnTimeoutSeconds: defaultTimeout,
		ReentryMode:        tablev1.ReentryMode_REENTRY_MODE_PRESERVE_STACK,
		RebuyMode:          tablev1.RebuyMode_REBUY_MODE_PLAYER_REQUEST,
		RebuyLimit:         3,
	}

	if input == nil {
		return rules, nil
	}

	// validation for non-optional fields
	switch mode := input.GetReentryMode(); mode {
	case tablev1.ReentryMode_REENTRY_MODE_UNSPECIFIED:

	case tablev1.ReentryMode_REENTRY_MODE_PRESERVE_STACK,
		tablev1.ReentryMode_REENTRY_MODE_RESTORE_STARTING_STACK:
		rules.ReentryMode = mode

	default:
		return nil, errors.New("invalid reentry_mode")
	}

	switch mode := input.GetRebuyMode(); mode {
	case tablev1.RebuyMode_REBUY_MODE_UNSPECIFIED:

	case tablev1.RebuyMode_REBUY_MODE_DISABLED,
		tablev1.RebuyMode_REBUY_MODE_PLAYER_REQUEST,
		tablev1.RebuyMode_REBUY_MODE_OWNER_APPROVAL:
		rules.RebuyMode = mode

	default:
		return nil, errors.New("invalid rebuy_mode")
	}

	// Apply optional overrides before validating related fields.
	if input.MaxSeats != nil {
		rules.MaxSeats = input.GetMaxSeats()
	}
	if input.StartingStack != nil {
		rules.StartingStack = input.GetStartingStack()
	}
	if input.SmallBlind != nil {
		rules.SmallBlind = input.GetSmallBlind()
	}
	if input.BigBlind != nil {
		rules.BigBlind = input.GetBigBlind()
	}
	if input.TurnTimeoutSeconds != nil {
		rules.TurnTimeoutSeconds = input.GetTurnTimeoutSeconds()
	}
	if input.RebuyLimit != nil {
		rules.RebuyLimit = input.GetRebuyLimit()
	}

	if rules.MaxSeats < minSeats || rules.MaxSeats > maxSeats {
		return nil, errors.New("invalid max_seats")
	}
	if rules.SmallBlind == 0 {
		return nil, errors.New("small_blind must be greater than zero")
	}
	if rules.BigBlind <= rules.SmallBlind {
		return nil, errors.New("big_blind must be greater than small_blind")
	}
	// Prevent overflow when converting the stack limits from BB to chips.
	if rules.BigBlind > math.MaxUint64/maxStartingStackBB {
		return nil, errors.New("big_blind is too large")
	}
	if rules.StartingStack < minStartingStackBB*rules.BigBlind ||
		rules.StartingStack > maxStartingStackBB*rules.BigBlind {
		return nil, errors.New("starting_stack must be between 20 and 200 big blinds")
	}
	switch rules.TurnTimeoutSeconds {
	case minTimeout, defaultTimeout, maxTimeout:
	default:
		return nil, errors.New("turn_timeout_seconds must be 15, 30 or 60")
	}

	return rules, nil
}
