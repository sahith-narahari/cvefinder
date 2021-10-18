package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/noandrea/cvefinder/testutil/keeper"
	"github.com/noandrea/cvefinder/x/cvefinder/keeper"
	"github.com/noandrea/cvefinder/x/cvefinder/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CvefinderKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
