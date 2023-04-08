// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRailwayCompany = "t_railway_company"

// RailwayCompany mapped from table <t_railway_company>
type RailwayCompany struct {
	FID             int64     `gorm:"column:f_id;primaryKey" json:"f_id"`                 // 主键
	FStatus         int16     `gorm:"column:f_status;not null;default:1" json:"f_status"` // 是否启用unknown：0，是：1 ,否：2
	FCreatedBy      string    `gorm:"column:f_created_by" json:"f_created_by"`            // 创建人
	FUpdatedBy      string    `gorm:"column:f_updated_by" json:"f_updated_by"`            // 修改人
	FCreatedAt      time.Time `gorm:"column:f_created_at;not null" json:"f_created_at"`   // 创建时间
	FUpdatedAt      time.Time `gorm:"column:f_updated_at;not null" json:"f_updated_at"`   // 最后更新时间
	FVersion        int64     `gorm:"column:f_version;not null" json:"f_version"`         // 乐观锁
	FRemark         string    `gorm:"column:f_remark" json:"f_remark"`                    // 备注
	FDeleted        int16     `gorm:"column:f_deleted;not null" json:"f_deleted"`         // 删除标记 1：删除 0：未删除
	FNameZh         string    `gorm:"column:f_name_zh;not null" json:"f_name_zh"`         // 中文名
	FTenantID       int64     `gorm:"column:f_tenant_id" json:"f_tenant_id"`
	FNameEn         string    `gorm:"column:f_name_en" json:"f_name_en"`                                                        // 英文名
	FAddressID      int64     `gorm:"column:f_address_id;not null" json:"f_address_id"`                                         // 地址id
	FAbbreviationZh string    `gorm:"column:f_abbreviation_zh;not null;default:''::character varying" json:"f_abbreviation_zh"` // 中文简称
	FAbbreviationEn string    `gorm:"column:f_abbreviation_en;not null;default:''::character varying" json:"f_abbreviation_en"` // 英文简称
}

// TableName RailwayCompany's table name
func (*RailwayCompany) TableName() string {
	return TableNameRailwayCompany
}
