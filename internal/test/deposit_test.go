package test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/notify"
	"github.com/skyisboss/pay-system/ent/txn"
	"github.com/skyisboss/pay-system/internal/app/task"
	"github.com/skyisboss/pay-system/internal/app/task/eth"
	"github.com/skyisboss/pay-system/internal/app/task/tron"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
	"github.com/test-go/testify/assert"
)

func TestCheckDeposit(t *testing.T) {
	ctx, _, boot, log := Setup()

	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&eth.Provider{Blockchain: wallet.TRON})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		log,
	)

	// 使用测试时需要到 tasks.EthProvider().CheckDeposit() #makeScanBlockNumber 开启测试模式
	// 注意点：hotWallet 会过滤手续费钱包
	// status 会控制是否检测该区块交易
	// min_deposit 控制最小检测充值金额

	t.Run("check-coin", func(t *testing.T) {
		boot.Ioc().BlockchainService().UpdateScanBlockByID(ctx, 1, 4641969)
		txid := "0x6179ee1e51aae8418ae856000a446de13d8e22106117b9047b441d047e5adb53"
		boot.Ioc().DBClient().Txn.Delete().Where(txn.TxIDEQ(txid)).Exec(ctx)
		boot.Ioc().DBClient().Notify.Delete().Where(notify.ChainID(1)).Where(notify.ProductID(1)).Exec(ctx)

		tasks.EthProvider().CheckDeposit()

		ret, err := boot.Ioc().DBClient().Txn.Query().Where(txn.TxIDEQ(txid)).First(ctx)
		if ent.IsNotFound(err) {
			t.Log("数据不存在123")
			return
		}
		assert.Equal(t, txid, ret.TxID)
	})

	t.Run("check-token", func(t *testing.T) {
		boot.Ioc().BlockchainService().UpdateScanBlockByID(ctx, 2, 4653863)
		txid := "0x1e2daac081a3e611f40c8d29c4b3c494ba73df85701ccfe6ddd0af98d65db4fb"
		boot.Ioc().DBClient().Txn.Delete().Where(txn.TxIDEQ(txid)).Exec(ctx)
		boot.Ioc().DBClient().Notify.Delete().Where(notify.ChainID(2)).Where(notify.ProductID(1)).Exec(ctx)

		tasks.EthProvider().CheckDeposit()

		ret, err := boot.Ioc().DBClient().Txn.Query().Where(txn.TxIDEQ(txid)).First(ctx)
		if ent.IsNotFound(err) {
			t.Log("数据不存在123")
			return
		}
		assert.Equal(t, txid, ret.TxID)
	})
}

func TestNewBalanceAccount(t *testing.T) {
	ctx, _, boot, log := Setup()

	ret, err := boot.Ioc().BalanceService().NewAccount(ctx, 1)
	if err != nil {
		log.Error().Err(err).Msg("err12")
		t.Fail()
	}
	util.ToJson(ret)
}

func TestCheckDepositTron(t *testing.T) {
	ctx, _, boot, log := Setup()

	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&tron.Provider{Blockchain: wallet.TRON, Container: boot.Ioc()})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		log,
	)

	// 使用测试时需要到 tasks.EthProvider().CheckDeposit() #makeScanBlockNumber 开启测试模式
	// 注意点：hotWallet 会过滤手续费钱包
	// status 会控制是否检测该区块交易
	// min_deposit 控制最小检测充值金额

	t.Run("check-coin", func(t *testing.T) {
		boot.Ioc().BlockchainService().UpdateScanBlockByID(ctx, 3, 41246717-1)
		txid := "f7d295ff130a73095a4491fbfaedebffd73fe206754aecea9a2bf0774ad65766"
		amountRaw := decimal.NewFromInt(1414000)
		boot.Ioc().DBClient().Txn.Delete().Where(txn.TxIDEQ(txid)).Exec(ctx)
		boot.Ioc().DBClient().Notify.Delete().Where(notify.ChainID(3)).Where(notify.ProductID(1)).Exec(ctx)

		// tasks.GetProvider(wallet.TRON).(*tron.Provider).CheckDeposit()
		tasks.TronProvider().CheckDeposit()

		ret, err := boot.Ioc().DBClient().Txn.Query().Where(txn.TxIDEQ(txid)).First(ctx)
		if ent.IsNotFound(err) {
			t.Log("数据不存在123")
			return
		}
		assert.Equal(t, txid, ret.TxID, "交易哈希")
		assert.Equal(t, amountRaw, ret.AmountRaw, "交易金额")
	})

	t.Run("check-token", func(t *testing.T) {
		boot.Ioc().BlockchainService().UpdateScanBlockByID(ctx, 4, 41257590-1)
		txid := "35012442a1394e4baeaf407d60c5e8601df165f256e83a3ac8f88dcff15efbf5"
		amountRaw := decimal.NewFromInt(46800000000)
		boot.Ioc().DBClient().Txn.Delete().Where(txn.TxIDEQ(txid)).Exec(ctx)
		boot.Ioc().DBClient().Notify.Delete().Where(notify.ChainID(4)).Where(notify.ProductID(1)).Exec(ctx)

		tasks.TronProvider().CheckDeposit()

		ret, err := boot.Ioc().DBClient().Txn.Query().Where(txn.TxIDEQ(txid)).First(ctx)
		if ent.IsNotFound(err) {
			t.Log("数据不存在123")
			return
		}
		assert.Equal(t, txid, ret.TxID, "交易哈希")
		assert.Equal(t, amountRaw, ret.AmountRaw, "交易金额")
	})
}
