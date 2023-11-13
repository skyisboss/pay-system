package blockchain

import "gorm.io/gorm"

type ModelBlockchain struct {
	gorm.Model

	// 区块链名称
	Chain string `json:"chain" gorm:"type:varchar(32);NOT NULL"`
	// 显示名称
	Token string `json:"token" gorm:"type:varchar(32);NOT NULL"`
	// 类型 coin/token
	Type string `json:"type" gorm:"type:varchar(16);NOT NULL"`
	// 币种
	Symbol string `json:"symbol" gorm:"type:varchar(32);NOT NULL"`
	// 精度
	Decimals int64 `json:"decimals" gorm:"type:int(11);NOT NULL"`
	// 状态
	Status int64 `json:"status" gorm:"type:tinyint(2);NOT NULL"`
	// token合约地址
	TokenAddress string `json:"token_address" gorm:"type:varchar(128);NOT NULL"`
	// 冷钱包地址 零钱整理时转移到这里
	ColdAddress string `json:"cold_address" gorm:"type:varchar(128);NOT NULL"`
	// 热钱包地址 用户提款时从这里扣款
	HotAddress string `json:"hot_address" gorm:"type:varchar(128);NOT NULL"`
	// 已处理的区块索引
	ScanBlockNum int64 `json:"scan_block_num" gorm:"type:int(11);NOT NULL"`
	// 最小空闲地址数
	MinFreeNum int64 `json:"min_free_num" gorm:"type:int(11);NOT NULL"`
	// 最小区块确认数
	MinConfirmNum int64 `json:"min_confirm_num" gorm:"type:int(11);NOT NULL"`
	// 提款手续费费用
	WithdrawFee float64 `json:"withdraw_fee" gorm:"type:decimal(10,2);NOT NULL"`
	// 提款手续费类型 1-百分比收取 2-单笔固定收取
	WithdrawFeeType int64 `json:"withdraw_fee_type" gorm:"type:int(2);NOT NULL"`
	// 最小充值金额
	MinDeposit int64 `json:"min_deposit" gorm:"type:int(11);NOT NULL"`
	// 最小提款金额
	MinWithdraw int64 `json:"min_withdraw" gorm:"type:int(11);NOT NULL"`
	// 最小转账金额 可用作零钱整理状态，可能暂时用不到
	// MinTransfer int64 `json:"min_transfer" gorm:"type:int(11);NOT NULL"`
}

func (table *ModelBlockchain) TableName() string {
	return "blockchain"
}
