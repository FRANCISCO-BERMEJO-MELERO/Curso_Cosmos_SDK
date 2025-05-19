package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "curso/testutil/keeper"
	"curso/x/mimodulo/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MimoduloKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
