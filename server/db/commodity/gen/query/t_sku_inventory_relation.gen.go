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

func newSkuInventoryRelation(db *gorm.DB, opts ...gen.DOOption) skuInventoryRelation {
	_skuInventoryRelation := skuInventoryRelation{}

	_skuInventoryRelation.skuInventoryRelationDo.UseDB(db, opts...)
	_skuInventoryRelation.skuInventoryRelationDo.UseModel(&model.SkuInventoryRelation{})

	tableName := _skuInventoryRelation.skuInventoryRelationDo.TableName()
	_skuInventoryRelation.ALL = field.NewAsterisk(tableName)
	_skuInventoryRelation.FID = field.NewInt64(tableName, "f_id")
	_skuInventoryRelation.FTenantID = field.NewString(tableName, "f_tenant_id")
	_skuInventoryRelation.FGroupID = field.NewInt64(tableName, "f_group_id")
	_skuInventoryRelation.FSkuID = field.NewInt64(tableName, "f_sku_id")
	_skuInventoryRelation.FInventoryID = field.NewInt64(tableName, "f_inventory_id")
	_skuInventoryRelation.FStatus = field.NewInt16(tableName, "f_status")
	_skuInventoryRelation.FCreatedBy = field.NewString(tableName, "f_created_by")
	_skuInventoryRelation.FUpdatedBy = field.NewString(tableName, "f_updated_by")
	_skuInventoryRelation.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_skuInventoryRelation.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_skuInventoryRelation.FVersion = field.NewInt64(tableName, "f_version")
	_skuInventoryRelation.FRemark = field.NewString(tableName, "f_remark")
	_skuInventoryRelation.FDeleted = field.NewInt16(tableName, "f_deleted")

	_skuInventoryRelation.fillFieldMap()

	return _skuInventoryRelation
}

type skuInventoryRelation struct {
	skuInventoryRelationDo skuInventoryRelationDo

	ALL          field.Asterisk
	FID          field.Int64  // 主键
	FTenantID    field.String // 租户id
	FGroupID     field.Int64  // 数据隔离组织ID
	FSkuID       field.Int64  // 商品id
	FInventoryID field.Int64  // 库存id
	FStatus      field.Int16  // 是否启用 unknown：0，是：1 ,否：2
	FCreatedBy   field.String // 创建人
	FUpdatedBy   field.String // 修改人
	FCreatedAt   field.Time   // 创建时间
	FUpdatedAt   field.Time   // 最后更新时间
	FVersion     field.Int64  // 乐观锁
	FRemark      field.String // 备注
	FDeleted     field.Int16  // 删除标记

	fieldMap map[string]field.Expr
}

func (s skuInventoryRelation) Table(newTableName string) *skuInventoryRelation {
	s.skuInventoryRelationDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skuInventoryRelation) As(alias string) *skuInventoryRelation {
	s.skuInventoryRelationDo.DO = *(s.skuInventoryRelationDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skuInventoryRelation) updateTableName(table string) *skuInventoryRelation {
	s.ALL = field.NewAsterisk(table)
	s.FID = field.NewInt64(table, "f_id")
	s.FTenantID = field.NewString(table, "f_tenant_id")
	s.FGroupID = field.NewInt64(table, "f_group_id")
	s.FSkuID = field.NewInt64(table, "f_sku_id")
	s.FInventoryID = field.NewInt64(table, "f_inventory_id")
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

func (s *skuInventoryRelation) WithContext(ctx context.Context) *skuInventoryRelationDo {
	return s.skuInventoryRelationDo.WithContext(ctx)
}

func (s skuInventoryRelation) TableName() string { return s.skuInventoryRelationDo.TableName() }

func (s skuInventoryRelation) Alias() string { return s.skuInventoryRelationDo.Alias() }

func (s *skuInventoryRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skuInventoryRelation) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 13)
	s.fieldMap["f_id"] = s.FID
	s.fieldMap["f_tenant_id"] = s.FTenantID
	s.fieldMap["f_group_id"] = s.FGroupID
	s.fieldMap["f_sku_id"] = s.FSkuID
	s.fieldMap["f_inventory_id"] = s.FInventoryID
	s.fieldMap["f_status"] = s.FStatus
	s.fieldMap["f_created_by"] = s.FCreatedBy
	s.fieldMap["f_updated_by"] = s.FUpdatedBy
	s.fieldMap["f_created_at"] = s.FCreatedAt
	s.fieldMap["f_updated_at"] = s.FUpdatedAt
	s.fieldMap["f_version"] = s.FVersion
	s.fieldMap["f_remark"] = s.FRemark
	s.fieldMap["f_deleted"] = s.FDeleted
}

