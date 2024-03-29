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

	"github.com/dk-sirius/scaf-fold-demo/server/db/configcenter/gen/model"
)

func newRouteGroup(db *gorm.DB, opts ...gen.DOOption) routeGroup {
	_routeGroup := routeGroup{}

	_routeGroup.routeGroupDo.UseDB(db, opts...)
	_routeGroup.routeGroupDo.UseModel(&model.RouteGroup{})

	tableName := _routeGroup.routeGroupDo.TableName()
	_routeGroup.ALL = field.NewAsterisk(tableName)
	_routeGroup.FID = field.NewInt64(tableName, "f_id")
	_routeGroup.FName = field.NewString(tableName, "f_name")
	_routeGroup.FRemark = field.NewString(tableName, "f_remark")
	_routeGroup.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_routeGroup.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_routeGroup.FDeleted = field.NewInt16(tableName, "f_deleted")
	_routeGroup.FCreatedBy = field.NewString(tableName, "f_created_by")
	_routeGroup.FUpdatedBy = field.NewString(tableName, "f_updated_by")
	_routeGroup.FVersion = field.NewInt64(tableName, "f_version")
	_routeGroup.FStatus = field.NewInt16(tableName, "f_status")
	_routeGroup.FTenantID = field.NewInt64(tableName, "f_tenant_id")

	_routeGroup.fillFieldMap()

	return _routeGroup
}

type routeGroup struct {
	routeGroupDo routeGroupDo

	ALL        field.Asterisk
	FID        field.Int64  // 自增主键
	FName      field.String // 航线组名称
	FRemark    field.String // 备注，默认为空
	FCreatedAt field.Time   // 创建时间
	FUpdatedAt field.Time   // 更新时间
	FDeleted   field.Int16
	FCreatedBy field.String // 创建人
	FUpdatedBy field.String // 修改人
	FVersion   field.Int64  // 乐观锁
	FStatus    field.Int16  // 是否启用unknown：0，是：1 ,否：2
	FTenantID  field.Int64

	fieldMap map[string]field.Expr
}

func (r routeGroup) Table(newTableName string) *routeGroup {
	r.routeGroupDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r routeGroup) As(alias string) *routeGroup {
	r.routeGroupDo.DO = *(r.routeGroupDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *routeGroup) updateTableName(table string) *routeGroup {
	r.ALL = field.NewAsterisk(table)
	r.FID = field.NewInt64(table, "f_id")
	r.FName = field.NewString(table, "f_name")
	r.FRemark = field.NewString(table, "f_remark")
	r.FCreatedAt = field.NewTime(table, "f_created_at")
	r.FUpdatedAt = field.NewTime(table, "f_updated_at")
	r.FDeleted = field.NewInt16(table, "f_deleted")
	r.FCreatedBy = field.NewString(table, "f_created_by")
	r.FUpdatedBy = field.NewString(table, "f_updated_by")
	r.FVersion = field.NewInt64(table, "f_version")
	r.FStatus = field.NewInt16(table, "f_status")
	r.FTenantID = field.NewInt64(table, "f_tenant_id")

	r.fillFieldMap()

	return r
}

func (r *routeGroup) WithContext(ctx context.Context) *routeGroupDo {
	return r.routeGroupDo.WithContext(ctx)
}

func (r routeGroup) TableName() string { return r.routeGroupDo.TableName() }

func (r routeGroup) Alias() string { return r.routeGroupDo.Alias() }

func (r *routeGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *routeGroup) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 11)
	r.fieldMap["f_id"] = r.FID
	r.fieldMap["f_name"] = r.FName
	r.fieldMap["f_remark"] = r.FRemark
	r.fieldMap["f_created_at"] = r.FCreatedAt
	r.fieldMap["f_updated_at"] = r.FUpdatedAt
	r.fieldMap["f_deleted"] = r.FDeleted
	r.fieldMap["f_created_by"] = r.FCreatedBy
	r.fieldMap["f_updated_by"] = r.FUpdatedBy
	r.fieldMap["f_version"] = r.FVersion
	r.fieldMap["f_status"] = r.FStatus
	r.fieldMap["f_tenant_id"] = r.FTenantID
}

