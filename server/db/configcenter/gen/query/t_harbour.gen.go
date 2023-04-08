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

func newHarbour(db *gorm.DB, opts ...gen.DOOption) harbour {
	_harbour := harbour{}

	_harbour.harbourDo.UseDB(db, opts...)
	_harbour.harbourDo.UseModel(&model.Harbour{})

	tableName := _harbour.harbourDo.TableName()
	_harbour.ALL = field.NewAsterisk(tableName)
	_harbour.FID = field.NewInt64(tableName, "f_id")
	_harbour.FStatus = field.NewInt16(tableName, "f_status")
	_harbour.FCreatedBy = field.NewString(tableName, "f_created_by")
	_harbour.FUpdatedBy = field.NewString(tableName, "f_updated_by")
	_harbour.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_harbour.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_harbour.FVersion = field.NewInt64(tableName, "f_version")
	_harbour.FRemark = field.NewString(tableName, "f_remark")
	_harbour.FDeleted = field.NewInt16(tableName, "f_deleted")
	_harbour.FNameZh = field.NewString(tableName, "f_name_zh")
	_harbour.FTenantID = field.NewInt64(tableName, "f_tenant_id")
	_harbour.FNameEn = field.NewString(tableName, "f_name_en")
	_harbour.FAddressID = field.NewInt64(tableName, "f_address_id")
	_harbour.FCode = field.NewString(tableName, "f_code")
	_harbour.FLine = field.NewString(tableName, "f_line")
	_harbour.FLng = field.NewString(tableName, "f_lng")
	_harbour.FLat = field.NewString(tableName, "f_lat")

	_harbour.fillFieldMap()

	return _harbour
}

type harbour struct {
	harbourDo harbourDo

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
	FCode      field.String // 港口代码
	FLine      field.String // 航线
	FLng       field.String // 经度
	FLat       field.String // 纬度

	fieldMap map[string]field.Expr
}

func (h harbour) Table(newTableName string) *harbour {
	h.harbourDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h harbour) As(alias string) *harbour {
	h.harbourDo.DO = *(h.harbourDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *harbour) updateTableName(table string) *harbour {
	h.ALL = field.NewAsterisk(table)
	h.FID = field.NewInt64(table, "f_id")
	h.FStatus = field.NewInt16(table, "f_status")
	h.FCreatedBy = field.NewString(table, "f_created_by")
	h.FUpdatedBy = field.NewString(table, "f_updated_by")
	h.FCreatedAt = field.NewTime(table, "f_created_at")
	h.FUpdatedAt = field.NewTime(table, "f_updated_at")
	h.FVersion = field.NewInt64(table, "f_version")
	h.FRemark = field.NewString(table, "f_remark")
	h.FDeleted = field.NewInt16(table, "f_deleted")
	h.FNameZh = field.NewString(table, "f_name_zh")
	h.FTenantID = field.NewInt64(table, "f_tenant_id")
	h.FNameEn = field.NewString(table, "f_name_en")
	h.FAddressID = field.NewInt64(table, "f_address_id")
	h.FCode = field.NewString(table, "f_code")
	h.FLine = field.NewString(table, "f_line")
	h.FLng = field.NewString(table, "f_lng")
	h.FLat = field.NewString(table, "f_lat")

	h.fillFieldMap()

	return h
}

func (h *harbour) WithContext(ctx context.Context) *harbourDo { return h.harbourDo.WithContext(ctx) }

func (h harbour) TableName() string { return h.harbourDo.TableName() }

func (h harbour) Alias() string { return h.harbourDo.Alias() }

func (h *harbour) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *harbour) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 17)
	h.fieldMap["f_id"] = h.FID
	h.fieldMap["f_status"] = h.FStatus
	h.fieldMap["f_created_by"] = h.FCreatedBy
	h.fieldMap["f_updated_by"] = h.FUpdatedBy
	h.fieldMap["f_created_at"] = h.FCreatedAt
	h.fieldMap["f_updated_at"] = h.FUpdatedAt
	h.fieldMap["f_version"] = h.FVersion
	h.fieldMap["f_remark"] = h.FRemark
	h.fieldMap["f_deleted"] = h.FDeleted
	h.fieldMap["f_name_zh"] = h.FNameZh
	h.fieldMap["f_tenant_id"] = h.FTenantID
	h.fieldMap["f_name_en"] = h.FNameEn
	h.fieldMap["f_address_id"] = h.FAddressID
	h.fieldMap["f_code"] = h.FCode
	h.fieldMap["f_line"] = h.FLine
	h.fieldMap["f_lng"] = h.FLng
	h.fieldMap["f_lat"] = h.FLat
}

