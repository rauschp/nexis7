package commands

import (
	"github.com/rauschp/nexis7/blockchain"
	"github.com/rauschp/nexis7/types"
)

func ProcessRunCommand() error {
	node := types.CreateNode(":9999", "0.1.0")

	service := &blockchain.BlockchainService{
		ParentNode: "",
		Node:       node,
	}
	service.StartService()

	return nil
}
