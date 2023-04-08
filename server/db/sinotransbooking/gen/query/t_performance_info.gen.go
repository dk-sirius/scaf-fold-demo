// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/dk-sirius/scaf-fold-demo/server/db/sinotransbooking/gen/model"
)

func newPerformanceInfo(db *gorm.DB, opts ...gen.DOOption) performanceInfo {
	_performanceInfo := performanceInfo{}

	_performanceInfo.performanceInfoDo.UseDB(db, opts...)
	_performanceInfo.performanceInfoDo.UseModel(&model.PerformanceInfo{})

	tableName := _performanceInfo.performanceInfoDo.TableName()
	_performanceInfo.ALL = field.NewAsterisk(tableName)
	_performanceInfo.FID = field.NewInt64(tableName, "f_id")
	_performanceInfo.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_performanceInfo.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_performanceInfo.FDeletedAt = field.NewTime(tableName, "f_deleted_at")
	_performanceInfo.FDeleted = field.NewInt16(tableName, "f_deleted")
	_performanceInfo.FOrderID = field.NewInt64(tableName, "f_order_id")
	_performanceInfo.FOrderNo = field.NewString(tableName, "f_order_no")
	_performanceInfo.FPerformanceID = field.NewInt64(tableName, "f_performance_id")
	_performanceInfo.FBuyerCompanyID = field.NewInt64(tableName, "f_buyer_company_id")
	_performanceInfo.FEnterpriseID = field.NewInt64(tableName, "f_enterprise_id")
	_performanceInfo.FCommodityItemID = field.NewInt64(tableName, "f_commodity_item_id")
	_performanceInfo.FCommodityItemVersionID = field.NewInt64(tableName, "f_commodity_item_version_id")
	_performanceInfo.FStatus = field.NewInt32(tableName, "f_status")
	_performanceInfo.FBuyerBusinessName = field.NewString(tableName, "f_buyer_business_name")
	_performanceInfo.FBusinessName = field.NewString(tableName, "f_business_name")
	_performanceInfo.FPayStatus = field.NewInt16(tableName, "f_pay_status")

	_performanceInfo.fillFieldMap()

	return _performanceInfo
}

type performanceInfo struct {
	performanceInfoDo performanceInfoDo

	ALL                     field.Asterisk
	FID                     field.Int64  // 自增id
	FCreatedAt              field.Time   // 创建时间
	FUpdatedAt              field.Time   // 更新时间
	FDeletedAt              field.Time   // 删除时间
	FDeleted                field.Int16  // 删除标识0-未删除，1-删除
	FOrderID                field.Int64  // 订单id
	FOrderNo                field.String // 订单号
	FPerformanceID          field.Int64  // 履约单id
	FBuyerCompanyID         field.Int64  // 买家企业id
	FEnterpriseID           field.Int64  // 卖家企业id
	FCommodityItemID        field.Int64  // 商品item id
	FCommodityItemVersionID field.Int64  // 商品版本id
	FStatus                 field.Int32  // 履约单状态
	FBuyerBusinessName      field.String // 买家企业用户名称
	FBusinessName           field.String // 卖家企业用户名称
	FPayStatus              field.Int16  // 履约单支付状态    1;//处理中2;//支付成功 3;//支付失败4;//支付关闭

	fieldMap map[string]field.Expr
}

func (p performanceInfo) Table(newTableName string) *performanceInfo {
	p.performanceInfoDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p performanceInfo) As(alias string) *performanceInfo {
	p.performanceInfoDo.DO = *(p.performanceInfoDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *performanceInfo) updateTableName(table string) *performanceInfo {
	p.ALL = field.NewAsterisk(table)
	p.FID = field.NewInt64(table, "f_id")
	p.FCreatedAt = field.NewTime(table, "f_created_at")
	p.FUpdatedAt = field.NewTime(table, "f_updated_at")
	p.FDeletedAt = field.NewTime(table, "f_deleted_at")
	p.FDeleted = field.NewInt16(table, "f_deleted")
	p.FOrderID = field.NewInt64(table, "f_order_id")
	p.FOrderNo = field.NewString(table, "f_order_no")
	p.FPerformanceID = field.NewInt64(table, "f_performance_id")
	p.FBuyerCompanyID = field.NewInt64(table, "f_buyer_company_id")
	p.FEnterpriseID = field.NewInt64(table, "f_enterprise_id")
	p.FCommodityItemID = field.NewInt64(table, "f_commodity_item_id")
	p.FCommodityItemVersionID = field.NewInt64(table, "f_commodity_item_version_id")
	p.FStatus = field.NewInt32(table, "f_status")
	p.FBuyerBusinessName = field.NewString(table, "f_buyer_business_name")
	p.FBusinessName = field.NewString(table, "f_business_name")
	p.FPayStatus = field.NewInt16(table, "f_pay_status")

	p.fillFieldMap()

	return p
}

func (p *performanceInfo) WithContext(ctx context.Context) *performanceInfoDo {
	return p.performanceInfoDo.WithContext(ctx)
}

func (p performanceInfo) TableName() string { return p.performanceInfoDo.TableName() }

func (p performanceInfo) Alias() string { return p.performanceInfoDo.Alias() }

func (p *performanceInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *performanceInfo) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 16)
	p.fieldMap["f_id"] = p.FID
	p.fieldMap["f_created_at"] = p.FCreatedAt
	p.fieldMap["f_updated_at"] = p.FUpdatedAt
	p.fieldMap["f_deleted_at"] = p.FDeletedAt
	p.fieldMap["f_deleted"] = p.FDeleted
	p.fieldMap["f_order_id"] = p.FOrderID
	p.fieldMap["f_order_no"] = p.FOrderNo
	p.fieldMap["f_performance_id"] = p.FPerformanceID
	p.fieldMap["f_buyer_company_id"] = p.FBuyerCompanyID
	p.fieldMap["f_enterprise_id"] = p.FEnterpriseID
	p.fieldMap["f_commodity_item_id"] = p.FCommodityItemID
	p.fieldMap["f_commodity_item_version_id"] = p.FCommodityItemVersionID
	p.fieldMap["f_status"] = p.FStatus
	p.fieldMap["f_buyer_business_name"] = p.FBuyerBusinessName
	p.fieldMap["f_business_name"] = p.FBusinessName
	p.fieldMap["f_pay_status"] = p.FPayStatus
}

