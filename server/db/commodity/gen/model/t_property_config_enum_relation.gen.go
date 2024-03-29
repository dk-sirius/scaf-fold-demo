// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePropertyConfigEnumRelation = "t_property_config_enum_relation"

// PropertyConfigEnumRelation mapped from table <t_property_config_enum_relation>
type PropertyConfigEnumRelation struct {
	FID        int64     `gorm:"column:f_id;primaryKey" json:"f_id"`                                             // 主键
	FTenantID  string    `gorm:"column:f_tenant_id" json:"f_tenant_id"`                                          // 租户id
	FGroupID   int64     `gorm:"column:f_group_id" json:"f_group_id"`                                            // 数据隔离组织ID
	FConfigID  int64     `gorm:"column:f_config_id;not null" json:"f_config_id"`                                 // 属性配置id
	FEnumID    int64     `gorm:"column:f_enum_id;not null" json:"f_enum_id"`                                     // 属性枚举id
	FStatus    int16     `gorm:"column:f_status;not null;default:1" json:"f_status"`                             // 是否启用 unknown：0，是：1 ,否：2
	FCreatedBy string    `gorm:"column:f_created_by;not null;default:''::character varying" json:"f_created_by"` // 创建人
	FUpdatedBy string    `gorm:"column:f_updated_by;not null;default:''::character varying" json:"f_updated_by"` // 修改人
	FCreatedAt time.Time `gorm:"column:f_created_at;not null" json:"f_created_at"`                               // 创建时间
	FUpdatedAt time.Time `gorm:"column:f_updated_at;not null" json:"f_updated_at"`                               // 最后更新时间
	FVersion   int64     `gorm:"column:f_version;not null" json:"f_version"`                                     // 乐观锁
	FRemark    string    `gorm:"column:f_remark;not null;default:''::character varying" json:"f_remark"`         // 备注
	FDeleted   int16     `gorm:"column:f_deleted;not null" json:"f_deleted"`                                     // 删除标记
}

// TableName PropertyConfigEnumRelation's table name
func (*PropertyConfigEnumRelation) TableName() string {
	return TableNamePropertyConfigEnumRelation
}
