package keychain

import (
	"github.com/DE-labtory/zulu/types"
)

type PubKey interface {
	Encode() string
}

type PvtKey interface {
	Encode() string
}

func GenerateKey(id string, platform types.Platform) (PubKey, PvtKey) {
	return nil, nil
}

func Sign() {}
