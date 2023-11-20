package blockchain

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/blockchain"
	"github.com/skyisboss/pay-system/ent/schema"
	"github.com/skyisboss/pay-system/internal/wallet"
)

type Service struct {
	db     *ent.BlockchainClient
	client *ent.Client
	logger *zerolog.Logger
}

func New(client *ent.Client, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "product_service").Logger()

	return &Service{
		db:     client.Blockchain,
		client: client,
		logger: &log,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]*ent.Blockchain, error) {
	ret, err := s.db.Query().All(ctx)
	return ret, err
}

// 根据币种获取配置
func (s *Service) GetBySymbol(ctx context.Context, symbol string) (*ent.Blockchain, error) {
	ret, err := s.db.Query().Where(blockchain.SymbolEQ(symbol)).First(ctx)
	return ret, err
}

// 根据区块链类型获取配置
func (s *Service) GetByChainAndType(ctx context.Context, chain string, types string) (*ent.Blockchain, *GasPrice, error) {
	ret, err := s.db.Query().
		Where(blockchain.ChainEQ(chain)).
		Where(blockchain.TypesEQ(types)).
		First(ctx)
	if err != nil {
		return nil, nil, err
	}

	return ret, nil, err
}

// 根据chain和symbol获取配置信息
func (s *Service) GetByChainAndSymbol(ctx context.Context, chain, symbol string) (*ent.Blockchain, error) {
	ret, err := s.db.Query().
		Where(blockchain.ChainEQ(chain)).
		Where(blockchain.SymbolEQ(symbol)).
		First(ctx)
	return ret, err
}

// 根据chain获取所有同类网络区块链配置
func (s *Service) GetByChain(ctx context.Context, chain wallet.Blockchain) ([]*ent.Blockchain, error) {
	ret, err := s.db.Query().
		Where(blockchain.ChainEQ(chain.String())).
		All(ctx)
	return ret, err
}

// 更新已处理的区块
func (s *Service) UpdateScanBlockBySymbol(ctx context.Context, symbol string, num int64) error {
	ret, err := s.db.Update().Where(blockchain.SymbolEQ(symbol)).SetScanBlockNum(num).Save(ctx)
	if ret == 0 {
		return errors.New("record not found")
	}
	return err
}
func (s *Service) UpdateScanBlockByID(ctx context.Context, id uint64, num int64) error {
	ret, err := s.db.Update().Where(blockchain.IDEQ(id)).SetScanBlockNum(num).Save(ctx)
	if ret == 0 {
		return errors.New("record not found")
	}
	return err
}

// 获取区块配置
func (s *Service) GetBlockchain_XXXXX(ctx context.Context) ([]*ent.Blockchain, error) {
	ret, err := s.db.Query().Where(blockchain.TypesEQ("coin")).All(ctx)
	return ret, err
}

// 获取区块配置 废弃
func (s *Service) GetByChainXXX(ctx context.Context, chain string) ([]*ent.Blockchain, error) {
	ret, err := s.db.Query().Where(blockchain.ChainEQ(chain)).All(ctx)
	return ret, err
}

// 更新汽油价格
type GasPrice struct {
	UserGasPrice int64 `json:"user_gas_price"`
	ColdGasPrice int64 `json:"cold_gas_price"`
	TipFeePrice  int64 `json:"tip_fee_price"`
}

func (s *Service) UpdateEthGasPrice(ctx context.Context, chainName string, gas schema.GasPrice) (int, error) {
	ret, err := s.db.Update().
		Where(blockchain.ChainEQ(chainName)).
		SetGasPrice(gas).
		Save(ctx)
	return ret, err
}

func (s *Service) GetEthGasPrice1(gasJson string) (*GasPrice, error) {
	var ret GasPrice
	err := json.Unmarshal([]byte(gasJson), &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}
