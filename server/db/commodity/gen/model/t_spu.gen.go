// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSpu = "t_spu"

// Spu mapped from table <t_spu>
type Spu struct {
	FID          int64     `gorm:"column:f_id;primaryKey" json:"f_id"`               // 主键
	FTenantID    string    `gorm:"column:f_tenant_id" json:"f_tenant_id"`            // 租户id
	FGroupID     int64     `gorm:"column:f_group_id" json:"f_group_id"`              // 数据隔离组织ID
	FCode        string    `gorm:"column:f_code" json:"f_code"`                      // 编码
	FName        string    `gorm:"column:f_name" json:"f_name"`                      // 名称
	FDescription string    `gorm:"column:f_description" json:"f_description"`        // 描述
	FStatus      int16     `gorm:"column:f_status;default:1" json:"f_status"`        // 是否启用 unknown：0，是：1 ,否：2
	FCreatedBy   string    `gorm:"column:f_created_by" json:"f_created_by"`          // 创建人
	FUpdatedBy   string    `gorm:"column:f_updated_by" json:"f_updated_by"`          // 修改人
	FCreatedAt   time.Time `gorm:"column:f_created_at;not null" json:"f_created_at"` // 创建时间
	FUpdatedAt   time.Time `gorm:"column:f_updated_at;not null" json:"f_updated_at"` // 最后更新时间
	FVersion     int64     `gorm:"column:f_version;not null" json:"f_version"`       // 乐观锁
	FRemark      string    `gorm:"column:f_remark" json:"f_remark"`                  // 备注
	FDeleted     int16     `gorm:"column:f_deleted;not null" json:"f_deleted"`       // 删除标记
	FCategoryID  int64     `gorm:"column:f_category_id" json:"f_category_id"`        // 分类id
}

// TableName Spu's table name
func (*Spu) TableName() string {
	return TableNameSpu
}
