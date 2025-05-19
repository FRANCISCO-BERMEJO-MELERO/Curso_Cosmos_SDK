package rps_test

import (
	"testing"

	keepertest "curso/testutil/keeper"
	"curso/testutil/nullify"
	rps "curso/x/rps/module"
	"curso/x/rps/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RpsKeeper(t)
	rps.InitGenesis(ctx, k, genesisState)
	got := rps.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
