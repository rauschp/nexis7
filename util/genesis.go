package util

import (
	pb "github.com/rauschp/nexis7/proto"
	"os"
	"time"
)

func CreateGenesisBlock() (*pb.Block, *PublicKey) {
	// Create genesis wallet
	key := GenerateNewPrivateKey()
	encoded := key.ToBase64()

	f, err := os.Create(DefaultGenesisKeyPath)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	_, err = f.WriteString(encoded)
	if err != nil {
		panic(err)
	}

	err = f.Sync()
	if err != nil {
		panic(err)
	}

	block := &pb.Block{
		Header: &pb.Header{
			Version:   "0.1.0",
			State:     pb.BlockState_Finalized,
			Height:    0,
			Timestamp: time.Now().Unix(),
		},
		Transactions: []*pb.Transaction{
			{
				Version:   "0.1.0",
				State:     pb.TransactionState_Confirmed,
				Nonce:     1,
				Amount:    10000000,
				ToAddress: key.Public().GetAddress().ToBytes(),
				Timestamp: time.Now().Unix(),
			},
		},
	}

	return block, key.Public()
}
