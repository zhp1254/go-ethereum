package state

import (
	"errors"

	"github.com/OffchainLabs/go-ethereum/common"
	"github.com/OffchainLabs/go-ethereum/core/rawdb"
	"github.com/OffchainLabs/go-ethereum/ethdb"
)

func (db *cachingDB) ActivatedAsm(target ethdb.WasmTarget, moduleHash common.Hash) ([]byte, error) {
	cacheKey := activatedAsmCacheKey{moduleHash, target}
	if asm, _ := db.activatedAsmCache.Get(cacheKey); len(asm) > 0 {
		return asm, nil
	}
	if asm := rawdb.ReadActivatedAsm(db.wasmdb, target, moduleHash); len(asm) > 0 {
		db.activatedAsmCache.Add(cacheKey, asm)
		return asm, nil
	}
	return nil, errors.New("not found")
}
