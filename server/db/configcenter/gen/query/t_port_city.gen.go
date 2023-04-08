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

func newPortCity(db *gorm.DB, opts ...gen.DOOption) portCity {
	_portCity := portCity{}

	_portCity.portCityDo.UseDB(db, opts...)
	_portCity.portCityDo.UseModel(&model.PortCity{})

	tableName := _portCity.portCityDo.TableName()
	_portCity.ALL = field.NewAsterisk(tableName)
	_portCity.FID = field.NewInt64(tableName, "f_id")
	_portCity.FStatus = field.NewInt16(tableName, "f_status")
	_portCity.FCreatedBy = field.NewString(tableName, "f_created_by")
	_portCity.FUpdatedBy = field.NewString(tableName, "f_updated_by")
	_portCity.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_portCity.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_portCity.FVersion = field.NewInt64(tableName, "f_version")
	_portCity.FRemark = field.NewString(tableName, "f_remark")
	_portCity.FDeleted = field.NewInt16(tableName, "f_deleted")
	_portCity.FNameZh = field.NewString(tableName, "f_name_zh")
	_portCity.FTenantID = field.NewInt64(tableName, "f_tenant_id")
	_portCity.FNameEn = field.NewString(tableName, "f_name_en")
	_portCity.FAddressID = field.NewInt64(tableName, "f_address_id")

	_portCity.fillFieldMap()

	return _portCity
}

type portCity struct {
	portCityDo portCityDo

	ALL        field.Asterisk
	FID        field.Int64  // 主键
	FStatus    field.Int16  // 是否启用unknown：0，是：1 ,否：2
	FCreatedBy field.String // 创建人
	FUpdatedBy field.String // 修改人
	FCreatedAt field.Time   // 创建时间
	FUpdatedAt field.Time   // 最后更新时间
	FVersion   field.Int64  // 乐观锁
	FRemark    field.String // 备注
	FDeleted   field.Int16  // 删除标记 1：删除 0：未删除
	FNameZh    field.String // 中文名
	FTenantID  field.Int64
	FNameEn    field.String // 英文名
	FAddressID field.Int64  // 地址id

	fieldMap map[string]field.Expr
}

func (p portCity) Table(newTableName string) *portCity {
	p.portCityDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p portCity) As(alias string) *portCity {
	p.portCityDo.DO = *(p.portCityDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *portCity) updateTableName(table string) *portCity {
	p.ALL = field.NewAsterisk(table)
	p.FID = field.NewInt64(table, "f_id")
	p.FStatus = field.NewInt16(table, "f_status")
	p.FCreatedBy = field.NewString(table, "f_created_by")
	p.FUpdatedBy = field.NewString(table, "f_updated_by")
	p.FCreatedAt = field.NewTime(table, "f_created_at")
	p.FUpdatedAt = field.NewTime(table, "f_updated_at")
	p.FVersion = field.NewInt64(table, "f_version")
	p.FRemark = field.NewString(table, "f_remark")
	p.FDeleted = field.NewInt16(table, "f_deleted")
	p.FNameZh = field.NewString(table, "f_name_zh")
	p.FTenantID = field.NewInt64(table, "f_tenant_id")
	p.FNameEn = field.NewString(table, "f_name_en")
	p.FAddressID = field.NewInt64(table, "f_address_id")

	p.fillFieldMap()

	return p
}

func (p *portCity) WithContext(ctx context.Context) *portCityDo { return p.portCityDo.WithContext(ctx) }

func (p portCity) TableName() string { return p.portCityDo.TableName() }

func (p portCity) Alias() string { return p.portCityDo.Alias() }

func (p *portCity) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *portCity) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 13)
	p.fieldMap["f_id"] = p.FID
	p.fieldMap["f_status"] = p.FStatus
	p.fieldMap["f_created_by"] = p.FCreatedBy
	p.fieldMap["f_updated_by"] = p.FUpdatedBy
	p.fieldMap["f_created_at"] = p.FCreatedAt
	p.fieldMap["f_updated_at"] = p.FUpdatedAt
	p.fieldMap["f_version"] = p.FVersion
	p.fieldMap["f_remark"] = p.FRemark
	p.fieldMap["f_deleted"] = p.FDeleted
	p.fieldMap["f_name_zh"] = p.FNameZh
	p.fieldMap["f_tenant_id"] = p.FTenantID
	p.fieldMap["f_name_en"] = p.FNameEn
	p.fieldMap["f_address_id"] = p.FAddressID
}

