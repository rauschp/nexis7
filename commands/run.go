package commands

import (
	"google.golang.org/grpc"
	"nexis7/persistence"
	pb "nexis7/proto"
	"nexis7/util"
	"os"
	"sync"
)

type NexisNode struct {
	Version     string
	Host        string
	PrivateKey  util.PrivateKey
	PeerManager *PeerManager
	Mempool     *Mempool
	Datastore   *PersistenceManager
}

type PeerManager struct {
	Peers map[string]*Peer
	Lock  sync.RWMutex
}

type PersistenceManager struct {
	BlockStore  persistence.BlockStore
	WalletStore persistence.WalletStore
}

type Mempool struct {
	Transactions map[string]*pb.Transaction
	Lock         sync.RWMutex
}

type Peer struct {
	Version    string
	Host       string
	Connection *grpc.ClientConn
}

func ProcessRun() error {
	return nil
}

func createNode(parentNode string) *NexisNode {
	var key, err = readPrivateKey()
	if err != nil {
		panic(err)
	}

	return &NexisNode{
		Version:    "0.0.1",
		Host:       ":9999",
		PrivateKey: *key,
		PeerManager: &PeerManager{
			Peers: make(map[string]*Peer),
		},
		Mempool: &Mempool{
			Transactions: make(map[string]*pb.Transaction),
		},
		Datastore: &PersistenceManager{
			BlockStore:  persistence.CreatePersistentBlockstore(),
			WalletStore: persistence.CreatePersistentWalletStore(),
		},
	}
}

func readPrivateKey() (*util.PrivateKey, error) {
	f, err := os.ReadFile(util.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	contents := string(f)

	key, err := util.GeneratePrivateKeyFromBase64(contents)
	if err != nil {
		return nil, err
	}

	return key, nil
}
