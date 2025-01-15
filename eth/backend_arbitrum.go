package eth

import (
	"context"

	"github.com/OffchainLabs/go-ethereum/core"
	"github.com/OffchainLabs/go-ethereum/core/state"
	"github.com/OffchainLabs/go-ethereum/core/types"
	"github.com/OffchainLabs/go-ethereum/core/vm"
	"github.com/OffchainLabs/go-ethereum/eth/tracers"
	"github.com/OffchainLabs/go-ethereum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*types.Transaction, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}
