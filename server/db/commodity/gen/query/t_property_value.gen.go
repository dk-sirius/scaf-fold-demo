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

func newPropertyValue(db *gorm.DB, opts ...gen.DOOption) propertyValue {
	_propertyValue := propertyValue{}

	_propertyValue.propertyValueDo.UseDB(db, opts...)
	_propertyValue.propertyValueDo.UseModel(&model.PropertyValue{})

	tableName := _propertyValue.propertyValueDo.TableName()
	_propertyValue.ALL = field.NewAsterisk(tableName)
	_propertyValue.FID = field.NewInt64(tableName, "f_id")
	_propertyValue.FTenantID = field.NewString(tableName, "f_tenant_id")
	_propertyValue.FRelationID = field.NewInt64(tableName, "f_relation_id")
	_propertyValue.FPropertyID = field.NewInt64(tableName, "f_property_id")
	_propertyValue.FPropertyEnumID = field.NewInt64(tableName, "f_property_enum_id")
	_propertyValue.FPropertyValue = field.NewString(tableName, "f_property_value")
	_propertyValue.FSort = field.NewInt32(tableName, "f_sort")
	_propertyValue.FDeleted = field.NewInt16(tableName, "f_deleted")
	_propertyValue.FCreatedBy = field.NewString(tableName, "f_created_by")
	_propertyValue.FUpdatedBy = field.NewString(tableName, "f_updated_by")
	_propertyValue.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_propertyValue.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_propertyValue.FVersion = field.NewInt64(tableName, "f_version")
	_propertyValue.FPropertyCode = field.NewString(tableName, "f_property_code")
	_propertyValue.FStatus = field.NewInt16(tableName, "f_status")
	_propertyValue.FType = field.NewInt16(tableName, "f_type")

	_propertyValue.fillFieldMap()

	return _propertyValue
}

type propertyValue struct {
	propertyValueDo propertyValueDo

	ALL             field.Asterisk
	FID             field.Int64  // 主键
	FTenantID       field.String // 租户id
	FRelationID     field.Int64  // 关联id（sku、spu）
	FPropertyID     field.Int64  // 商品关联属性id
	FPropertyEnumID field.Int64  // 关联属性枚举ID
	FPropertyValue  field.String // 属性值
	FSort           field.Int32  // 属性排序
	FDeleted        field.Int16  // 删除标记
	FCreatedBy      field.String // 创建者
	FUpdatedBy      field.String // 修改者
	FCreatedAt      field.Time   // 创建时间
	FUpdatedAt      field.Time   // 修改时间
	FVersion        field.Int64  // 乐观锁版本号
	FPropertyCode   field.String // 商品code
	FStatus         field.Int16  // 是否启用 unknown：0，是：1 ,否：2
	FType           field.Int16  // 关联类型 unknown：0，sku：1 ,spu：2

	fieldMap map[string]field.Expr
}

func (p propertyValue) Table(newTableName string) *propertyValue {
	p.propertyValueDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p propertyValue) As(alias string) *propertyValue {
	p.propertyValueDo.DO = *(p.propertyValueDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *propertyValue) updateTableName(table string) *propertyValue {
	p.ALL = field.NewAsterisk(table)
	p.FID = field.NewInt64(table, "f_id")
	p.FTenantID = field.NewString(table, "f_tenant_id")
	p.FRelationID = field.NewInt64(table, "f_relation_id")
	p.FPropertyID = field.NewInt64(table, "f_property_id")
	p.FPropertyEnumID = field.NewInt64(table, "f_property_enum_id")
	p.FPropertyValue = field.NewString(table, "f_property_value")
	p.FSort = field.NewInt32(table, "f_sort")
	p.FDeleted = field.NewInt16(table, "f_deleted")
	p.FCreatedBy = field.NewString(table, "f_created_by")
	p.FUpdatedBy = field.NewString(table, "f_updated_by")
	p.FCreatedAt = field.NewTime(table, "f_created_at")
	p.FUpdatedAt = field.NewTime(table, "f_updated_at")
	p.FVersion = field.NewInt64(table, "f_version")
	p.FPropertyCode = field.NewString(table, "f_property_code")
	p.FStatus = field.NewInt16(table, "f_status")
	p.FType = field.NewInt16(table, "f_type")

	p.fillFieldMap()

	return p
}

func (p *propertyValue) WithContext(ctx context.Context) *propertyValueDo {
	return p.propertyValueDo.WithContext(ctx)
}

func (p propertyValue) TableName() string { return p.propertyValueDo.TableName() }

func (p propertyValue) Alias() string { return p.propertyValueDo.Alias() }

func (p *propertyValue) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *propertyValue) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 16)
	p.fieldMap["f_id"] = p.FID
	p.fieldMap["f_tenant_id"] = p.FTenantID
	p.fieldMap["f_relation_id"] = p.FRelationID
	p.fieldMap["f_property_id"] = p.FPropertyID
	p.fieldMap["f_property_enum_id"] = p.FPropertyEnumID
	p.fieldMap["f_property_value"] = p.FPropertyValue
	p.fieldMap["f_sort"] = p.FSort
	p.fieldMap["f_deleted"] = p.FDeleted
	p.fieldMap["f_created_by"] = p.FCreatedBy
	p.fieldMap["f_updated_by"] = p.FUpdatedBy
	p.fieldMap["f_created_at"] = p.FCreatedAt
	p.fieldMap["f_updated_at"] = p.FUpdatedAt
	p.fieldMap["f_version"] = p.FVersion
	p.fieldMap["f_property_code"] = p.FPropertyCode
	p.fieldMap["f_status"] = p.FStatus
	p.fieldMap["f_type"] = p.FType
}

