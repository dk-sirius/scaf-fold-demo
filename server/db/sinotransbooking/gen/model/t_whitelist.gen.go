// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameWhitelist = "t_whitelist"

// Whitelist mapped from table <t_whitelist>
type Whitelist struct {
	FID           int64     `gorm:"column:f_id;primaryKey" json:"f_id"`                     // 自增主键
	FCreatedAt    time.Time `gorm:"column:f_created_at;not null" json:"f_created_at"`       // 创建时间
	FUpdatedAt    time.Time `gorm:"column:f_updated_at;not null" json:"f_updated_at"`       // 更新时间
	FDeleted      int16     `gorm:"column:f_deleted;not null" json:"f_deleted"`             // 逻辑软删除 0-未删除，1-删除
	FEnable       int32     `gorm:"column:f_enable;not null" json:"f_enable"`               // 启动：0-启动，2-取消
	FEnterpriseID int64     `gorm:"column:f_enterprise_id;not null" json:"f_enterprise_id"` // 企业id
}

// TableName Whitelist's table name
func (*Whitelist) TableName() string {
	return TableNameWhitelist
}