func (p portCity) clone(db *gorm.DB) portCity {
	p.portCityDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p portCity) replaceDB(db *gorm.DB) portCity {
	p.portCityDo.ReplaceDB(db)
	return p
}

type portCityDo struct{ gen.DO }

func (p portCityDo) Debug() *portCityDo {
	return p.withDO(p.DO.Debug())
}

func (p portCityDo) WithContext(ctx context.Context) *portCityDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p portCityDo) ReadDB() *portCityDo {
	return p.Clauses(dbresolver.Read)
}

func (p portCityDo) WriteDB() *portCityDo {
	return p.Clauses(dbresolver.Write)
}

func (p portCityDo) Session(config *gorm.Session) *portCityDo {
	return p.withDO(p.DO.Session(config))
}

func (p portCityDo) Clauses(conds ...clause.Expression) *portCityDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p portCityDo) Returning(value interface{}, columns ...string) *portCityDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p portCityDo) Not(conds ...gen.Condition) *portCityDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p portCityDo) Or(conds ...gen.Condition) *portCityDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p portCityDo) Select(conds ...field.Expr) *portCityDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p portCityDo) Where(conds ...gen.Condition) *portCityDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p portCityDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *portCityDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p portCityDo) Order(conds ...field.Expr) *portCityDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p portCityDo) Distinct(cols ...field.Expr) *portCityDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p portCityDo) Omit(cols ...field.Expr) *portCityDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p portCityDo) Join(table schema.Tabler, on ...field.Expr) *portCityDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p portCityDo) LeftJoin(table schema.Tabler, on ...field.Expr) *portCityDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p portCityDo) RightJoin(table schema.Tabler, on ...field.Expr) *portCityDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p portCityDo) Group(cols ...field.Expr) *portCityDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p portCityDo) Having(conds ...gen.Condition) *portCityDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p portCityDo) Limit(limit int) *portCityDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p portCityDo) Offset(offset int) *portCityDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p portCityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *portCityDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p portCityDo) Unscoped() *portCityDo {
	return p.withDO(p.DO.Unscoped())
}

func (p portCityDo) Create(values ...*model.PortCity) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p portCityDo) CreateInBatches(values []*model.PortCity, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p portCityDo) Save(values ...*model.PortCity) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p portCityDo) First() (*model.PortCity, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PortCity), nil
	}
}

func (p portCityDo) Take() (*model.PortCity, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PortCity), nil
	}
}

func (p portCityDo) Last() (*model.PortCity, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PortCity), nil
	}
}

func (p portCityDo) Find() ([]*model.PortCity, error) {
	result, err := p.DO.Find()
	return result.([]*model.PortCity), err
}

func (p portCityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PortCity, err error) {
	buf := make([]*model.PortCity, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p portCityDo) FindInBatches(result *[]*model.PortCity, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p portCityDo) Attrs(attrs ...field.AssignExpr) *portCityDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p portCityDo) Assign(attrs ...field.AssignExpr) *portCityDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p portCityDo) Joins(fields ...field.RelationField) *portCityDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p portCityDo) Preload(fields ...field.RelationField) *portCityDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p portCityDo) FirstOrInit() (*model.PortCity, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PortCity), nil
	}
}

func (p portCityDo) FirstOrCreate() (*model.PortCity, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PortCity), nil
	}
}

func (p portCityDo) FindByPage(offset int, limit int) (result []*model.PortCity, count int64, err error) {
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

func (p portCityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p portCityDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p portCityDo) Delete(models ...*model.PortCity) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *portCityDo) withDO(do gen.Dao) *portCityDo {
	p.DO = *do.(*gen.DO)
	return p
}
