package txn

import (
	"time"

	"gorm.io/gorm"
)

type ModelTxn struct {
	gorm.Model
	// 交易id
	TxID string `json:"tx_id" gorm:"type:varchar(128);NOT NULL"`
	// 币种
	Symbol string `json:"symbol" gorm:"type:varchar(64);NOT NULL"`
	// 商户id
	ProductID int64 `json:"product_id" gorm:"type:int(11);NOT NULL"`
	// 原来地址
	FromAddress string `json:"from_address" gorm:"type:varchar(128);NOT NULL"`
	// 目标地址
	ToAddress string `json:"to_address" gorm:"type:varchar(128);NOT NULL"`
	// 金额
	Amount string `json:"amount" gorm:"type:varchar(128);NOT NULL"`
	// 处理状态
	HandleStatus int64 `json:"handle_status" gorm:"type:varchar(32);NOT NULL"`
	// 处理信息
	HandleMsg string `json:"handle_msg" gorm:"type:varchar(128);NOT NULL"`
	// 处理时间
	HandleTime time.Time `json:"handle_time"`
	// 零钱归集
	TransferStatus int64     `json:"transfer_status" gorm:"type:varchar(32);NOT NULL"`
	TransferMsg    string    `json:"transfer_msg" gorm:"type:varchar(128);NOT NULL"`
	TransferTime   time.Time `json:"transfer_time"`
}

func (table *ModelTxn) TableName() string {
	return "txn"
}
