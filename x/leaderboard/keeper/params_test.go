package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/cosmonaut/leaderboard/testutil/keeper"
	"github.com/cosmonaut/leaderboard/x/leaderboard/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LeaderboardKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
