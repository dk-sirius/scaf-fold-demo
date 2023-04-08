// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameShippingConditions = "t_shipping_conditions"

// ShippingConditions mapped from table <t_shipping_conditions>
type ShippingConditions struct {
	FID              int64     `gorm:"column:f_id;primaryKey;autoIncrement:true" json:"f_id"`        // 主键自增
	FScacs           string    `gorm:"column:f_scacs" json:"f_scacs"`                                // 查询船公司
	FSearchDate      string    `gorm:"column:f_search_date" json:"f_search_date"`                    // 有效期开始
	FOriginPort      string    `gorm:"column:f_origin_port;not null" json:"f_origin_port"`           // 起始港code
	FDestinationPort string    `gorm:"column:f_destination_port;not null" json:"f_destination_port"` // 目的港code
	FWeeksOut        string    `gorm:"column:f_weeks_out;not null;default:8" json:"f_weeks_out"`     // 查询周数
	FCreatedAt       time.Time `gorm:"column:f_created_at;not null" json:"f_created_at"`             // 创建时间
	FUpdatedAt       time.Time `gorm:"column:f_updated_at;not null" json:"f_updated_at"`             // 更新时间
	FDeleted         int16     `gorm:"column:f_deleted;not null" json:"f_deleted"`                   // 是否已经删除 0 未删除 1已经删除
	FIsTransit       int16     `gorm:"column:f_is_transit" json:"f_is_transit"`                      // 是否中转0  不中转 1 中转 2 全部查
	FDeletedAt       time.Time `gorm:"column:f_deleted_at" json:"f_deleted_at"`
}

// TableName ShippingConditions's table name
func (*ShippingConditions) TableName() string {
	return TableNameShippingConditions
}