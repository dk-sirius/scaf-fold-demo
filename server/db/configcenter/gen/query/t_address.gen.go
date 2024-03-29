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

func newAddress(db *gorm.DB, opts ...gen.DOOption) address {
	_address := address{}

	_address.addressDo.UseDB(db, opts...)
	_address.addressDo.UseModel(&model.Address{})

	tableName := _address.addressDo.TableName()
	_address.ALL = field.NewAsterisk(tableName)
	_address.FID = field.NewInt64(tableName, "f_id")
	_address.FStatus = field.NewInt16(tableName, "f_status")
	_address.FCreatedBy = field.NewString(tableName, "f_created_by")
	_address.FUpdatedBy = field.NewString(tableName, "f_updated_by")
	_address.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_address.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_address.FVersion = field.NewInt64(tableName, "f_version")
	_address.FRemark = field.NewString(tableName, "f_remark")
	_address.FDeleted = field.NewInt16(tableName, "f_deleted")
	_address.FNameZh = field.NewString(tableName, "f_name_zh")
	_address.FTenantID = field.NewInt64(tableName, "f_tenant_id")
	_address.FLevel = field.NewInt16(tableName, "f_level")
	_address.FNameEn = field.NewString(tableName, "f_name_en")
	_address.FParentID = field.NewString(tableName, "f_parent_id")
	_address.FLng = field.NewString(tableName, "f_lng")
	_address.FLat = field.NewString(tableName, "f_lat")
	_address.FGeom = field.NewString(tableName, "f_geom")
	_address.FOverseasFlag = field.NewInt16(tableName, "f_overseas_flag")
	_address.FCode = field.NewString(tableName, "f_code")

	_address.fillFieldMap()

	return _address
}

type address struct {
	addressDo addressDo

	ALL           field.Asterisk
	FID           field.Int64  // 主键
	FStatus       field.Int16  // 是否启用unknown：0，是：1 ,否：2
	FCreatedBy    field.String // 创建人
	FUpdatedBy    field.String // 修改人
	FCreatedAt    field.Time   // 创建时间
	FUpdatedAt    field.Time   // 最后更新时间
	FVersion      field.Int64  // 乐观锁
	FRemark       field.String // 备注
	FDeleted      field.Int16  // 删除标记 1：删除 0：未删除
	FNameZh       field.String // 中文名
	FTenantID     field.Int64
	FLevel        field.Int16  // 地址级别：unknown：0，洲：1，国家：2，省：3，城市/地区：4
	FNameEn       field.String // 英文名
	FParentID     field.String // 父级id以逗号隔开
	FLng          field.String // 经度
	FLat          field.String // 纬度
	FGeom         field.String // 坐标
	FOverseasFlag field.Int16  // 境外标识：unknown：0，境内：1，境外：2
	FCode         field.String // 地址代码

	fieldMap map[string]field.Expr
}

func (a address) Table(newTableName string) *address {
	a.addressDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a address) As(alias string) *address {
	a.addressDo.DO = *(a.addressDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *address) updateTableName(table string) *address {
	a.ALL = field.NewAsterisk(table)
	a.FID = field.NewInt64(table, "f_id")
	a.FStatus = field.NewInt16(table, "f_status")
	a.FCreatedBy = field.NewString(table, "f_created_by")
	a.FUpdatedBy = field.NewString(table, "f_updated_by")
	a.FCreatedAt = field.NewTime(table, "f_created_at")
	a.FUpdatedAt = field.NewTime(table, "f_updated_at")
	a.FVersion = field.NewInt64(table, "f_version")
	a.FRemark = field.NewString(table, "f_remark")
	a.FDeleted = field.NewInt16(table, "f_deleted")
	a.FNameZh = field.NewString(table, "f_name_zh")
	a.FTenantID = field.NewInt64(table, "f_tenant_id")
	a.FLevel = field.NewInt16(table, "f_level")
	a.FNameEn = field.NewString(table, "f_name_en")
	a.FParentID = field.NewString(table, "f_parent_id")
	a.FLng = field.NewString(table, "f_lng")
	a.FLat = field.NewString(table, "f_lat")
	a.FGeom = field.NewString(table, "f_geom")
	a.FOverseasFlag = field.NewInt16(table, "f_overseas_flag")
	a.FCode = field.NewString(table, "f_code")

	a.fillFieldMap()

	return a
}

func (a *address) WithContext(ctx context.Context) *addressDo { return a.addressDo.WithContext(ctx) }

func (a address) TableName() string { return a.addressDo.TableName() }

func (a address) Alias() string { return a.addressDo.Alias() }

func (a *address) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *address) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 19)
	a.fieldMap["f_id"] = a.FID
	a.fieldMap["f_status"] = a.FStatus
	a.fieldMap["f_created_by"] = a.FCreatedBy
	a.fieldMap["f_updated_by"] = a.FUpdatedBy
	a.fieldMap["f_created_at"] = a.FCreatedAt
	a.fieldMap["f_updated_at"] = a.FUpdatedAt
	a.fieldMap["f_version"] = a.FVersion
	a.fieldMap["f_remark"] = a.FRemark
	a.fieldMap["f_deleted"] = a.FDeleted
	a.fieldMap["f_name_zh"] = a.FNameZh
	a.fieldMap["f_tenant_id"] = a.FTenantID
	a.fieldMap["f_level"] = a.FLevel
	a.fieldMap["f_name_en"] = a.FNameEn
	a.fieldMap["f_parent_id"] = a.FParentID
	a.fieldMap["f_lng"] = a.FLng
	a.fieldMap["f_lat"] = a.FLat
	a.fieldMap["f_geom"] = a.FGeom
	a.fieldMap["f_overseas_flag"] = a.FOverseasFlag
	a.fieldMap["f_code"] = a.FCode
}

