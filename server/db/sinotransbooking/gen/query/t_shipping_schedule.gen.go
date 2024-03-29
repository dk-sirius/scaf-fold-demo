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

func newShippingSchedule(db *gorm.DB, opts ...gen.DOOption) shippingSchedule {
	_shippingSchedule := shippingSchedule{}

	_shippingSchedule.shippingScheduleDo.UseDB(db, opts...)
	_shippingSchedule.shippingScheduleDo.UseModel(&model.ShippingSchedule{})

	tableName := _shippingSchedule.shippingScheduleDo.TableName()
	_shippingSchedule.ALL = field.NewAsterisk(tableName)
	_shippingSchedule.ID = field.NewInt64(tableName, "id")
	_shippingSchedule.FScac = field.NewString(tableName, "f_scac")
	_shippingSchedule.FCarrierName = field.NewString(tableName, "f_carrier_name")
	_shippingSchedule.FServiceName = field.NewString(tableName, "f_service_name")
	_shippingSchedule.FVesselName = field.NewString(tableName, "f_vessel_name")
	_shippingSchedule.FVoyageNumber = field.NewString(tableName, "f_voyage_number")
	_shippingSchedule.FImoNumber = field.NewString(tableName, "f_imo_number")
	_shippingSchedule.FOriginUnloc = field.NewString(tableName, "f_origin_unloc")
	_shippingSchedule.FOriginCountry = field.NewString(tableName, "f_origin_country")
	_shippingSchedule.FOriginCityName = field.NewString(tableName, "f_origin_city_name")
	_shippingSchedule.FOriginSubdivision = field.NewString(tableName, "f_origin_subdivision")
	_shippingSchedule.FDestinationUnloc = field.NewString(tableName, "f_destination_unloc")
	_shippingSchedule.FDestinationCountry = field.NewString(tableName, "f_destination_country")
	_shippingSchedule.FDestinationSubdivision = field.NewString(tableName, "f_destination_subdivision")
	_shippingSchedule.FDestinationCityName = field.NewString(tableName, "f_destination_city_name")
	_shippingSchedule.FTotalDuration = field.NewInt32(tableName, "f_total_duration")
	_shippingSchedule.FScheduleType = field.NewString(tableName, "f_schedule_type")
	_shippingSchedule.FDeleted = field.NewInt16(tableName, "f_deleted")
	_shippingSchedule.FType = field.NewInt16(tableName, "f_type")
	_shippingSchedule.FOriginDepartureDate = field.NewTime(tableName, "f_origin_departure_date")
	_shippingSchedule.FDestinationArrivalDate = field.NewTime(tableName, "f_destination_arrival_date")
	_shippingSchedule.FEstimatedTerminalCutoff = field.NewTime(tableName, "f_estimated_terminal_cutoff")
	_shippingSchedule.FCreatedAt = field.NewTime(tableName, "f_created_at")
	_shippingSchedule.FUpdatedAt = field.NewTime(tableName, "f_updated_at")
	_shippingSchedule.FDataJSON = field.NewString(tableName, "f_data_json")
	_shippingSchedule.FRouteCode = field.NewString(tableName, "f_route_code")
	_shippingSchedule.FTerminalCutoff = field.NewTime(tableName, "f_terminal_cutoff")
	_shippingSchedule.FWayID = field.NewString(tableName, "f_way_id")
	_shippingSchedule.FEstimatedOriginDepartureDate = field.NewTime(tableName, "f_estimated_origin_departure_date")
	_shippingSchedule.FDeletedAt = field.NewTime(tableName, "f_deleted_at")

	_shippingSchedule.fillFieldMap()

	return _shippingSchedule
}

type shippingSchedule struct {
	shippingScheduleDo shippingScheduleDo

	ALL                           field.Asterisk
	ID                            field.Int64  // 自增主键
	FScac                         field.String // 船公司简称
	FCarrierName                  field.String // 船公司
	FServiceName                  field.String // 服务名
	FVesselName                   field.String // 船名
	FVoyageNumber                 field.String // 航次
	FImoNumber                    field.String // imo号
	FOriginUnloc                  field.String // 起始港
	FOriginCountry                field.String // 起始国家code
	FOriginCityName               field.String // 起始城市名
	FOriginSubdivision            field.String // 起始省份
	FDestinationUnloc             field.String // 目的港code
	FDestinationCountry           field.String // 目的国家
	FDestinationSubdivision       field.String // 目的省
	FDestinationCityName          field.String // 目的城市名
	FTotalDuration                field.Int32  // 船期
	FScheduleType                 field.String // 船期类型
	FDeleted                      field.Int16  // 是否删除 0 未删除 1 已删除
	FType                         field.Int16  // 数据类型 0 inttra 数据 1 手动录入数据 2 自动生成数据 3飞拓
	FOriginDepartureDate          field.Time   // 离港时间
	FDestinationArrivalDate       field.Time   // 到港时间
	FEstimatedTerminalCutoff      field.Time   // 预估截港时间
	FCreatedAt                    field.Time   // 创建时间
	FUpdatedAt                    field.Time   // 更新时间
	FDataJSON                     field.String // 第三方字段
	FRouteCode                    field.String // 航线代码
	FTerminalCutoff               field.Time   // 最终截港时间
	FWayID                        field.String // 飞拓wayid
	FEstimatedOriginDepartureDate field.Time   // 预估离港时间
	FDeletedAt                    field.Time

	fieldMap map[string]field.Expr
}

