package arbitrum

import (
	"context"

	"github.com/OffchainLabs/go-ethereum/core"
	"github.com/OffchainLabs/go-ethereum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