func (s skuInventoryRelation) clone(db *gorm.DB) skuInventoryRelation {
	s.skuInventoryRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skuInventoryRelation) replaceDB(db *gorm.DB) skuInventoryRelation {
	s.skuInventoryRelationDo.ReplaceDB(db)
	return s
}

type skuInventoryRelationDo struct{ gen.DO }

func (s skuInventoryRelationDo) Debug() *skuInventoryRelationDo {
	return s.withDO(s.DO.Debug())
}

func (s skuInventoryRelationDo) WithContext(ctx context.Context) *skuInventoryRelationDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skuInventoryRelationDo) ReadDB() *skuInventoryRelationDo {
	return s.Clauses(dbresolver.Read)
}

func (s skuInventoryRelationDo) WriteDB() *skuInventoryRelationDo {
	return s.Clauses(dbresolver.Write)
}

func (s skuInventoryRelationDo) Session(config *gorm.Session) *skuInventoryRelationDo {
	return s.withDO(s.DO.Session(config))
}

func (s skuInventoryRelationDo) Clauses(conds ...clause.Expression) *skuInventoryRelationDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skuInventoryRelationDo) Returning(value interface{}, columns ...string) *skuInventoryRelationDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skuInventoryRelationDo) Not(conds ...gen.Condition) *skuInventoryRelationDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skuInventoryRelationDo) Or(conds ...gen.Condition) *skuInventoryRelationDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skuInventoryRelationDo) Select(conds ...field.Expr) *skuInventoryRelationDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skuInventoryRelationDo) Where(conds ...gen.Condition) *skuInventoryRelationDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skuInventoryRelationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *skuInventoryRelationDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s skuInventoryRelationDo) Order(conds ...field.Expr) *skuInventoryRelationDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skuInventoryRelationDo) Distinct(cols ...field.Expr) *skuInventoryRelationDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skuInventoryRelationDo) Omit(cols ...field.Expr) *skuInventoryRelationDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skuInventoryRelationDo) Join(table schema.Tabler, on ...field.Expr) *skuInventoryRelationDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skuInventoryRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skuInventoryRelationDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skuInventoryRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) *skuInventoryRelationDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skuInventoryRelationDo) Group(cols ...field.Expr) *skuInventoryRelationDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skuInventoryRelationDo) Having(conds ...gen.Condition) *skuInventoryRelationDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skuInventoryRelationDo) Limit(limit int) *skuInventoryRelationDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skuInventoryRelationDo) Offset(offset int) *skuInventoryRelationDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skuInventoryRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skuInventoryRelationDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skuInventoryRelationDo) Unscoped() *skuInventoryRelationDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skuInventoryRelationDo) Create(values ...*model.SkuInventoryRelation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skuInventoryRelationDo) CreateInBatches(values []*model.SkuInventoryRelation, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skuInventoryRelationDo) Save(values ...*model.SkuInventoryRelation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skuInventoryRelationDo) First() (*model.SkuInventoryRelation, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryRelation), nil
	}
}

func (s skuInventoryRelationDo) Take() (*model.SkuInventoryRelation, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryRelation), nil
	}
}

func (s skuInventoryRelationDo) Last() (*model.SkuInventoryRelation, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryRelation), nil
	}
}

func (s skuInventoryRelationDo) Find() ([]*model.SkuInventoryRelation, error) {
	result, err := s.DO.Find()
	return result.([]*model.SkuInventoryRelation), err
}

func (s skuInventoryRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SkuInventoryRelation, err error) {
	buf := make([]*model.SkuInventoryRelation, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skuInventoryRelationDo) FindInBatches(result *[]*model.SkuInventoryRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skuInventoryRelationDo) Attrs(attrs ...field.AssignExpr) *skuInventoryRelationDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skuInventoryRelationDo) Assign(attrs ...field.AssignExpr) *skuInventoryRelationDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skuInventoryRelationDo) Joins(fields ...field.RelationField) *skuInventoryRelationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skuInventoryRelationDo) Preload(fields ...field.RelationField) *skuInventoryRelationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skuInventoryRelationDo) FirstOrInit() (*model.SkuInventoryRelation, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryRelation), nil
	}
}

func (s skuInventoryRelationDo) FirstOrCreate() (*model.SkuInventoryRelation, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryRelation), nil
	}
}

func (s skuInventoryRelationDo) FindByPage(offset int, limit int) (result []*model.SkuInventoryRelation, count int64, err error) {
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

func (s skuInventoryRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skuInventoryRelationDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skuInventoryRelationDo) Delete(models ...*model.SkuInventoryRelation) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skuInventoryRelationDo) withDO(do gen.Dao) *skuInventoryRelationDo {
	s.DO = *do.(*gen.DO)
	return s
}