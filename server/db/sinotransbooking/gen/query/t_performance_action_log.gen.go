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

func newActionLog(db *gorm.DB, opts ...gen.DOOption) actionLog {
	_actionLog := actionLog{}

	_actionLog.actionLogDo.UseDB(db, opts...)
	_actionLog.actionLogDo.UseModel(&model.ActionLog{})

	tableName := _actionLog.actionLogDo.TableName()
	_actionLog.ALL = field.NewAsterisk(tableName)
	_actionLog.FID = field.NewInt64(tableName, "f_id")
	_actionLog.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_actionLog.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_actionLog.FDeleted = field.NewInt16(tableName, "f_deleted")
	_actionLog.FAction = field.NewInt32(tableName, "f_action")
	_actionLog.FActionContent = field.NewString(tableName, "f_action_content")
	_actionLog.FOperatorID = field.NewInt64(tableName, "f_operator_id")
	_actionLog.FPerformanceID = field.NewInt64(tableName, "f_performance_id")
	_actionLog.FActionID = field.NewInt64(tableName, "f_action_id")

	_actionLog.fillFieldMap()

	return _actionLog
}

type actionLog struct {
	actionLogDo actionLogDo

	ALL            field.Asterisk
	FID            field.Int64  // 自增主键
	FCreatedAt     field.Time   // 创建时间
	FUpdatedAt     field.Time   // 更新时间
	FDeleted       field.Int16  // 逻辑删除标识0-未删除，1-删除
	FAction        field.Int32  // 行为枚举值
	FActionContent field.String // 行为操作内容
	FOperatorID    field.Int64  // 操作人id
	FPerformanceID field.Int64  // 履约id
	FActionID      field.Int64  // 行为id

	fieldMap map[string]field.Expr
}

func (a actionLog) Table(newTableName string) *actionLog {
	a.actionLogDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a actionLog) As(alias string) *actionLog {
	a.actionLogDo.DO = *(a.actionLogDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *actionLog) updateTableName(table string) *actionLog {
	a.ALL = field.NewAsterisk(table)
	a.FID = field.NewInt64(table, "f_id")
	a.FCreatedAt = field.NewTime(table, "f_created_at")
	a.FUpdatedAt = field.NewTime(table, "f_updated_at")
	a.FDeleted = field.NewInt16(table, "f_deleted")
	a.FAction = field.NewInt32(table, "f_action")
	a.FActionContent = field.NewString(table, "f_action_content")
	a.FOperatorID = field.NewInt64(table, "f_operator_id")
	a.FPerformanceID = field.NewInt64(table, "f_performance_id")
	a.FActionID = field.NewInt64(table, "f_action_id")

	a.fillFieldMap()

	return a
}

func (a *actionLog) WithContext(ctx context.Context) *actionLogDo {
	return a.actionLogDo.WithContext(ctx)
}

func (a actionLog) TableName() string { return a.actionLogDo.TableName() }

func (a actionLog) Alias() string { return a.actionLogDo.Alias() }

func (a *actionLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *actionLog) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["f_id"] = a.FID
	a.fieldMap["f_created_at"] = a.FCreatedAt
	a.fieldMap["f_updated_at"] = a.FUpdatedAt
	a.fieldMap["f_deleted"] = a.FDeleted
	a.fieldMap["f_action"] = a.FAction
	a.fieldMap["f_action_content"] = a.FActionContent
	a.fieldMap["f_operator_id"] = a.FOperatorID
	a.fieldMap["f_performance_id"] = a.FPerformanceID
	a.fieldMap["f_action_id"] = a.FActionID
}

func (a actionLog) clone(db *gorm.DB) actionLog {
	a.actionLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a actionLog) replaceDB(db *gorm.DB) actionLog {
	a.actionLogDo.ReplaceDB(db)
	return a
}

type actionLogDo struct{ gen.DO }

func (a actionLogDo) Debug() *actionLogDo {
	return a.withDO(a.DO.Debug())
}

func (a actionLogDo) WithContext(ctx context.Context) *actionLogDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a actionLogDo) ReadDB() *actionLogDo {
	return a.Clauses(dbresolver.Read)
}

func (a actionLogDo) WriteDB() *actionLogDo {
	return a.Clauses(dbresolver.Write)
}

func (a actionLogDo) Session(config *gorm.Session) *actionLogDo {
	return a.withDO(a.DO.Session(config))
}

func (a actionLogDo) Clauses(conds ...clause.Expression) *actionLogDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a actionLogDo) Returning(value interface{}, columns ...string) *actionLogDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a actionLogDo) Not(conds ...gen.Condition) *actionLogDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a actionLogDo) Or(conds ...gen.Condition) *actionLogDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a actionLogDo) Select(conds ...field.Expr) *actionLogDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a actionLogDo) Where(conds ...gen.Condition) *actionLogDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a actionLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *actionLogDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a actionLogDo) Order(conds ...field.Expr) *actionLogDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a actionLogDo) Distinct(cols ...field.Expr) *actionLogDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a actionLogDo) Omit(cols ...field.Expr) *actionLogDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a actionLogDo) Join(table schema.Tabler, on ...field.Expr) *actionLogDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a actionLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *actionLogDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a actionLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *actionLogDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a actionLogDo) Group(cols ...field.Expr) *actionLogDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a actionLogDo) Having(conds ...gen.Condition) *actionLogDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a actionLogDo) Limit(limit int) *actionLogDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a actionLogDo) Offset(offset int) *actionLogDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a actionLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *actionLogDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a actionLogDo) Unscoped() *actionLogDo {
	return a.withDO(a.DO.Unscoped())
}

func (a actionLogDo) Create(values ...*model.ActionLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a actionLogDo) CreateInBatches(values []*model.ActionLog, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a actionLogDo) Save(values ...*model.ActionLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a actionLogDo) First() (*model.ActionLog, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActionLog), nil
	}
}

func (a actionLogDo) Take() (*model.ActionLog, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActionLog), nil
	}
}

func (a actionLogDo) Last() (*model.ActionLog, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActionLog), nil
	}
}

func (a actionLogDo) Find() ([]*model.ActionLog, error) {
	result, err := a.DO.Find()
	return result.([]*model.ActionLog), err
}

func (a actionLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActionLog, err error) {
	buf := make([]*model.ActionLog, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a actionLogDo) FindInBatches(result *[]*model.ActionLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a actionLogDo) Attrs(attrs ...field.AssignExpr) *actionLogDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a actionLogDo) Assign(attrs ...field.AssignExpr) *actionLogDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a actionLogDo) Joins(fields ...field.RelationField) *actionLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a actionLogDo) Preload(fields ...field.RelationField) *actionLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a actionLogDo) FirstOrInit() (*model.ActionLog, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActionLog), nil
	}
}

func (a actionLogDo) FirstOrCreate() (*model.ActionLog, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActionLog), nil
	}
}

func (a actionLogDo) FindByPage(offset int, limit int) (result []*model.ActionLog, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a actionLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a actionLogDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a actionLogDo) Delete(models ...*model.ActionLog) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *actionLogDo) withDO(do gen.Dao) *actionLogDo {
	a.DO = *do.(*gen.DO)
	return a
}
