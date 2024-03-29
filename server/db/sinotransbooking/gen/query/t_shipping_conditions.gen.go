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

func newShippingConditions(db *gorm.DB, opts ...gen.DOOption) shippingConditions {
	_shippingConditions := shippingConditions{}

	_shippingConditions.shippingConditionsDo.UseDB(db, opts...)
	_shippingConditions.shippingConditionsDo.UseModel(&model.ShippingConditions{})

	tableName := _shippingConditions.shippingConditionsDo.TableName()
	_shippingConditions.ALL = field.NewAsterisk(tableName)
	_shippingConditions.FID = field.NewInt64(tableName, "f_id")
	_shippingConditions.FScacs = field.NewString(tableName, "f_scacs")
	_shippingConditions.FSearchDate = field.NewString(tableName, "f_search_date")
	_shippingConditions.FOriginPort = field.NewString(tableName, "f_origin_port")
	_shippingConditions.FDestinationPort = field.NewString(tableName, "f_destination_port")
	_shippingConditions.FWeeksOut = field.NewString(tableName, "f_weeks_out")
	_shippingConditions.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_shippingConditions.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_shippingConditions.FDeleted = field.NewInt16(tableName, "f_deleted")
	_shippingConditions.FIsTransit = field.NewInt16(tableName, "f_is_transit")
	_shippingConditions.FDeletedAt = field.NewTime(tableName, "f_deleted_at")

	_shippingConditions.fillFieldMap()

	return _shippingConditions
}

type shippingConditions struct {
	shippingConditionsDo shippingConditionsDo

	ALL              field.Asterisk
	FID              field.Int64  // 主键自增
	FScacs           field.String // 查询船公司
	FSearchDate      field.String // 有效期开始
	FOriginPort      field.String // 起始港code
	FDestinationPort field.String // 目的港code
	FWeeksOut        field.String // 查询周数
	FCreatedAt       field.Time   // 创建时间
	FUpdatedAt       field.Time   // 更新时间
	FDeleted         field.Int16  // 是否已经删除 0 未删除 1已经删除
	FIsTransit       field.Int16  // 是否中转0  不中转 1 中转 2 全部查
	FDeletedAt       field.Time

	fieldMap map[string]field.Expr
}

func (s shippingConditions) Table(newTableName string) *shippingConditions {
	s.shippingConditionsDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s shippingConditions) As(alias string) *shippingConditions {
	s.shippingConditionsDo.DO = *(s.shippingConditionsDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *shippingConditions) updateTableName(table string) *shippingConditions {
	s.ALL = field.NewAsterisk(table)
	s.FID = field.NewInt64(table, "f_id")
	s.FScacs = field.NewString(table, "f_scacs")
	s.FSearchDate = field.NewString(table, "f_search_date")
	s.FOriginPort = field.NewString(table, "f_origin_port")
	s.FDestinationPort = field.NewString(table, "f_destination_port")
	s.FWeeksOut = field.NewString(table, "f_weeks_out")
	s.FCreatedAt = field.NewTime(table, "f_created_at")
	s.FUpdatedAt = field.NewTime(table, "f_updated_at")
	s.FDeleted = field.NewInt16(table, "f_deleted")
	s.FIsTransit = field.NewInt16(table, "f_is_transit")
	s.FDeletedAt = field.NewTime(table, "f_deleted_at")

	s.fillFieldMap()

	return s
}

func (s *shippingConditions) WithContext(ctx context.Context) *shippingConditionsDo {
	return s.shippingConditionsDo.WithContext(ctx)
}

func (s shippingConditions) TableName() string { return s.shippingConditionsDo.TableName() }

func (s shippingConditions) Alias() string { return s.shippingConditionsDo.Alias() }

func (s *shippingConditions) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *shippingConditions) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["f_id"] = s.FID
	s.fieldMap["f_scacs"] = s.FScacs
	s.fieldMap["f_search_date"] = s.FSearchDate
	s.fieldMap["f_origin_port"] = s.FOriginPort
	s.fieldMap["f_destination_port"] = s.FDestinationPort
	s.fieldMap["f_weeks_out"] = s.FWeeksOut
	s.fieldMap["f_created_at"] = s.FCreatedAt
	s.fieldMap["f_updated_at"] = s.FUpdatedAt
	s.fieldMap["f_deleted"] = s.FDeleted
	s.fieldMap["f_is_transit"] = s.FIsTransit
	s.fieldMap["f_deleted_at"] = s.FDeletedAt
}

