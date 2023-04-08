// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameShipPerformanceInfo = "t_ship_performance_info"

// ShipPerformanceInfo mapped from table <t_ship_performance_info>
type ShipPerformanceInfo struct {
	ID                int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 自增主键
	PerformanceNo     int64     `gorm:"column:performance_no;not null" json:"performance_no"`                     // 履约编号
	BuyerID           int64     `gorm:"column:buyer_id;not null" json:"buyer_id"`                                 // 买方id
	Status            int32     `gorm:"column:status;not null" json:"status"`                                     // 履约状态
	Remark            string    `gorm:"column:remark" json:"remark"`                                              // 备注
	CreateTime        time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	ModifiedTime      time.Time `gorm:"column:modified_time" json:"modified_time"`                                // 修改时间
	OrderID           int64     `gorm:"column:order_id;not null" json:"order_id"`                                 // 订单id
	BuyerCompanyID    int64     `gorm:"column:buyer_company_id;not null" json:"buyer_company_id"`                 // 买方公司Id
	BuyerCompanyName  string    `gorm:"column:buyer_company_name;not null" json:"buyer_company_name"`             // 买方公司名称
	SellerCompanyID   int64     `gorm:"column:seller_company_id;not null" json:"seller_company_id"`               // 卖方公司id
	SellerCompanyName string    `gorm:"column:seller_company_name;not null" json:"seller_company_name"`           // 卖方公司名称
	BuyerName         string    `gorm:"column:buyer_name;not null" json:"buyer_name"`                             // 买方名称
	PayStatus         int32     `gorm:"column:pay_status;not null;default:1" json:"pay_status"`                   // 付款状态1未付款2部分付款3全付款
	OrderNo           string    `gorm:"column:order_no;not null" json:"order_no"`                                 // 订单编号
	SellerID          int64     `gorm:"column:seller_id" json:"seller_id"`                                        // 买家id
	SellerName        string    `gorm:"column:seller_name" json:"seller_name"`                                    // 买家名称
}

// TableName ShipPerformanceInfo's table name
func (*ShipPerformanceInfo) TableName() string {
	return TableNameShipPerformanceInfo
}