// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePropertyValue = "t_property_value"

// PropertyValue mapped from table <t_property_value>
type PropertyValue struct {
	FID             int64     `gorm:"column:f_id;primaryKey" json:"f_id"`                       // 主键
	FTenantID       string    `gorm:"column:f_tenant_id" json:"f_tenant_id"`                    // 租户id
	FRelationID     int64     `gorm:"column:f_relation_id;not null" json:"f_relation_id"`       // 关联id（sku、spu）
	FPropertyID     int64     `gorm:"column:f_property_id;not null" json:"f_property_id"`       // 商品关联属性id
	FPropertyEnumID int64     `gorm:"column:f_property_enum_id" json:"f_property_enum_id"`      // 关联属性枚举ID
	FPropertyValue  string    `gorm:"column:f_property_value;not null" json:"f_property_value"` // 属性值
	FSort           int32     `gorm:"column:f_sort" json:"f_sort"`                              // 属性排序
	FDeleted        int16     `gorm:"column:f_deleted;not null" json:"f_deleted"`               // 删除标记
	FCreatedBy      string    `gorm:"column:f_created_by" json:"f_created_by"`                  // 创建者
	FUpdatedBy      string    `gorm:"column:f_updated_by" json:"f_updated_by"`                  // 修改者
	FCreatedAt      time.Time `gorm:"column:f_created_at" json:"f_created_at"`                  // 创建时间
	FUpdatedAt      time.Time `gorm:"column:f_updated_at" json:"f_updated_at"`                  // 修改时间
	FVersion        int64     `gorm:"column:f_version" json:"f_version"`                        // 乐观锁版本号
	FPropertyCode   string    `gorm:"column:f_property_code;not null" json:"f_property_code"`   // 商品code
	FStatus         int16     `gorm:"column:f_status;not null;default:1" json:"f_status"`       // 是否启用 unknown：0，是：1 ,否：2
	FType           int16     `gorm:"column:f_type;not null;default:2" json:"f_type"`           // 关联类型 unknown：0，sku：1 ,spu：2
}

// TableName PropertyValue's table name
func (*PropertyValue) TableName() string {
	return TableNamePropertyValue
}
