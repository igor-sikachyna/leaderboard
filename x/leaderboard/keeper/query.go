package keeper

import (
	"github.com/igor-sikachyna/leaderboard/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
