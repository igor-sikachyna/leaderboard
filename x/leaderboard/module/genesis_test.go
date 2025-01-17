package leaderboard_test

import (
	"testing"

	keepertest "github.com/igor-sikachyna/leaderboard/testutil/keeper"
	"github.com/igor-sikachyna/leaderboard/testutil/nullify"
	leaderboard "github.com/igor-sikachyna/leaderboard/x/leaderboard/module"
	"github.com/igor-sikachyna/leaderboard/x/leaderboard/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LeaderboardKeeper(t)
	leaderboard.InitGenesis(ctx, k, genesisState)
	got := leaderboard.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
