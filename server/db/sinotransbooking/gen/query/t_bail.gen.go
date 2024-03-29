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

func newBail(db *gorm.DB, opts ...gen.DOOption) bail {
	_bail := bail{}

	_bail.bailDo.UseDB(db, opts...)
	_bail.bailDo.UseModel(&model.Bail{})

	tableName := _bail.bailDo.TableName()
	_bail.ALL = field.NewAsterisk(tableName)
	_bail.FID = field.NewInt64(tableName, "f_id")
	_bail.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_bail.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_bail.FDeleted = field.NewInt16(tableName, "f_deleted")
	_bail.FCurrency = field.NewFloat64(tableName, "f_currency")
	_bail.FCurrencyType = field.NewInt32(tableName, "f_currency_type")
	_bail.FEnable = field.NewInt32(tableName, "f_enable")
	_bail.FBailID = field.NewInt64(tableName, "f_bail_id")

	_bail.fillFieldMap()

	return _bail
}

type bail struct {
	bailDo bailDo

	ALL           field.Asterisk
	FID           field.Int64   // 自增主键
	FCreatedAt    field.Time    // 创建时间
	FUpdatedAt    field.Time    // 更新时间
	FDeleted      field.Int16   // 删除：0-未删除，1-删除
	FCurrency     field.Float64 // 保证金
	FCurrencyType field.Int32   // 货币类型
	FEnable       field.Int32   // 保证金是否生效，0-生效，1-不生效
	FBailID       field.Int64   // 保证金id

	fieldMap map[string]field.Expr
}

func (b bail) Table(newTableName string) *bail {
	b.bailDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bail) As(alias string) *bail {
	b.bailDo.DO = *(b.bailDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bail) updateTableName(table string) *bail {
	b.ALL = field.NewAsterisk(table)
	b.FID = field.NewInt64(table, "f_id")
	b.FCreatedAt = field.NewTime(table, "f_created_at")
	b.FUpdatedAt = field.NewTime(table, "f_updated_at")
	b.FDeleted = field.NewInt16(table, "f_deleted")
	b.FCurrency = field.NewFloat64(table, "f_currency")
	b.FCurrencyType = field.NewInt32(table, "f_currency_type")
	b.FEnable = field.NewInt32(table, "f_enable")
	b.FBailID = field.NewInt64(table, "f_bail_id")

	b.fillFieldMap()

	return b
}

func (b *bail) WithContext(ctx context.Context) *bailDo { return b.bailDo.WithContext(ctx) }

func (b bail) TableName() string { return b.bailDo.TableName() }

func (b bail) Alias() string { return b.bailDo.Alias() }

func (b *bail) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bail) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 8)
	b.fieldMap["f_id"] = b.FID
	b.fieldMap["f_created_at"] = b.FCreatedAt
	b.fieldMap["f_updated_at"] = b.FUpdatedAt
	b.fieldMap["f_deleted"] = b.FDeleted
	b.fieldMap["f_currency"] = b.FCurrency
	b.fieldMap["f_currency_type"] = b.FCurrencyType
	b.fieldMap["f_enable"] = b.FEnable
	b.fieldMap["f_bail_id"] = b.FBailID
}

func (b bail) clone(db *gorm.DB) bail {
	b.bailDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bail) replaceDB(db *gorm.DB) bail {
	b.bailDo.ReplaceDB(db)
	return b
}

type bailDo struct{ gen.DO }

func (b bailDo) Debug() *bailDo {
	return b.withDO(b.DO.Debug())
}

func (b bailDo) WithContext(ctx context.Context) *bailDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bailDo) ReadDB() *bailDo {
	return b.Clauses(dbresolver.Read)
}

func (b bailDo) WriteDB() *bailDo {
	return b.Clauses(dbresolver.Write)
}

func (b bailDo) Session(config *gorm.Session) *bailDo {
	return b.withDO(b.DO.Session(config))
}

func (b bailDo) Clauses(conds ...clause.Expression) *bailDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bailDo) Returning(value interface{}, columns ...string) *bailDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bailDo) Not(conds ...gen.Condition) *bailDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bailDo) Or(conds ...gen.Condition) *bailDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bailDo) Select(conds ...field.Expr) *bailDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bailDo) Where(conds ...gen.Condition) *bailDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bailDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *bailDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b bailDo) Order(conds ...field.Expr) *bailDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bailDo) Distinct(cols ...field.Expr) *bailDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bailDo) Omit(cols ...field.Expr) *bailDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bailDo) Join(table schema.Tabler, on ...field.Expr) *bailDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bailDo) LeftJoin(table schema.Tabler, on ...field.Expr) *bailDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bailDo) RightJoin(table schema.Tabler, on ...field.Expr) *bailDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bailDo) Group(cols ...field.Expr) *bailDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bailDo) Having(conds ...gen.Condition) *bailDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bailDo) Limit(limit int) *bailDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bailDo) Offset(offset int) *bailDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bailDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *bailDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bailDo) Unscoped() *bailDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bailDo) Create(values ...*model.Bail) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bailDo) CreateInBatches(values []*model.Bail, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bailDo) Save(values ...*model.Bail) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bailDo) First() (*model.Bail, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bail), nil
	}
}

func (b bailDo) Take() (*model.Bail, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bail), nil
	}
}

func (b bailDo) Last() (*model.Bail, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bail), nil
	}
}

func (b bailDo) Find() ([]*model.Bail, error) {
	result, err := b.DO.Find()
	return result.([]*model.Bail), err
}

func (b bailDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Bail, err error) {
	buf := make([]*model.Bail, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bailDo) FindInBatches(result *[]*model.Bail, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bailDo) Attrs(attrs ...field.AssignExpr) *bailDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bailDo) Assign(attrs ...field.AssignExpr) *bailDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bailDo) Joins(fields ...field.RelationField) *bailDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bailDo) Preload(fields ...field.RelationField) *bailDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bailDo) FirstOrInit() (*model.Bail, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bail), nil
	}
}

func (b bailDo) FirstOrCreate() (*model.Bail, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bail), nil
	}
}

func (b bailDo) FindByPage(offset int, limit int) (result []*model.Bail, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bailDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bailDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bailDo) Delete(models ...*model.Bail) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bailDo) withDO(do gen.Dao) *bailDo {
	b.DO = *do.(*gen.DO)
	return b
}
