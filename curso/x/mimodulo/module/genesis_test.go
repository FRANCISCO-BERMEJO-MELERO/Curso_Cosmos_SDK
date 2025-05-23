package mimodulo_test

import (
	"testing"

	keepertest "curso/testutil/keeper"
	"curso/testutil/nullify"
	mimodulo "curso/x/mimodulo/module"
	"curso/x/mimodulo/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MimoduloKeeper(t)
	mimodulo.InitGenesis(ctx, k, genesisState)
	got := mimodulo.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
