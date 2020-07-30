package btc

type Tx struct{}

type TxOutput struct{}

type TxOutputLister interface {
	ListByAddr(addr string) []TxOutput
}

type txFactory struct {
	lister TxOutputLister
}

func (f *txFactory) create() *Tx {
	return nil
}
