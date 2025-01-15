package arbitrum

import (
	"context"

	"github.com/OffchainLabs/go-ethereum/common/hexutil"
	"github.com/OffchainLabs/go-ethereum/core"
	"github.com/OffchainLabs/go-ethereum/internal/ethapi"
	"github.com/OffchainLabs/go-ethereum/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *ethapi.StateOverride, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, overrides, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}
