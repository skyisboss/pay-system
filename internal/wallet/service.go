package wallet

import (
	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/internal/util"
)

type Service struct {
	dbClient *ent.Client
	logger   *zerolog.Logger
	provider *Provider
}

func New(client *ent.Client, provider *Provider, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "wallet_service").Logger()
	return &Service{
		dbClient: client,
		provider: provider,
		logger:   &log,
	}
}

func (s *Service) GetProvider(blockchain Blockchain) IProvider {
	p, _ := s.provider.GetProvider(blockchain)
	return p
}

type Blockchain string

const (
	ETH   Blockchain = "eth"
	BTC   Blockchain = "btc"
	BSC   Blockchain = "bsc"
	TRON  Blockchain = "tron"
	MATIC Blockchain = "matic"
)

var blockchains = []Blockchain{BTC, ETH, TRON, MATIC, BSC}

func (b Blockchain) IsValid() bool {
	for _, bc := range blockchains {
		if b == bc {
			return true
		}
	}

	return false
}
func (b Blockchain) String() string {
	return string(b)
}

// 钱包地址信息
type Wallet struct {
	UUID       string     `json:"uuid"`
	Address    string     `json:"address"`
	Blockchain Blockchain `json:"blockchain"`
	PrivateKey string     `json:"private_key"`
}

// 加密密钥
func (w *Wallet) EncodePrivateKey(saltKey string) (string, error) {
	key := util.StrReverse(w.Address+w.UUID) + saltKey
	ret, err := util.AesEncrypt(w.PrivateKey, key)
	if err != nil {
		return "", err
	}
	return ret, nil
}

// 解密密钥
func (w *Wallet) DecodePrivateKey(saltKey string) (string, error) {
	key := util.StrReverse(w.Address+w.UUID) + saltKey
	ret, err := util.AesDecrypt(w.PrivateKey, key)
	if err != nil {
		return "", err
	}
	return ret, nil
}
