package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/babylonchain/babylon/testutil/keeper"
	"github.com/babylonchain/babylon/x/btcstaking/keeper"
	"github.com/babylonchain/babylon/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BTCStakingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
