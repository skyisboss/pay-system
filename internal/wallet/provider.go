package wallet

import (
	"sync"

	"github.com/pkg/errors"
)

var (
	ErrInvalidAddress         = errors.New("invalid address")
	ErrInvalidContractAddress = errors.New("invalid contract address")
	ErrInvalidAmount          = errors.New("invalid amount")
	ErrInvalidNetwork         = errors.New("invalid network")
	ErrInvalidGasSettings     = errors.New("invalid network gas settings")
	ErrInvalidNonce           = errors.New("invalid nonce")
	ErrTronResponse           = errors.New("invalid response from TRON node")
	ErrInsufficientBalance    = errors.New("sender balance is insufficient")
	ErrUnknownBlockchain      = errors.New("unknown blockchain")
)

// 接口
type IProvider interface {
	CreateWallet() *Wallet
	GetBlockchain() Blockchain
}

// 区块网络提供者
type Provider struct {
	mu        sync.RWMutex
	providers map[Blockchain]IProvider
}

func NewProvider() *Provider {
	return &Provider{
		mu:        sync.RWMutex{},
		providers: make(map[Blockchain]IProvider),
	}
}

func (p *Provider) AddProvider(ip IProvider) *Provider {
	p.mu.Lock()
	p.providers[ip.GetBlockchain()] = ip
	p.mu.Unlock()

	return p
}

func (p *Provider) GetProvider(blockchain Blockchain) (IProvider, error) {
	if !blockchain.IsValid() {
		return nil, ErrUnknownBlockchain
	}

	var selectedProvider IProvider

	p.mu.RLock()
	if selected, ok := p.providers[blockchain]; ok {
		selectedProvider = selected
	}
	p.mu.RUnlock()

	if selectedProvider == nil {
		return nil, errors.Wrap(ErrUnknownBlockchain, "provider not found")
	}

	return selectedProvider, nil
}
