package tron

import (
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/wallet"
)

type Provider struct {
	Blockchain wallet.Blockchain
	Container  *ioc.Container
}

func (p *Provider) GetBlockchain() wallet.Blockchain {
	return p.Blockchain
}

func (p *Provider) Ioc() *ioc.Container {
	return p.Container
}
