// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePropertyConfig = "t_property_config"

// PropertyConfig mapped from table <t_property_config>
type PropertyConfig struct {
	FID           int64     `gorm:"column:f_id;primaryKey" json:"f_id"`                                                   // 主键
	FTenantID     string    `gorm:"column:f_tenant_id" json:"f_tenant_id"`                                                // 租户id
	FGroupID      int64     `gorm:"column:f_group_id" json:"f_group_id"`                                                  // 数据隔离组织ID
	FDescription  string    `gorm:"column:f_description;not null;default:''::character varying" json:"f_description"`     // 描述
	FStatus       int16     `gorm:"column:f_status;not null;default:1" json:"f_status"`                                   // 是否启用 unknown：0，是：1 ,否：2
	FCreatedBy    string    `gorm:"column:f_created_by;not null;default:''::character varying" json:"f_created_by"`       // 创建人
	FUpdatedBy    string    `gorm:"column:f_updated_by;not null;default:''::character varying" json:"f_updated_by"`       // 修改人
	FCreatedAt    time.Time `gorm:"column:f_created_at;not null" json:"f_created_at"`                                     // 创建时间
	FUpdatedAt    time.Time `gorm:"column:f_updated_at;not null" json:"f_updated_at"`                                     // 最后更新时间
	FVersion      int64     `gorm:"column:f_version;not null" json:"f_version"`                                           // 乐观锁
	FRemark       string    `gorm:"column:f_remark;not null;default:''::character varying" json:"f_remark"`               // 备注
	FDeleted      int16     `gorm:"column:f_deleted;not null" json:"f_deleted"`                                           // 删除标记
	FCategoryID   int64     `gorm:"column:f_category_id;not null" json:"f_category_id"`                                   // 分类id
	FPropertyID   int64     `gorm:"column:f_property_id;not null" json:"f_property_id"`                                   // propertyId
	FInputType    int16     `gorm:"column:f_input_type;not null" json:"f_input_type"`                                     // 填写方式：unknown：0，单选：1，多选：2，填写：3
	FRequired     int16     `gorm:"column:f_required;not null" json:"f_required"`                                         // 是否必填：unknown：0，是：1 ,否：2
	FUseType      int16     `gorm:"column:f_use_type;not null" json:"f_use_type"`                                         // 用途类型：unknown：0，基本属性：1，营销属性：2，售后属性：3，描述属性：4
	FShow         int16     `gorm:"column:f_show;not null" json:"f_show"`                                                 // 是否展示：unknown：0，是：1 ,否：2
	FAlias        string    `gorm:"column:f_alias;not null;default:''::character varying" json:"f_alias"`                 // 别名
	FConfigType   int16     `gorm:"column:f_config_type;not null" json:"f_config_type"`                                   // 配置类型：unknown：0，sku：1，spu：2
	FPropertyCode string    `gorm:"column:f_property_code;not null;default:''::character varying" json:"f_property_code"` // 属性code
}

// TableName PropertyConfig's table name
func (*PropertyConfig) TableName() string {
	return TableNamePropertyConfig
}
