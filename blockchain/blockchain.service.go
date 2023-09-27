package blockchain

import (
	"github.com/rauschp/nexis7/types"
	"github.com/rauschp/nexis7/util"
)

type BlockchainService struct {
	ParentNode string
	Height     int64
	Node       *types.NexisNode
}

func (b *BlockchainService) StartService() {
	if b.ParentNode == "" {
		b.startPantheonNode()
	} else {
		b.startReaderNode()
	}
}

func (b *BlockchainService) startPantheonNode() {
	height := b.Node.Datastore.BlockStore.Height()
	if height == 0 {
		// There are no existing blocks, create the genesis block and genesis wallet
		genesisBlock, pubKey := util.CreateGenesisBlock()

		err := b.Node.Datastore.BlockStore.Set(genesisBlock)
		if err != nil {
			panic(err)
		}

		wallet := types.Wallet{
			PublicKey: pubKey,
			Address:   pubKey.GetAddress(),
			Balance:   genesisBlock.Transactions[0].Amount,
			Nonce:     1,
		}

		err = b.Node.Datastore.WalletStore.SaveNewWallet(wallet)
		if err != nil {
			panic(err)
		}
	} else {
		// The chain has already been established, start the node
	}
}

func (b *BlockchainService) startReaderNode() {
	// TODO: Implement later :)
}
