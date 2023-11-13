package task

import (
	"sync"

	"github.com/skyisboss/pay-system/internal/wallet"
)

type Provider interface {
	GetBlockchain() wallet.Blockchain
}

type Handler struct {
	mu        sync.RWMutex
	providers map[wallet.Blockchain]Provider
}

func NewProvider() *Handler {
	return &Handler{
		mu:        sync.RWMutex{},
		providers: make(map[wallet.Blockchain]Provider),
	}
}

func (h *Handler) AddProvider(provider Provider) *Handler {
	h.mu.Lock()
	h.providers[provider.GetBlockchain()] = provider
	h.mu.Unlock()

	return h
}
