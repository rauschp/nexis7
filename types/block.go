package types

import (
	"crypto/sha256"
	protobuf "github.com/golang/protobuf/proto"
	pb "github.com/rauschp/nexis7/proto"
	"github.com/rauschp/nexis7/util"
)

func HashBlock(block *pb.Block) []byte {
	bl, err := protobuf.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(bl)

	return hash[:]
}

func SignBlock(pk *util.PrivateKey, block *pb.Block) []byte {
	sig := pk.Sign(HashBlock(block))

	return sig
}
