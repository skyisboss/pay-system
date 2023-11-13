package cmd

import (
	"context"
	"fmt"

	"github.com/skyisboss/pay-system/internal/boot"
	"github.com/skyisboss/pay-system/internal/service/address"
	"github.com/skyisboss/pay-system/internal/service/balance"
	"github.com/skyisboss/pay-system/internal/service/blockchain"
	"github.com/skyisboss/pay-system/internal/service/notify"
	"github.com/skyisboss/pay-system/internal/service/product"
	"github.com/skyisboss/pay-system/internal/service/transfer"
	"github.com/skyisboss/pay-system/internal/service/txn"
	"github.com/skyisboss/pay-system/internal/service/user"
	"github.com/skyisboss/pay-system/internal/service/withdraw"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var migrateCommand = &cobra.Command{
	Use:     "migrate",
	Short:   "数据库迁移",
	Example: "migrate [TableName|all]",
	Run:     migrRun,
}

func migrRun(_ *cobra.Command, args []string) {
	var (
		ctx     = context.Background()
		cfg     = RegisterConfig()
		service = boot.New(ctx, cfg)
		logger  = service.Logger()
		// db      = service.Ioc().DB().Instance()
	)
	var db *gorm.DB

	switch args[0] {
	case "blockchain":
		initBlockchain(db)
	case "product":
		initProduct(db)
	case "user":
		initUser(db)
	case "address":
		initAddress(db)
	case "notify":
		initNotify(db)
	case "withdraw":
		initWithdraw(db)
	case "transfer":
		initTransfer(db)
	case "txn":
		initTxn(db)
	case "balance":
		initBalance(db)
	case "all":
		initBlockchain(db)
		initProduct(db)
		initUser(db)
		initAddress(db)
		initNotify(db)
		initWithdraw(db)
		initTransfer(db)
		initTxn(db)
		initBalance(db)
	default:
		logger.Error().Msgf("table=%s 不存在", args[0])
		return
	}
	fmt.Println("\r\nsuccess")
}

func initBlockchain(db *gorm.DB) {
	table := blockchain.ModelBlockchain{}

	db.Migrator().DropTable(&table)
	db.AutoMigrate(&table)
	type TokenConfigIndex struct {
		blockchain.ModelBlockchain
		Chain  string `gorm:"index:chain_symbol,unique"`
		Symbol string `gorm:"index:chain_symbol,unique"`
	}
	db.Migrator().CreateIndex(&TokenConfigIndex{}, "chain_symbol")

	var data = []blockchain.ModelBlockchain{
		{
			Chain:           "eth",
			Type:            "coin",
			Symbol:          "eth",
			Token:           "eth",
			Status:          1,
			MinFreeNum:      10,
			MinConfirmNum:   6,
			WithdrawFeeType: 1,
			WithdrawFee:     6.8,
			MinWithdraw:     10,
			MinDeposit:      1,
			Decimals:        18,
		},
		{
			Chain:           "eth",
			Type:            "token",
			Symbol:          "usdt",
			Token:           "erc20",
			Status:          1,
			MinConfirmNum:   6,
			WithdrawFeeType: 1,
			WithdrawFee:     6.8,
			MinWithdraw:     10,
			MinDeposit:      1,
			Decimals:        6,
		},

		{
			Chain:           "tron",
			Type:            "coin",
			Symbol:          "trx",
			Token:           "trx",
			Status:          1,
			MinFreeNum:      10,
			MinConfirmNum:   6,
			WithdrawFeeType: 1,
			WithdrawFee:     6.8,
			MinWithdraw:     10,
			MinDeposit:      1,
			Decimals:        6,
		},
		{
			Chain:           "tron",
			Type:            "token",
			Symbol:          "usdt",
			Token:           "erc20",
			Status:          1,
			MinConfirmNum:   6,
			WithdrawFeeType: 1,
			WithdrawFee:     6.8,
			MinWithdraw:     10,
			MinDeposit:      1,
			Decimals:        6,
		},
	}

	db.Create(&data)
}

func initProduct(db *gorm.DB) {
	table := product.ModelProduct{}

	db.Migrator().DropTable(&table)
	db.AutoMigrate(&table)
	type Index struct {
		product.ModelProduct
		AppID string `gorm:"index:app_idx,unique"`
	}
	db.Migrator().CreateIndex(&Index{}, "app_idx")
}

func initUser(db *gorm.DB) {
	table := user.ModelUser{}

	// 创建数据表 如果存在先删除
	db.Migrator().DropTable(&table)
	// 再创建
	db.AutoMigrate(&table)
	// 设置主键
	type AdminUserIndex struct {
		user.ModelUser
		Username string `gorm:"index:username,unique"`
	}
	db.Migrator().CreateIndex(&AdminUserIndex{}, "username")

	// 默认账号密码 admin/admin
	db.Create(&user.ModelUser{
		Role:     1,
		Username: "admin",
		Password: "$2a$10$kZ3uj8wkpk9U70LLyeIdGe9azgZJm3AIz4vRddClQBKMEz2vfjxYa",
	})
}

func initAddress(db *gorm.DB) {
	table := address.ModelAddress{}

	db.Migrator().DropTable(&table)
	db.AutoMigrate(&table)
	type Index struct {
		address.ModelAddress
		// ID      string `gorm:"index:id,unique"`
		Chain   string `gorm:"index:chain_address_idx,unique"`
		Address string `gorm:"index:chain_address_idx,unique"`
	}
	// db.Migrator().CreateIndex(&Index{}, "id")
	db.Migrator().CreateIndex(&Index{}, "chain_address_idx")
}

func initNotify(db *gorm.DB) {
	table := notify.ModelNotify{}

	db.Migrator().DropTable(&table)
	db.AutoMigrate(&table)
	type Index struct {
		notify.ModelNotify
		ProductID  int64  `gorm:"index:product_id,unique"`
		ItemType   int64  `gorm:"index:product_id,unique"`
		ItemID     int64  `gorm:"index:product_id,unique"`
		NotifyType string `gorm:"index:product_id,unique"`
		// TokenSymbol string `gorm:"index:product_id,unique"`
	}
	db.Migrator().CreateIndex(&Index{}, "product_id")
}

func initWithdraw(db *gorm.DB) {
	table := withdraw.ModelWithdraw{}

	db.Migrator().DropTable(&table)
	db.AutoMigrate(&table)
	type WithdrawIndex struct {
		withdraw.ModelWithdraw
		TxHash    string `gorm:"index:tx_hash"`
		SerialID  string `gorm:"index:serial_idx,unique"`
		ProductID string `gorm:"index:serial_idx,unique"`

		HandleStatus string `gorm:"index:handle_status_symbol_idx,unique"`
		Symbol       string `gorm:"index:handle_status_symbol_idx,unique"`
	}
	db.Migrator().CreateIndex(&WithdrawIndex{}, "tx_hash")
	db.Migrator().CreateIndex(&WithdrawIndex{}, "serial_idx")
	db.Migrator().CreateIndex(&WithdrawIndex{}, "handle_status_symbol_idx")
}

func initTransfer(db *gorm.DB) {
	table := transfer.ModelTransfer{}

	db.Migrator().DropTable(&table)
	db.AutoMigrate(&table)

	type TransferIndex struct {
		transfer.ModelTransfer
		FromAddress string `gorm:"index:from_address"`
		TxID        string `gorm:"index:related_id,unique"`
		RelatedID   int64  `gorm:"index:related_id,unique"`
		RelatedType int64  `gorm:"index:related_id,unique"`
	}
	type TransferIndex2 struct {
		transfer.ModelTransfer
		TxID string `gorm:"index:tx_id"`
	}
	db.Migrator().CreateIndex(&TransferIndex2{}, "tx_id")
	db.Migrator().CreateIndex(&TransferIndex{}, "from_address")
	db.Migrator().CreateIndex(&TransferIndex{}, "related_id")
}

func initTxn(db *gorm.DB) {
	table := txn.ModelTxn{}

	db.Migrator().DropTable(&table)
	db.AutoMigrate(&table)

	type TxIndex struct {
		txn.ModelTxn
		TxID           string `gorm:"index:tx_id,unique"`
		TransferStatus string `gorm:"index:transfer_status_idx"`
	}
	db.Migrator().CreateIndex(&TxIndex{}, "tx_id")
	db.Migrator().CreateIndex(&TxIndex{}, "transfer_status_idx")
}

func initBalance(db *gorm.DB) {
	table := balance.ModelBalance{}

	db.Migrator().DropTable(&table)
	db.AutoMigrate(&table)

	type BalanceIndex struct {
		balance.ModelBalance
		ProductID string `gorm:"index:product_symbol_idx,unique"`
		Symbol    string `gorm:"index:product_symbol_idx,unique"`
	}
	db.Migrator().CreateIndex(&BalanceIndex{}, "product_symbol_idx")

	type BalanceIndex2 struct {
		balance.ModelBalance
		ProductID string `gorm:"index:product_idx"`
	}
	db.Migrator().CreateIndex(&BalanceIndex2{}, "product_idx")
}