func (p propertyValue) clone(db *gorm.DB) propertyValue {
	p.propertyValueDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p propertyValue) replaceDB(db *gorm.DB) propertyValue {
	p.propertyValueDo.ReplaceDB(db)
	return p
}

type propertyValueDo struct{ gen.DO }

func (p propertyValueDo) Debug() *propertyValueDo {
	return p.withDO(p.DO.Debug())
}

func (p propertyValueDo) WithContext(ctx context.Context) *propertyValueDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p propertyValueDo) ReadDB() *propertyValueDo {
	return p.Clauses(dbresolver.Read)
}

func (p propertyValueDo) WriteDB() *propertyValueDo {
	return p.Clauses(dbresolver.Write)
}

func (p propertyValueDo) Session(config *gorm.Session) *propertyValueDo {
	return p.withDO(p.DO.Session(config))
}

func (p propertyValueDo) Clauses(conds ...clause.Expression) *propertyValueDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p propertyValueDo) Returning(value interface{}, columns ...string) *propertyValueDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p propertyValueDo) Not(conds ...gen.Condition) *propertyValueDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p propertyValueDo) Or(conds ...gen.Condition) *propertyValueDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p propertyValueDo) Select(conds ...field.Expr) *propertyValueDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p propertyValueDo) Where(conds ...gen.Condition) *propertyValueDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p propertyValueDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *propertyValueDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p propertyValueDo) Order(conds ...field.Expr) *propertyValueDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p propertyValueDo) Distinct(cols ...field.Expr) *propertyValueDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p propertyValueDo) Omit(cols ...field.Expr) *propertyValueDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p propertyValueDo) Join(table schema.Tabler, on ...field.Expr) *propertyValueDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p propertyValueDo) LeftJoin(table schema.Tabler, on ...field.Expr) *propertyValueDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p propertyValueDo) RightJoin(table schema.Tabler, on ...field.Expr) *propertyValueDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p propertyValueDo) Group(cols ...field.Expr) *propertyValueDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p propertyValueDo) Having(conds ...gen.Condition) *propertyValueDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p propertyValueDo) Limit(limit int) *propertyValueDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p propertyValueDo) Offset(offset int) *propertyValueDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p propertyValueDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *propertyValueDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p propertyValueDo) Unscoped() *propertyValueDo {
	return p.withDO(p.DO.Unscoped())
}

func (p propertyValueDo) Create(values ...*model.PropertyValue) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p propertyValueDo) CreateInBatches(values []*model.PropertyValue, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p propertyValueDo) Save(values ...*model.PropertyValue) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p propertyValueDo) First() (*model.PropertyValue, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PropertyValue), nil
	}
}

func (p propertyValueDo) Take() (*model.PropertyValue, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PropertyValue), nil
	}
}

func (p propertyValueDo) Last() (*model.PropertyValue, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PropertyValue), nil
	}
}

func (p propertyValueDo) Find() ([]*model.PropertyValue, error) {
	result, err := p.DO.Find()
	return result.([]*model.PropertyValue), err
}

func (p propertyValueDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PropertyValue, err error) {
	buf := make([]*model.PropertyValue, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p propertyValueDo) FindInBatches(result *[]*model.PropertyValue, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p propertyValueDo) Attrs(attrs ...field.AssignExpr) *propertyValueDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p propertyValueDo) Assign(attrs ...field.AssignExpr) *propertyValueDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p propertyValueDo) Joins(fields ...field.RelationField) *propertyValueDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p propertyValueDo) Preload(fields ...field.RelationField) *propertyValueDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p propertyValueDo) FirstOrInit() (*model.PropertyValue, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PropertyValue), nil
	}
}

func (p propertyValueDo) FirstOrCreate() (*model.PropertyValue, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PropertyValue), nil
	}
}

func (p propertyValueDo) FindByPage(offset int, limit int) (result []*model.PropertyValue, count int64, err error) {
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

func (p propertyValueDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p propertyValueDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p propertyValueDo) Delete(models ...*model.PropertyValue) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *propertyValueDo) withDO(do gen.Dao) *propertyValueDo {
	p.DO = *do.(*gen.DO)
	return p
}
