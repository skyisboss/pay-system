package notify

import (
	"time"

	"gorm.io/gorm"
)

type ModelNotify struct {
	gorm.Model
	Nonce string `json:"nonce" gorm:"type:varchar(32);NOT NULL"`
	// 商户id
	ProductID int64 `json:"product_id" gorm:"type:int(11);NOT NULL"`
	// 操作类型
	ItemType int64 `json:"item_type" gorm:"type:int(11);NOT NULL"`
	// 关联id
	ItemID int64 `json:"item_id" gorm:"type:int(11);NOT NULL"`
	// 通知类型 1-冲币到账通知 2-提币广播通知 3-提币到账通知
	NotifyType string `json:"notify_type" gorm:"type:varchar(32);NOT NULL"`
	// 币种
	Symbol string `json:"symbol" gorm:"type:varchar(64);NOT NULL"`
	// 发送通知地址
	SendUrl string `json:"send_url" gorm:"type:varchar(512);NOT NULL"`
	// 发送通知内容
	SendBody string `json:"send_body" gorm:"type:varchar(4089);NOT NULL"`
	// 通知失败重试次数
	SendRetry int64 `json:"send_retry" gorm:"type:int(11);NOT NULL;DEFAULT 0"`
	// 处理状态 0-默认 1-通知成功 2-通知失败
	HandleStatus string `json:"handle_status" gorm:"type:varchar(32);NOT NULL"`
	// 处理时间
	HandleTime time.Time `json:"handle_time"`
	// 处理信息
	HandleMsg string `json:"handle_msg" gorm:"type:varchar(512);NOT NULL"`
}

func (table *ModelNotify) TableName() string {
	return "notify"
}
