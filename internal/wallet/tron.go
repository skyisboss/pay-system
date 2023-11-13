package wallet

import "github.com/skyisboss/pay-system/internal/rpc/tronrpc"

type TronProvider struct {
	Blockchain Blockchain
	Client     *tronrpc.Client
}

func (pv *TronProvider) GetBlockchain() Blockchain {
	return pv.Blockchain
}

// 创建钱包
func (p *TronProvider) CreateWallet() *Wallet {
	return nil
}
