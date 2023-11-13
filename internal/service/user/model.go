package user

import "gorm.io/gorm"

type ModelUser struct {
	gorm.Model
	// 角色
	Role int `json:"role" gorm:"type:int(11);NOT NULL"`
	// 账号
	Username string `json:"username" gorm:"type:varchar(64);NOT NULL"`
	// 密码
	Password string `json:"password" gorm:"type:varchar(128);NOT NULL"`
	// 校验码
	AuthGoogle string `json:"auth_google" gorm:"type:varchar(64);NOT NULL"`
	// token
	AuthToken string `json:"auth_token" gorm:"type:varchar(64);NOT NULL"`
	// setting
	Setting string `json:"setting" gorm:"type:varchar(1024);NOT NULL"`
}

func (table *ModelUser) TableName() string {
	return "user"
}
