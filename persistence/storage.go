package persistence

import (
	pb "github.com/rauschp/nexis7/proto"
	"github.com/rauschp/nexis7/types"
	"github.com/rauschp/nexis7/util"
)

type BlockStore interface {
	Height() int64
	Set(block *pb.Block) error
	Get(hash string) (*pb.Block, error)
}

type WalletStore interface {
	GetByPublicKey(pc *util.PublicKey) (*types.Wallet, error)
	GetByAddress(address util.Address) (*types.Wallet, error)
	SaveNewWallet(wallet types.Wallet) error
	DepositCurrency(address util.Address, amount float32) error
	WithdrawCurrency(address util.Address, amount float32) error
}