func (r routeGroup) clone(db *gorm.DB) routeGroup {
	r.routeGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r routeGroup) replaceDB(db *gorm.DB) routeGroup {
	r.routeGroupDo.ReplaceDB(db)
	return r
}

type routeGroupDo struct{ gen.DO }

func (r routeGroupDo) Debug() *routeGroupDo {
	return r.withDO(r.DO.Debug())
}

func (r routeGroupDo) WithContext(ctx context.Context) *routeGroupDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r routeGroupDo) ReadDB() *routeGroupDo {
	return r.Clauses(dbresolver.Read)
}

func (r routeGroupDo) WriteDB() *routeGroupDo {
	return r.Clauses(dbresolver.Write)
}

func (r routeGroupDo) Session(config *gorm.Session) *routeGroupDo {
	return r.withDO(r.DO.Session(config))
}

func (r routeGroupDo) Clauses(conds ...clause.Expression) *routeGroupDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r routeGroupDo) Returning(value interface{}, columns ...string) *routeGroupDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r routeGroupDo) Not(conds ...gen.Condition) *routeGroupDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r routeGroupDo) Or(conds ...gen.Condition) *routeGroupDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r routeGroupDo) Select(conds ...field.Expr) *routeGroupDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r routeGroupDo) Where(conds ...gen.Condition) *routeGroupDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r routeGroupDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *routeGroupDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r routeGroupDo) Order(conds ...field.Expr) *routeGroupDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r routeGroupDo) Distinct(cols ...field.Expr) *routeGroupDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r routeGroupDo) Omit(cols ...field.Expr) *routeGroupDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r routeGroupDo) Join(table schema.Tabler, on ...field.Expr) *routeGroupDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r routeGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) *routeGroupDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r routeGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) *routeGroupDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r routeGroupDo) Group(cols ...field.Expr) *routeGroupDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r routeGroupDo) Having(conds ...gen.Condition) *routeGroupDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r routeGroupDo) Limit(limit int) *routeGroupDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r routeGroupDo) Offset(offset int) *routeGroupDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r routeGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *routeGroupDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r routeGroupDo) Unscoped() *routeGroupDo {
	return r.withDO(r.DO.Unscoped())
}

func (r routeGroupDo) Create(values ...*model.RouteGroup) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r routeGroupDo) CreateInBatches(values []*model.RouteGroup, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r routeGroupDo) Save(values ...*model.RouteGroup) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r routeGroupDo) First() (*model.RouteGroup, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteGroup), nil
	}
}

func (r routeGroupDo) Take() (*model.RouteGroup, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteGroup), nil
	}
}

func (r routeGroupDo) Last() (*model.RouteGroup, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteGroup), nil
	}
}

func (r routeGroupDo) Find() ([]*model.RouteGroup, error) {
	result, err := r.DO.Find()
	return result.([]*model.RouteGroup), err
}

func (r routeGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RouteGroup, err error) {
	buf := make([]*model.RouteGroup, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r routeGroupDo) FindInBatches(result *[]*model.RouteGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r routeGroupDo) Attrs(attrs ...field.AssignExpr) *routeGroupDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r routeGroupDo) Assign(attrs ...field.AssignExpr) *routeGroupDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r routeGroupDo) Joins(fields ...field.RelationField) *routeGroupDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r routeGroupDo) Preload(fields ...field.RelationField) *routeGroupDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r routeGroupDo) FirstOrInit() (*model.RouteGroup, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteGroup), nil
	}
}

func (r routeGroupDo) FirstOrCreate() (*model.RouteGroup, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteGroup), nil
	}
}

func (r routeGroupDo) FindByPage(offset int, limit int) (result []*model.RouteGroup, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r routeGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r routeGroupDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r routeGroupDo) Delete(models ...*model.RouteGroup) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *routeGroupDo) withDO(do gen.Dao) *routeGroupDo {
	r.DO = *do.(*gen.DO)
	return r
}
