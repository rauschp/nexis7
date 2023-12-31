package types

import (
	"github.com/rauschp/nexis7/util"
)

type Wallet struct {
	PublicKey *util.PublicKey
	Address   util.Address
	Balance   float32
	Nonce     int64
}
