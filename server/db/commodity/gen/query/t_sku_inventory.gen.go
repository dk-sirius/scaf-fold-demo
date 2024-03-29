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

	"github.com/dk-sirius/scaf-fold-demo/server/db/commodity/gen/model"
)

func newSkuInventory(db *gorm.DB, opts ...gen.DOOption) skuInventory {
	_skuInventory := skuInventory{}

	_skuInventory.skuInventoryDo.UseDB(db, opts...)
	_skuInventory.skuInventoryDo.UseModel(&model.SkuInventory{})

	tableName := _skuInventory.skuInventoryDo.TableName()
	_skuInventory.ALL = field.NewAsterisk(tableName)
	_skuInventory.FID = field.NewInt64(tableName, "f_id")
	_skuInventory.FTenantID = field.NewString(tableName, "f_tenant_id")
	_skuInventory.FTotalQty = field.NewInt32(tableName, "f_total_qty")
	_skuInventory.FAvailableQty = field.NewInt32(tableName, "f_available_qty")
	_skuInventory.FLockQty = field.NewInt32(tableName, "f_lock_qty")
	_skuInventory.FSaledQty = field.NewInt32(tableName, "f_saled_qty")
	_skuInventory.FStatus = field.NewInt16(tableName, "f_status")
	_skuInventory.FCreatedBy = field.NewString(tableName, "f_created_by")
	_skuInventory.FUpdatedBy = field.NewString(tableName, "f_updated_by")
	_skuInventory.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_skuInventory.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_skuInventory.FVersion = field.NewInt64(tableName, "f_version")
	_skuInventory.FRemark = field.NewString(tableName, "f_remark")
	_skuInventory.FDeleted = field.NewInt16(tableName, "f_deleted")

	_skuInventory.fillFieldMap()

	return _skuInventory
}

type skuInventory struct {
	skuInventoryDo skuInventoryDo

	ALL           field.Asterisk
	FID           field.Int64  // 主键
	FTenantID     field.String // 租户id
	FTotalQty     field.Int32  // 销售总库存数量，实际库存数量+锁定库存数量
	FAvailableQty field.Int32  // 实际库存数量
	FLockQty      field.Int32  // 锁定库存数量
	FSaledQty     field.Int32  // 已销售库存数量
	FStatus       field.Int16  // 是否启用 unknown：0，是：1 ,否：2
	FCreatedBy    field.String // 创建人
	FUpdatedBy    field.String // 修改人
	FCreatedAt    field.Time   // 创建时间
	FUpdatedAt    field.Time   // 最后更新时间
	FVersion      field.Int64  // 乐观锁
	FRemark       field.String // 备注
	FDeleted      field.Int16  // 删除标记

	fieldMap map[string]field.Expr
}

func (s skuInventory) Table(newTableName string) *skuInventory {
	s.skuInventoryDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skuInventory) As(alias string) *skuInventory {
	s.skuInventoryDo.DO = *(s.skuInventoryDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skuInventory) updateTableName(table string) *skuInventory {
	s.ALL = field.NewAsterisk(table)
	s.FID = field.NewInt64(table, "f_id")
	s.FTenantID = field.NewString(table, "f_tenant_id")
	s.FTotalQty = field.NewInt32(table, "f_total_qty")
	s.FAvailableQty = field.NewInt32(table, "f_available_qty")
	s.FLockQty = field.NewInt32(table, "f_lock_qty")
	s.FSaledQty = field.NewInt32(table, "f_saled_qty")
	s.FStatus = field.NewInt16(table, "f_status")
	s.FCreatedBy = field.NewString(table, "f_created_by")
	s.FUpdatedBy = field.NewString(table, "f_updated_by")
	s.FCreatedAt = field.NewTime(table, "f_created_at")
	s.FUpdatedAt = field.NewTime(table, "f_updated_at")
	s.FVersion = field.NewInt64(table, "f_version")
	s.FRemark = field.NewString(table, "f_remark")
	s.FDeleted = field.NewInt16(table, "f_deleted")

	s.fillFieldMap()

	return s
}

func (s *skuInventory) WithContext(ctx context.Context) *skuInventoryDo {
	return s.skuInventoryDo.WithContext(ctx)
}

func (s skuInventory) TableName() string { return s.skuInventoryDo.TableName() }

func (s skuInventory) Alias() string { return s.skuInventoryDo.Alias() }

func (s *skuInventory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skuInventory) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["f_id"] = s.FID
	s.fieldMap["f_tenant_id"] = s.FTenantID
	s.fieldMap["f_total_qty"] = s.FTotalQty
	s.fieldMap["f_available_qty"] = s.FAvailableQty
	s.fieldMap["f_lock_qty"] = s.FLockQty
	s.fieldMap["f_saled_qty"] = s.FSaledQty
	s.fieldMap["f_status"] = s.FStatus
	s.fieldMap["f_created_by"] = s.FCreatedBy
	s.fieldMap["f_updated_by"] = s.FUpdatedBy
	s.fieldMap["f_created_at"] = s.FCreatedAt
	s.fieldMap["f_updated_at"] = s.FUpdatedAt
	s.fieldMap["f_version"] = s.FVersion
	s.fieldMap["f_remark"] = s.FRemark
	s.fieldMap["f_deleted"] = s.FDeleted
}