func (a address) clone(db *gorm.DB) address {
	a.addressDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a address) replaceDB(db *gorm.DB) address {
	a.addressDo.ReplaceDB(db)
	return a
}

type addressDo struct{ gen.DO }

func (a addressDo) Debug() *addressDo {
	return a.withDO(a.DO.Debug())
}

func (a addressDo) WithContext(ctx context.Context) *addressDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a addressDo) ReadDB() *addressDo {
	return a.Clauses(dbresolver.Read)
}

func (a addressDo) WriteDB() *addressDo {
	return a.Clauses(dbresolver.Write)
}

func (a addressDo) Session(config *gorm.Session) *addressDo {
	return a.withDO(a.DO.Session(config))
}

func (a addressDo) Clauses(conds ...clause.Expression) *addressDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a addressDo) Returning(value interface{}, columns ...string) *addressDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a addressDo) Not(conds ...gen.Condition) *addressDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a addressDo) Or(conds ...gen.Condition) *addressDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a addressDo) Select(conds ...field.Expr) *addressDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a addressDo) Where(conds ...gen.Condition) *addressDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a addressDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *addressDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a addressDo) Order(conds ...field.Expr) *addressDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a addressDo) Distinct(cols ...field.Expr) *addressDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a addressDo) Omit(cols ...field.Expr) *addressDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a addressDo) Join(table schema.Tabler, on ...field.Expr) *addressDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a addressDo) LeftJoin(table schema.Tabler, on ...field.Expr) *addressDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a addressDo) RightJoin(table schema.Tabler, on ...field.Expr) *addressDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a addressDo) Group(cols ...field.Expr) *addressDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a addressDo) Having(conds ...gen.Condition) *addressDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a addressDo) Limit(limit int) *addressDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a addressDo) Offset(offset int) *addressDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a addressDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *addressDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a addressDo) Unscoped() *addressDo {
	return a.withDO(a.DO.Unscoped())
}

func (a addressDo) Create(values ...*model.Address) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a addressDo) CreateInBatches(values []*model.Address, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a addressDo) Save(values ...*model.Address) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a addressDo) First() (*model.Address, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) Take() (*model.Address, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) Last() (*model.Address, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) Find() ([]*model.Address, error) {
	result, err := a.DO.Find()
	return result.([]*model.Address), err
}

func (a addressDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Address, err error) {
	buf := make([]*model.Address, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a addressDo) FindInBatches(result *[]*model.Address, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a addressDo) Attrs(attrs ...field.AssignExpr) *addressDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a addressDo) Assign(attrs ...field.AssignExpr) *addressDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a addressDo) Joins(fields ...field.RelationField) *addressDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a addressDo) Preload(fields ...field.RelationField) *addressDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a addressDo) FirstOrInit() (*model.Address, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) FirstOrCreate() (*model.Address, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) FindByPage(offset int, limit int) (result []*model.Address, count int64, err error) {
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

func (a addressDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a addressDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a addressDo) Delete(models ...*model.Address) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *addressDo) withDO(do gen.Dao) *addressDo {
	a.DO = *do.(*gen.DO)
	return a
}
