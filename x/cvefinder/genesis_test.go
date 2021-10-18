package cvefinder_test

import (
	"testing"

	keepertest "github.com/noandrea/cvefinder/testutil/keeper"
	"github.com/noandrea/cvefinder/x/cvefinder"
	"github.com/noandrea/cvefinder/x/cvefinder/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CvefinderKeeper(t)
	cvefinder.InitGenesis(ctx, *k, genesisState)
	got := cvefinder.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