func (s skuInventory) clone(db *gorm.DB) skuInventory {
	s.skuInventoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skuInventory) replaceDB(db *gorm.DB) skuInventory {
	s.skuInventoryDo.ReplaceDB(db)
	return s
}

type skuInventoryDo struct{ gen.DO }

func (s skuInventoryDo) Debug() *skuInventoryDo {
	return s.withDO(s.DO.Debug())
}

func (s skuInventoryDo) WithContext(ctx context.Context) *skuInventoryDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skuInventoryDo) ReadDB() *skuInventoryDo {
	return s.Clauses(dbresolver.Read)
}

func (s skuInventoryDo) WriteDB() *skuInventoryDo {
	return s.Clauses(dbresolver.Write)
}

func (s skuInventoryDo) Session(config *gorm.Session) *skuInventoryDo {
	return s.withDO(s.DO.Session(config))
}

func (s skuInventoryDo) Clauses(conds ...clause.Expression) *skuInventoryDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skuInventoryDo) Returning(value interface{}, columns ...string) *skuInventoryDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skuInventoryDo) Not(conds ...gen.Condition) *skuInventoryDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skuInventoryDo) Or(conds ...gen.Condition) *skuInventoryDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skuInventoryDo) Select(conds ...field.Expr) *skuInventoryDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skuInventoryDo) Where(conds ...gen.Condition) *skuInventoryDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skuInventoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *skuInventoryDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s skuInventoryDo) Order(conds ...field.Expr) *skuInventoryDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skuInventoryDo) Distinct(cols ...field.Expr) *skuInventoryDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skuInventoryDo) Omit(cols ...field.Expr) *skuInventoryDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skuInventoryDo) Join(table schema.Tabler, on ...field.Expr) *skuInventoryDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skuInventoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skuInventoryDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skuInventoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *skuInventoryDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skuInventoryDo) Group(cols ...field.Expr) *skuInventoryDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skuInventoryDo) Having(conds ...gen.Condition) *skuInventoryDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skuInventoryDo) Limit(limit int) *skuInventoryDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skuInventoryDo) Offset(offset int) *skuInventoryDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skuInventoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skuInventoryDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skuInventoryDo) Unscoped() *skuInventoryDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skuInventoryDo) Create(values ...*model.SkuInventory) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skuInventoryDo) CreateInBatches(values []*model.SkuInventory, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skuInventoryDo) Save(values ...*model.SkuInventory) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skuInventoryDo) First() (*model.SkuInventory, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventory), nil
	}
}

func (s skuInventoryDo) Take() (*model.SkuInventory, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventory), nil
	}
}

func (s skuInventoryDo) Last() (*model.SkuInventory, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventory), nil
	}
}

func (s skuInventoryDo) Find() ([]*model.SkuInventory, error) {
	result, err := s.DO.Find()
	return result.([]*model.SkuInventory), err
}

func (s skuInventoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SkuInventory, err error) {
	buf := make([]*model.SkuInventory, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skuInventoryDo) FindInBatches(result *[]*model.SkuInventory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skuInventoryDo) Attrs(attrs ...field.AssignExpr) *skuInventoryDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skuInventoryDo) Assign(attrs ...field.AssignExpr) *skuInventoryDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skuInventoryDo) Joins(fields ...field.RelationField) *skuInventoryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skuInventoryDo) Preload(fields ...field.RelationField) *skuInventoryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skuInventoryDo) FirstOrInit() (*model.SkuInventory, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventory), nil
	}
}

func (s skuInventoryDo) FirstOrCreate() (*model.SkuInventory, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventory), nil
	}
}

func (s skuInventoryDo) FindByPage(offset int, limit int) (result []*model.SkuInventory, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s skuInventoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skuInventoryDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skuInventoryDo) Delete(models ...*model.SkuInventory) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skuInventoryDo) withDO(do gen.Dao) *skuInventoryDo {
	s.DO = *do.(*gen.DO)
	return s
}
