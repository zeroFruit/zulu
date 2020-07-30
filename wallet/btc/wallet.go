package btc

import (
	"github.com/DE-labtory/zulu/keychain"
)

type Wallet struct {
	addr string
	pub  keychain.PubKey
	pvt  keychain.PvtKey
}
