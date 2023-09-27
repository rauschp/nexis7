package types

import (
	"github.com/rauschp/nexis7/persistence"
	pb "github.com/rauschp/nexis7/proto"
	"github.com/rauschp/nexis7/util"
	"google.golang.org/grpc"
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

func CreateNode(host, version string) *NexisNode {
	var key, err = readPrivateKey()
	if err != nil {
		panic(err)
	}

	return &NexisNode{
		Version:    version,
		Host:       host,
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
	f, err := os.ReadFile(util.DefaultPrivateKeyPath)
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
