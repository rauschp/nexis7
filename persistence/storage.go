package persistence

import (
	pb "nexis7/proto"
	"nexis7/types"
	"nexis7/util"
)

type BlockStore interface {
	Height() int64
	Set(block *pb.Block) error
	Get(hash string) (*pb.Block, error)
}

type WalletStore interface {
	GetByPublicKey(pc *util.PublicKey) (*types.Wallet, error)
	GetByAddress(address util.Address) (*types.Wallet, error)
	DepositCurrency(address util.Address, amount float32) error
	WithdrawCurrency(address util.Address, amount float32) error
}
