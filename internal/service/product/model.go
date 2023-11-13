package product

import "gorm.io/gorm"

type ModelProduct struct {
	gorm.Model
	// 商户id
	AppID string `json:"app_id" gorm:"type:varchar(64);NOT NULL"`
	// 商户应用名称
	AppName string `json:"app_name" gorm:"type:varchar(64);NOT NULL"`
	// 商户密钥
	AppSecret string `json:"app_secret" gorm:"type:varchar(64);NOT NULL"`
	// 商户状态 0-禁止 1-正常
	AppStatus int64 `json:"app_status" gorm:"type:tinyint(2);NOT NULL"`
	// 结算状态 0-禁止 1-自动结算 2-人工审核
	WithdrawStatus int64 `json:"withdraw_status" gorm:"type:tinyint(2);NOT NULL"`
	// 通知回调地址
	WebHook string `json:"web_hook" gorm:"type:varchar(512);NOT NULL"`
}

func (table *ModelProduct) TableName() string {
	return "product"
}
