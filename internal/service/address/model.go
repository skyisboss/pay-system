package address

import (
	"gorm.io/gorm"
)

type ModelAddress struct {
	gorm.Model
	// 所属区块
	Chain string `json:"chain" gorm:"type:varchar(32);NOT NULL"`
	// 钱包地址
	Address string `json:"address" gorm:"type:varchar(64);NOT NULL"`
	// 密钥
	Password string `json:"password" gorm:"type:varchar(512);NOT NULL"`
	// 占用标志 < 0 为系统作为热钱包使用 0=未占用, >0用户冲币地址占用,关联商户id。
	UseTo int64 `json:"use_to" gorm:"type:int(11);NOT NULL"`
}

func (table *ModelAddress) TableName() string {
	return "address"
}
