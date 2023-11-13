package withdraw

import (
	"time"

	"gorm.io/gorm"
)

type ModelWithdraw struct {
	gorm.Model
	// 产品id
	ProductID int64 `json:"product_id" gorm:"type:int(11);NOT NULL"`
	// 币种
	Symbol string `json:"symbol" gorm:"type:varchar(64);NOT NULL"`
	// 提币地址
	ToAddress string `json:"to_address" gorm:"type:varchar(128);NOT NULL"`
	// 提币金额
	Amount string `json:"amount" gorm:"type:varchar(128);NOT NULL"`
	// 提币序列号 唯一标示 (猜测可能是方便用户使用此号码查询记录)
	SerialID string `json:"serial_id" gorm:"type:varchar(64);NOT NULL"`
	// 提币tx hash
	TxHash string `json:"tx_hash" gorm:"type:varchar(128);NOT NULL"`
	// 处理状态
	HandleStatus string `json:"handle_status" gorm:"type:varchar(32);NOT NULL"`
	// 处理信息
	HandleMsg string `json:"handle_msg" gorm:"type:varchar(128);NOT NULL"`
	// 处理时间
	HandleTime time.Time `json:"handle_time" gorm:"type:TIMESTAMP;NOT NULL;"`
}

func (table *ModelWithdraw) TableName() string {
	return "withdraw"
}
