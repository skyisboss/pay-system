package wallet

import (
	btcec "github.com/btcsuite/btcd/btcec/v2"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/skyisboss/pay-system/internal/rpc/tronrpc"
	"github.com/skyisboss/pay-system/internal/util"
)

type TronProvider struct {
	Blockchain Blockchain
	Client     *tronrpc.Client
}

func (pv *TronProvider) GetBlockchain() Blockchain {
	return pv.Blockchain
}

// 创建钱包
func (p *TronProvider) CreateWallet() *Wallet {
	pri, err := btcec.NewPrivateKey()
	if err != nil {
		return nil
	}
	if len(pri.Key.Bytes()) != 32 {
		for {
			pri, err = btcec.NewPrivateKey()
			if err != nil {
				continue
			}
			if len(pri.Key.Bytes()) == 32 {
				break
			}
		}
	}

	address := addr.PubkeyToAddress(pri.ToECDSA().PublicKey).String()
	privateKey := pri.Key.String()

	return &Wallet{
		UUID:       util.GetUUID(),
		Blockchain: p.Blockchain,
		Address:    address,
		PrivateKey: privateKey,
	}
}

// 获取最新区块
func (p *TronProvider) RpcBlockNumber() (int64, error) {
	ret, err := p.Client.GRPC.GetNowBlock()
	if err != nil {
		return 0, err
	}
	return ret.BlockHeader.RawData.Number, nil
}
