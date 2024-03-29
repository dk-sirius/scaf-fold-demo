// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBail = "t_bail"

// Bail mapped from table <t_bail>
type Bail struct {
	FID           int64     `gorm:"column:f_id;primaryKey" json:"f_id"`                     // 自增主键
	FCreatedAt    time.Time `gorm:"column:f_created_at;not null" json:"f_created_at"`       // 创建时间
	FUpdatedAt    time.Time `gorm:"column:f_updated_at;not null" json:"f_updated_at"`       // 更新时间
	FDeleted      int16     `gorm:"column:f_deleted;not null" json:"f_deleted"`             // 删除：0-未删除，1-删除
	FCurrency     float64   `gorm:"column:f_currency;not null" json:"f_currency"`           // 保证金
	FCurrencyType int32     `gorm:"column:f_currency_type;not null" json:"f_currency_type"` // 货币类型
	FEnable       int32     `gorm:"column:f_enable;not null" json:"f_enable"`               // 保证金是否生效，0-生效，1-不生效
	FBailID       int64     `gorm:"column:f_bail_id;not null" json:"f_bail_id"`             // 保证金id
}

// TableName Bail's table name
func (*Bail) TableName() string {
	return TableNameBail
}
