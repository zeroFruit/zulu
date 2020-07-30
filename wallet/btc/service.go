package btc

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type Service struct{}

func NewService() *Service {
	return nil
}

func (s *Service) CreateNormalWallet(id string) *Wallet {
	keychain.GenerateKey(id, types.Bitcoin)
	return nil
}

func (s *Service) CreateMultiSigWallet(id string) *Wallet {
	return nil
}

func (s *Service) Transfer(id string, to string, amount uint) *Transaction {
	return nil
}
