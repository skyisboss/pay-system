package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

type Provider struct {
	Blockchain wallet.Blockchain
	ioc        *ioc.Container
}

func (p *Provider) SetInit(ioc *ioc.Container) *Provider {
	p.ioc = ioc
	return p
}

func (p *Provider) GetBlockchain() wallet.Blockchain {
	return p.Blockchain
}

// eth钱包地址
type Address string

func (a Address) String() string {
	return string(a)
}

// blockchain表ID主键
type ChainID int64

func (c ChainID) Uint() uint64 {
	return uint64(c)
}
func (c ChainID) String() string {
	return strconv.Itoa(int(c))
}

// 数据库id主键
type PkID uint64

func (d PkID) Int() int64 {
	return int64(d)
}
func (d PkID) Uint() uint64 {
	return uint64(d)
}

// 转换以太坊地址为Hex
func (a Address) HexAddress() common.Address {
	return common.HexToAddress(a.String())
}

// 钱包信息
type WalletInfo struct {
	ID      uint64
	UUID    string
	Address string
	Private *ecdsa.PrivateKey
	UseTo   int64
	ChainID uint64
}

// 获取钱包信息
func (p *Provider) GetWallet(ctx context.Context, address []Address) (map[Address]*WalletInfo, error) {
	addressRows, err := p.ioc.AddressService().ListByAddressIn(ctx, util.MapSlice(address, func(a Address) string {
		return a.String()
	}))
	if err != nil {
		p.ioc.Logger().Error().Err(err).Msg("err")
		return nil, err
	}
	// 解密私钥
	privateKeyMap := make(map[Address]*WalletInfo)
	for _, itemRow := range addressRows {
		saltKey := util.StrReverse(itemRow.Address+itemRow.UUID) + p.ioc.Config().Providers.SaltKey
		key, err := util.AesDecrypt(itemRow.Password, saltKey)
		if err != nil {
			p.ioc.Logger().Error().Err(err).Msg("err")
			return nil, err
		}
		if len(key) == 0 {
			err := errors.New("key len = 0")
			p.ioc.Logger().Error().Err(err).Msg(err.Error())
			return nil, err
		}
		key = strings.TrimPrefix(key, "0x")
		privateKey, err := crypto.HexToECDSA(key)
		if err != nil {
			p.ioc.Logger().Error().Err(err).Msg("err")
			return nil, err
		}
		privateKeyMap[Address(itemRow.Address)] = &WalletInfo{
			ID:      itemRow.ID,
			UUID:    itemRow.UUID,
			Address: itemRow.Address,
			Private: privateKey,
			UseTo:   itemRow.UseTo,
			ChainID: itemRow.ChainID,
		}
	}
	return privateKeyMap, nil
}

// 获取当前链上最新区块数
func (p *Provider) RpcGetBlockNumber(ctx context.Context) (int64, error) {
	rpcNum, err := p.ioc.EthClient().BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	return int64(rpcNum), nil
}

// 根据区块数获取区块链信息
func (p *Provider) RpcGetBlockByNumber(ctx context.Context, blockNum int64) (*types.Block, error) {
	// rpc获取block信息
	rpcBlock, err := p.ioc.EthClient().BlockByNumber(ctx, big.NewInt(blockNum))
	if err != nil {
		return nil, err
	}
	return rpcBlock, nil
}

// 从给定区块中获取log日志信息
func (p *Provider) RpcGetLogsByNumber(ctx context.Context, abi abi.ABI, blockNum int64) ([]types.Log, error) {
	// rpc获取block信息
	var warpAddresses []common.Address
	event := abi.Events["Transfer"]
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNum),
		ToBlock:   big.NewInt(blockNum),
		Addresses: warpAddresses,
		Topics: [][]common.Hash{
			{event.ID},
		},
	}

	// {
	// 	"address": "0x48cfd205698120a01fec4481cfc5e8a205677791",
	// 	"blockHash": "0x3b1e9c93bc0030a23cc9cf2dd430cd8d58aba27b488007b8173d13b95d9a2124",
	// 	"blockNumber": "0x4526b9",
	// 	"data": "0x00000000000000000000000000000000000000000000000000000000054c3380",
	// 	"logIndex": "0x14",
	// 	"removed": false,
	// 	"topics": [
	// 		"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
	// 		"0x00000000000000000000000006edba1bf11ff69bfaaec69cf6fc614856989272",
	// 		"0x00000000000000000000000006edba1bf11ff69bfaaec69cf6fc614856989272"
	// 		],
	// 	"transactionHash": "0x4358c01d5752a79449f6f6f2e58025f35c2d46d67b5e3eb5cad74c546213b612",
	// 	"transactionIndex": "0xb"
	// },

	rpcLogs, err := p.ioc.EthClient().FilterLogs(ctx, query)
	if err != nil {
		p.ioc.Logger().Error().Err(err).Msg("err")
		return nil, err
	}

	return rpcLogs, nil
}
