package wallet

import (
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/skyisboss/pay-system/internal/rpc/ethrpc"
	"github.com/skyisboss/pay-system/internal/util"
)

type EthProvider struct {
	Blockchain Blockchain
	Client     *ethrpc.Client
}

// ethAddressRegex see https://goethereumbook.org/en/address-check/
var ethAddressRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

func (pv *EthProvider) GetBlockchain() Blockchain {
	return pv.Blockchain
}

func (p *EthProvider) ValidateAddress(address string) bool {
	return ethAddressRegex.MatchString(address)
}

// 创建钱包
func (p *EthProvider) CreateWallet() *Wallet {
	key, err := crypto.GenerateKey()
	if err != nil {
		return &Wallet{}
	}
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := hexutil.Encode(crypto.FromECDSA(key))

	return &Wallet{
		UUID:       util.GetUUID(),
		Blockchain: p.Blockchain,
		Address:    strings.ToLower(address),
		PrivateKey: privateKey,
	}
}

// 获取钱包nonce
func (p *EthProvider) RpcNonce(address string) (uint64, error) {
	return 0, nil
}

// 获取最新区块
func (p *EthProvider) RpcBlockNumber() (uint64, error) {
	return 0, nil
}

// 根据区块数获取区块信息
func (p *EthProvider) RpcBlockByNumber(num uint64) (uint64, error) {
	return 0, nil
}

// 根据区块数获取日志信息
func (p *EthProvider) RpcLogsByNumber(num uint64) (uint64, error) {
	return 0, nil
}

// 获取主币余额
func (p *EthProvider) RpcBalanceAt(address string) (uint64, error) {
	return 0, nil
}

// 获取代币余额
func (p *EthProvider) RpcTokenBalanceAt(address string) (uint64, error) {
	return 0, nil
}

// 构建交易
func (p *EthProvider) NewTransaction(w *Wallet) (string, error) {
	return "", nil
}

// 发送交易
func (p *EthProvider) SendTransaction(w *Wallet) (string, error) {
	return "", nil
}
