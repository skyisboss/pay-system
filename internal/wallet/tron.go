package wallet

type TronProvider struct {
	Blockchain Blockchain
}

func (pv *TronProvider) GetBlockchain() Blockchain {
	return pv.Blockchain
}

// 创建钱包
func (p *TronProvider) CreateWallet() *Wallet {
	return nil
}
