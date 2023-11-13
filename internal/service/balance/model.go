package balance

import "gorm.io/gorm"

type ModelBalance struct {
	gorm.Model
	// 关联商户id
	ProductID int64 `json:"product_id" gorm:"type:int(11);NOT NULL"`
	// 币种
	Symbol string `json:"symbol" gorm:"type:varchar(64);NOT NULL"`
	// 余额单位小数点
	Decimals int64 `json:"decimals" gorm:"type:int(11);NOT NULL"`
	// 当前余额 金额用varchar是因为每个表有多条记录，每个币种的单位不一样
	AvailableAmount string `json:"available_amount" gorm:"type:varchar(128);NOT NULL"`
	// 冻结金额
	FreezeAmount string `json:"freeze_amount" gorm:"type:varchar(128);NOT NULL"`
	// 冻结时间 =0默认，>0 时间戳
	FreezeTime int64 `json:"freeze_time" gorm:"type:int(20);DEFAULT 0;NOT NULL"`
	// 解冻时间 =0默认，>0 时间戳
	UnFreezeTime int64 `json:"un_freeze_time" gorm:"type:int(20);DEFAULT 0;NOT NULL"`
	// 充值总额
	DepositTotal string `json:"deposit_total" gorm:"type:varchar(128);NOT NULL"`
	// 提款总额
	WithdrawTotal string `json:"withdraw_total" gorm:"type:varchar(128);NOT NULL"`
	// 充值总次数
	DepositCount int64 `json:"deposit_count" gorm:"type:int(11);DEFAULT 0;NOT NULL"`
	// 提款总次数
	WithdrawCount int64 `json:"withdraw_count" gorm:"type:int(11);DEFAULT 0;NOT NULL"`
	// 账户变更前的数据 json存储
	BeforeData string `json:"before_data" gorm:"type:json;NOT NULL"`
	// 账户变更后的数据 json存储
	AfterData string `json:"after_data" gorm:"type:json;NOT NULL"`
	// 乐观锁版本，交易时，先查询balance表，得到version值，交易成功后比较version+1和version关系，一致则交易成功，变更金额。
	//  1)先读task表的数据（实际上这个表只有一条记录），得到version的值为versionValue
	//  2)每次更新task表中的value字段时，为了防止发生冲突，需要这样操作
	// update xxx_table set value = newValue, version =  versionValue + 1 where version = versionValue;
	// 参考 https://blog.csdn.net/qq_38765404/article/details/82253834
	Version int64 `json:"version" gorm:"type:int(11);NOT NULL"`
}

func (table *ModelBalance) TableName() string {
	return "balance"
}