func (s shippingSchedule) Table(newTableName string) *shippingSchedule {
	s.shippingScheduleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s shippingSchedule) As(alias string) *shippingSchedule {
	s.shippingScheduleDo.DO = *(s.shippingScheduleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *shippingSchedule) updateTableName(table string) *shippingSchedule {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.FScac = field.NewString(table, "f_scac")
	s.FCarrierName = field.NewString(table, "f_carrier_name")
	s.FServiceName = field.NewString(table, "f_service_name")
	s.FVesselName = field.NewString(table, "f_vessel_name")
	s.FVoyageNumber = field.NewString(table, "f_voyage_number")
	s.FImoNumber = field.NewString(table, "f_imo_number")
	s.FOriginUnloc = field.NewString(table, "f_origin_unloc")
	s.FOriginCountry = field.NewString(table, "f_origin_country")
	s.FOriginCityName = field.NewString(table, "f_origin_city_name")
	s.FOriginSubdivision = field.NewString(table, "f_origin_subdivision")
	s.FDestinationUnloc = field.NewString(table, "f_destination_unloc")
	s.FDestinationCountry = field.NewString(table, "f_destination_country")
	s.FDestinationSubdivision = field.NewString(table, "f_destination_subdivision")
	s.FDestinationCityName = field.NewString(table, "f_destination_city_name")
	s.FTotalDuration = field.NewInt32(table, "f_total_duration")
	s.FScheduleType = field.NewString(table, "f_schedule_type")
	s.FDeleted = field.NewInt16(table, "f_deleted")
	s.FType = field.NewInt16(table, "f_type")
	s.FOriginDepartureDate = field.NewTime(table, "f_origin_departure_date")
	s.FDestinationArrivalDate = field.NewTime(table, "f_destination_arrival_date")
	s.FEstimatedTerminalCutoff = field.NewTime(table, "f_estimated_terminal_cutoff")
	s.FCreatedAt = field.NewTime(table, "f_created_at")
	s.FUpdatedAt = field.NewTime(table, "f_updated_at")
	s.FDataJSON = field.NewString(table, "f_data_json")
	s.FRouteCode = field.NewString(table, "f_route_code")
	s.FTerminalCutoff = field.NewTime(table, "f_terminal_cutoff")
	s.FWayID = field.NewString(table, "f_way_id")
	s.FEstimatedOriginDepartureDate = field.NewTime(table, "f_estimated_origin_departure_date")
	s.FDeletedAt = field.NewTime(table, "f_deleted_at")

	s.fillFieldMap()

	return s
}

func (s *shippingSchedule) WithContext(ctx context.Context) *shippingScheduleDo {
	return s.shippingScheduleDo.WithContext(ctx)
}

func (s shippingSchedule) TableName() string { return s.shippingScheduleDo.TableName() }

func (s shippingSchedule) Alias() string { return s.shippingScheduleDo.Alias() }

func (s *shippingSchedule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *shippingSchedule) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 30)
	s.fieldMap["id"] = s.ID
	s.fieldMap["f_scac"] = s.FScac
	s.fieldMap["f_carrier_name"] = s.FCarrierName
	s.fieldMap["f_service_name"] = s.FServiceName
	s.fieldMap["f_vessel_name"] = s.FVesselName
	s.fieldMap["f_voyage_number"] = s.FVoyageNumber
	s.fieldMap["f_imo_number"] = s.FImoNumber
	s.fieldMap["f_origin_unloc"] = s.FOriginUnloc
	s.fieldMap["f_origin_country"] = s.FOriginCountry
	s.fieldMap["f_origin_city_name"] = s.FOriginCityName
	s.fieldMap["f_origin_subdivision"] = s.FOriginSubdivision
	s.fieldMap["f_destination_unloc"] = s.FDestinationUnloc
	s.fieldMap["f_destination_country"] = s.FDestinationCountry
	s.fieldMap["f_destination_subdivision"] = s.FDestinationSubdivision
	s.fieldMap["f_destination_city_name"] = s.FDestinationCityName
	s.fieldMap["f_total_duration"] = s.FTotalDuration
	s.fieldMap["f_schedule_type"] = s.FScheduleType
	s.fieldMap["f_deleted"] = s.FDeleted
	s.fieldMap["f_type"] = s.FType
	s.fieldMap["f_origin_departure_date"] = s.FOriginDepartureDate
	s.fieldMap["f_destination_arrival_date"] = s.FDestinationArrivalDate
	s.fieldMap["f_estimated_terminal_cutoff"] = s.FEstimatedTerminalCutoff
	s.fieldMap["f_created_at"] = s.FCreatedAt
	s.fieldMap["f_updated_at"] = s.FUpdatedAt
	s.fieldMap["f_data_json"] = s.FDataJSON
	s.fieldMap["f_route_code"] = s.FRouteCode
	s.fieldMap["f_terminal_cutoff"] = s.FTerminalCutoff
	s.fieldMap["f_way_id"] = s.FWayID
	s.fieldMap["f_estimated_origin_departure_date"] = s.FEstimatedOriginDepartureDate
	s.fieldMap["f_deleted_at"] = s.FDeletedAt
}