func (s shippingConditions) clone(db *gorm.DB) shippingConditions {
	s.shippingConditionsDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s shippingConditions) replaceDB(db *gorm.DB) shippingConditions {
	s.shippingConditionsDo.ReplaceDB(db)
	return s
}

type shippingConditionsDo struct{ gen.DO }

func (s shippingConditionsDo) Debug() *shippingConditionsDo {
	return s.withDO(s.DO.Debug())
}

func (s shippingConditionsDo) WithContext(ctx context.Context) *shippingConditionsDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s shippingConditionsDo) ReadDB() *shippingConditionsDo {
	return s.Clauses(dbresolver.Read)
}

func (s shippingConditionsDo) WriteDB() *shippingConditionsDo {
	return s.Clauses(dbresolver.Write)
}

func (s shippingConditionsDo) Session(config *gorm.Session) *shippingConditionsDo {
	return s.withDO(s.DO.Session(config))
}

func (s shippingConditionsDo) Clauses(conds ...clause.Expression) *shippingConditionsDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s shippingConditionsDo) Returning(value interface{}, columns ...string) *shippingConditionsDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s shippingConditionsDo) Not(conds ...gen.Condition) *shippingConditionsDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s shippingConditionsDo) Or(conds ...gen.Condition) *shippingConditionsDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s shippingConditionsDo) Select(conds ...field.Expr) *shippingConditionsDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s shippingConditionsDo) Where(conds ...gen.Condition) *shippingConditionsDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s shippingConditionsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *shippingConditionsDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s shippingConditionsDo) Order(conds ...field.Expr) *shippingConditionsDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s shippingConditionsDo) Distinct(cols ...field.Expr) *shippingConditionsDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s shippingConditionsDo) Omit(cols ...field.Expr) *shippingConditionsDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s shippingConditionsDo) Join(table schema.Tabler, on ...field.Expr) *shippingConditionsDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s shippingConditionsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *shippingConditionsDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s shippingConditionsDo) RightJoin(table schema.Tabler, on ...field.Expr) *shippingConditionsDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s shippingConditionsDo) Group(cols ...field.Expr) *shippingConditionsDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s shippingConditionsDo) Having(conds ...gen.Condition) *shippingConditionsDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s shippingConditionsDo) Limit(limit int) *shippingConditionsDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s shippingConditionsDo) Offset(offset int) *shippingConditionsDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s shippingConditionsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *shippingConditionsDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s shippingConditionsDo) Unscoped() *shippingConditionsDo {
	return s.withDO(s.DO.Unscoped())
}

func (s shippingConditionsDo) Create(values ...*model.ShippingConditions) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s shippingConditionsDo) CreateInBatches(values []*model.ShippingConditions, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s shippingConditionsDo) Save(values ...*model.ShippingConditions) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s shippingConditionsDo) First() (*model.ShippingConditions, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingConditions), nil
	}
}

func (s shippingConditionsDo) Take() (*model.ShippingConditions, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingConditions), nil
	}
}

func (s shippingConditionsDo) Last() (*model.ShippingConditions, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingConditions), nil
	}
}

func (s shippingConditionsDo) Find() ([]*model.ShippingConditions, error) {
	result, err := s.DO.Find()
	return result.([]*model.ShippingConditions), err
}

func (s shippingConditionsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ShippingConditions, err error) {
	buf := make([]*model.ShippingConditions, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s shippingConditionsDo) FindInBatches(result *[]*model.ShippingConditions, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s shippingConditionsDo) Attrs(attrs ...field.AssignExpr) *shippingConditionsDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s shippingConditionsDo) Assign(attrs ...field.AssignExpr) *shippingConditionsDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s shippingConditionsDo) Joins(fields ...field.RelationField) *shippingConditionsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s shippingConditionsDo) Preload(fields ...field.RelationField) *shippingConditionsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s shippingConditionsDo) FirstOrInit() (*model.ShippingConditions, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingConditions), nil
	}
}

func (s shippingConditionsDo) FirstOrCreate() (*model.ShippingConditions, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingConditions), nil
	}
}

func (s shippingConditionsDo) FindByPage(offset int, limit int) (result []*model.ShippingConditions, count int64, err error) {
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

func (s shippingConditionsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s shippingConditionsDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s shippingConditionsDo) Delete(models ...*model.ShippingConditions) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *shippingConditionsDo) withDO(do gen.Dao) *shippingConditionsDo {
	s.DO = *do.(*gen.DO)
	return s
}
