package transfer

import (
	"time"

	"gorm.io/gorm"
)

type ModelTransfer struct {
	gorm.Model
	// 关联类型 1-零钱整理 2-提币
	RelatedType int64 `json:"related_type" gorm:"type:tinyint(4);NOT NULL"`
	// 关联id
	RelatedID int64 `json:"related_id" gorm:"type:int(11);NOT NULL"`
	// 交易id
	TxID string `json:"tx_id" gorm:"type:varchar(128);DEFAULT '';NOT NULL"`
	// 原来地址
	FromAddress string `json:"from_address" gorm:"type:varchar(128);DEFAULT '';NOT NULL"`
	// 目标地址
	ToAddress string `json:"to_address" gorm:"type:varchar(128);DEFAULT '';NOT NULL"`
	// 金额
	Amount string `json:"amount" gorm:"type:varchar(128);NOT NULL"`
	// 处理状态
	HandleStatus int64 `json:"handle_status" gorm:"type:varchar(32);NOT NULL"`
	// 处理信息
	HandleMsg string `json:"handle_msg" gorm:"type:varchar(128);DEFAULT '';NOT NULL"`
	// 处理时间
	HandleTime time.Time `json:"handle_time" gorm:"type:TIMESTAMP;NOT NULL;"`
	// gas
	Gas int64 `json:"gas" gorm:"type:bigint(20);NOT NULL"`
	// gas_price
	GasPrice int64 `json:"gas_price" gorm:"type:bigint(20);NOT NULL"`
	// nonce
	Nonce int64 `json:"nonce" gorm:"type:int(20);NOT NULL"`
	// hex
	Hex string `json:"hex" gorm:"type:varchar(2048);NOT NULL"`
}

func (table *ModelTransfer) TableName() string {
	return "transfer"
}