func (h harbour) clone(db *gorm.DB) harbour {
	h.harbourDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h harbour) replaceDB(db *gorm.DB) harbour {
	h.harbourDo.ReplaceDB(db)
	return h
}

type harbourDo struct{ gen.DO }

func (h harbourDo) Debug() *harbourDo {
	return h.withDO(h.DO.Debug())
}

func (h harbourDo) WithContext(ctx context.Context) *harbourDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h harbourDo) ReadDB() *harbourDo {
	return h.Clauses(dbresolver.Read)
}

func (h harbourDo) WriteDB() *harbourDo {
	return h.Clauses(dbresolver.Write)
}

func (h harbourDo) Session(config *gorm.Session) *harbourDo {
	return h.withDO(h.DO.Session(config))
}

func (h harbourDo) Clauses(conds ...clause.Expression) *harbourDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h harbourDo) Returning(value interface{}, columns ...string) *harbourDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h harbourDo) Not(conds ...gen.Condition) *harbourDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h harbourDo) Or(conds ...gen.Condition) *harbourDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h harbourDo) Select(conds ...field.Expr) *harbourDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h harbourDo) Where(conds ...gen.Condition) *harbourDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h harbourDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *harbourDo {
	return h.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (h harbourDo) Order(conds ...field.Expr) *harbourDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h harbourDo) Distinct(cols ...field.Expr) *harbourDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h harbourDo) Omit(cols ...field.Expr) *harbourDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h harbourDo) Join(table schema.Tabler, on ...field.Expr) *harbourDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h harbourDo) LeftJoin(table schema.Tabler, on ...field.Expr) *harbourDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h harbourDo) RightJoin(table schema.Tabler, on ...field.Expr) *harbourDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h harbourDo) Group(cols ...field.Expr) *harbourDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h harbourDo) Having(conds ...gen.Condition) *harbourDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h harbourDo) Limit(limit int) *harbourDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h harbourDo) Offset(offset int) *harbourDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h harbourDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *harbourDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h harbourDo) Unscoped() *harbourDo {
	return h.withDO(h.DO.Unscoped())
}

func (h harbourDo) Create(values ...*model.Harbour) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h harbourDo) CreateInBatches(values []*model.Harbour, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h harbourDo) Save(values ...*model.Harbour) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h harbourDo) First() (*model.Harbour, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Harbour), nil
	}
}

func (h harbourDo) Take() (*model.Harbour, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Harbour), nil
	}
}

func (h harbourDo) Last() (*model.Harbour, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Harbour), nil
	}
}

func (h harbourDo) Find() ([]*model.Harbour, error) {
	result, err := h.DO.Find()
	return result.([]*model.Harbour), err
}

func (h harbourDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Harbour, err error) {
	buf := make([]*model.Harbour, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h harbourDo) FindInBatches(result *[]*model.Harbour, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h harbourDo) Attrs(attrs ...field.AssignExpr) *harbourDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h harbourDo) Assign(attrs ...field.AssignExpr) *harbourDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h harbourDo) Joins(fields ...field.RelationField) *harbourDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h harbourDo) Preload(fields ...field.RelationField) *harbourDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h harbourDo) FirstOrInit() (*model.Harbour, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Harbour), nil
	}
}

func (h harbourDo) FirstOrCreate() (*model.Harbour, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Harbour), nil
	}
}

func (h harbourDo) FindByPage(offset int, limit int) (result []*model.Harbour, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h harbourDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h harbourDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h harbourDo) Delete(models ...*model.Harbour) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *harbourDo) withDO(do gen.Dao) *harbourDo {
	h.DO = *do.(*gen.DO)
	return h
}