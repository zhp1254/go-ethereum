package eth

import (
	"github.com/OffchainLabs/go-ethereum/core"
	"github.com/OffchainLabs/go-ethereum/core/state"
	"github.com/OffchainLabs/go-ethereum/core/types"
	"github.com/OffchainLabs/go-ethereum/core/vm"
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

func (eth *Ethereum) StateAtTransaction(block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, error) {
	return eth.stateAtTransaction(block, txIndex, reexec)
}