func (s shippingSchedule) clone(db *gorm.DB) shippingSchedule {
	s.shippingScheduleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s shippingSchedule) replaceDB(db *gorm.DB) shippingSchedule {
	s.shippingScheduleDo.ReplaceDB(db)
	return s
}

type shippingScheduleDo struct{ gen.DO }

func (s shippingScheduleDo) Debug() *shippingScheduleDo {
	return s.withDO(s.DO.Debug())
}

func (s shippingScheduleDo) WithContext(ctx context.Context) *shippingScheduleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s shippingScheduleDo) ReadDB() *shippingScheduleDo {
	return s.Clauses(dbresolver.Read)
}

func (s shippingScheduleDo) WriteDB() *shippingScheduleDo {
	return s.Clauses(dbresolver.Write)
}

func (s shippingScheduleDo) Session(config *gorm.Session) *shippingScheduleDo {
	return s.withDO(s.DO.Session(config))
}

func (s shippingScheduleDo) Clauses(conds ...clause.Expression) *shippingScheduleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s shippingScheduleDo) Returning(value interface{}, columns ...string) *shippingScheduleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s shippingScheduleDo) Not(conds ...gen.Condition) *shippingScheduleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s shippingScheduleDo) Or(conds ...gen.Condition) *shippingScheduleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s shippingScheduleDo) Select(conds ...field.Expr) *shippingScheduleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s shippingScheduleDo) Where(conds ...gen.Condition) *shippingScheduleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s shippingScheduleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *shippingScheduleDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s shippingScheduleDo) Order(conds ...field.Expr) *shippingScheduleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s shippingScheduleDo) Distinct(cols ...field.Expr) *shippingScheduleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s shippingScheduleDo) Omit(cols ...field.Expr) *shippingScheduleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s shippingScheduleDo) Join(table schema.Tabler, on ...field.Expr) *shippingScheduleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s shippingScheduleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *shippingScheduleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s shippingScheduleDo) RightJoin(table schema.Tabler, on ...field.Expr) *shippingScheduleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s shippingScheduleDo) Group(cols ...field.Expr) *shippingScheduleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s shippingScheduleDo) Having(conds ...gen.Condition) *shippingScheduleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s shippingScheduleDo) Limit(limit int) *shippingScheduleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s shippingScheduleDo) Offset(offset int) *shippingScheduleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s shippingScheduleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *shippingScheduleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s shippingScheduleDo) Unscoped() *shippingScheduleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s shippingScheduleDo) Create(values ...*model.ShippingSchedule) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s shippingScheduleDo) CreateInBatches(values []*model.ShippingSchedule, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s shippingScheduleDo) Save(values ...*model.ShippingSchedule) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s shippingScheduleDo) First() (*model.ShippingSchedule, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingSchedule), nil
	}
}

func (s shippingScheduleDo) Take() (*model.ShippingSchedule, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingSchedule), nil
	}
}

func (s shippingScheduleDo) Last() (*model.ShippingSchedule, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingSchedule), nil
	}
}

func (s shippingScheduleDo) Find() ([]*model.ShippingSchedule, error) {
	result, err := s.DO.Find()
	return result.([]*model.ShippingSchedule), err
}

func (s shippingScheduleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ShippingSchedule, err error) {
	buf := make([]*model.ShippingSchedule, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s shippingScheduleDo) FindInBatches(result *[]*model.ShippingSchedule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s shippingScheduleDo) Attrs(attrs ...field.AssignExpr) *shippingScheduleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s shippingScheduleDo) Assign(attrs ...field.AssignExpr) *shippingScheduleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s shippingScheduleDo) Joins(fields ...field.RelationField) *shippingScheduleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s shippingScheduleDo) Preload(fields ...field.RelationField) *shippingScheduleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s shippingScheduleDo) FirstOrInit() (*model.ShippingSchedule, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingSchedule), nil
	}
}

func (s shippingScheduleDo) FirstOrCreate() (*model.ShippingSchedule, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingSchedule), nil
	}
}

func (s shippingScheduleDo) FindByPage(offset int, limit int) (result []*model.ShippingSchedule, count int64, err error) {
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

func (s shippingScheduleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s shippingScheduleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s shippingScheduleDo) Delete(models ...*model.ShippingSchedule) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *shippingScheduleDo) withDO(do gen.Dao) *shippingScheduleDo {
	s.DO = *do.(*gen.DO)
	return s
}