func (p performanceInfo) clone(db *gorm.DB) performanceInfo {
	p.performanceInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p performanceInfo) replaceDB(db *gorm.DB) performanceInfo {
	p.performanceInfoDo.ReplaceDB(db)
	return p
}

type performanceInfoDo struct{ gen.DO }

func (p performanceInfoDo) Debug() *performanceInfoDo {
	return p.withDO(p.DO.Debug())
}

func (p performanceInfoDo) WithContext(ctx context.Context) *performanceInfoDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p performanceInfoDo) ReadDB() *performanceInfoDo {
	return p.Clauses(dbresolver.Read)
}

func (p performanceInfoDo) WriteDB() *performanceInfoDo {
	return p.Clauses(dbresolver.Write)
}

func (p performanceInfoDo) Session(config *gorm.Session) *performanceInfoDo {
	return p.withDO(p.DO.Session(config))
}

func (p performanceInfoDo) Clauses(conds ...clause.Expression) *performanceInfoDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p performanceInfoDo) Returning(value interface{}, columns ...string) *performanceInfoDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p performanceInfoDo) Not(conds ...gen.Condition) *performanceInfoDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p performanceInfoDo) Or(conds ...gen.Condition) *performanceInfoDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p performanceInfoDo) Select(conds ...field.Expr) *performanceInfoDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p performanceInfoDo) Where(conds ...gen.Condition) *performanceInfoDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p performanceInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *performanceInfoDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p performanceInfoDo) Order(conds ...field.Expr) *performanceInfoDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p performanceInfoDo) Distinct(cols ...field.Expr) *performanceInfoDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p performanceInfoDo) Omit(cols ...field.Expr) *performanceInfoDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p performanceInfoDo) Join(table schema.Tabler, on ...field.Expr) *performanceInfoDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p performanceInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *performanceInfoDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p performanceInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *performanceInfoDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p performanceInfoDo) Group(cols ...field.Expr) *performanceInfoDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p performanceInfoDo) Having(conds ...gen.Condition) *performanceInfoDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p performanceInfoDo) Limit(limit int) *performanceInfoDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p performanceInfoDo) Offset(offset int) *performanceInfoDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p performanceInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *performanceInfoDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p performanceInfoDo) Unscoped() *performanceInfoDo {
	return p.withDO(p.DO.Unscoped())
}

func (p performanceInfoDo) Create(values ...*model.PerformanceInfo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p performanceInfoDo) CreateInBatches(values []*model.PerformanceInfo, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p performanceInfoDo) Save(values ...*model.PerformanceInfo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p performanceInfoDo) First() (*model.PerformanceInfo, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PerformanceInfo), nil
	}
}

func (p performanceInfoDo) Take() (*model.PerformanceInfo, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PerformanceInfo), nil
	}
}

func (p performanceInfoDo) Last() (*model.PerformanceInfo, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PerformanceInfo), nil
	}
}

func (p performanceInfoDo) Find() ([]*model.PerformanceInfo, error) {
	result, err := p.DO.Find()
	return result.([]*model.PerformanceInfo), err
}

func (p performanceInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PerformanceInfo, err error) {
	buf := make([]*model.PerformanceInfo, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p performanceInfoDo) FindInBatches(result *[]*model.PerformanceInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p performanceInfoDo) Attrs(attrs ...field.AssignExpr) *performanceInfoDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p performanceInfoDo) Assign(attrs ...field.AssignExpr) *performanceInfoDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p performanceInfoDo) Joins(fields ...field.RelationField) *performanceInfoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p performanceInfoDo) Preload(fields ...field.RelationField) *performanceInfoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p performanceInfoDo) FirstOrInit() (*model.PerformanceInfo, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PerformanceInfo), nil
	}
}

func (p performanceInfoDo) FirstOrCreate() (*model.PerformanceInfo, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PerformanceInfo), nil
	}
}

func (p performanceInfoDo) FindByPage(offset int, limit int) (result []*model.PerformanceInfo, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p performanceInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p performanceInfoDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p performanceInfoDo) Delete(models ...*model.PerformanceInfo) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *performanceInfoDo) withDO(do gen.Dao) *performanceInfoDo {
	p.DO = *do.(*gen.DO)
	return p
}
