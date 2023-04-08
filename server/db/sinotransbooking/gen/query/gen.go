// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		ActionLog:          newActionLog(db, opts...),
		Attachment:         newAttachment(db, opts...),
		Bail:               newBail(db, opts...),
		Commission:         newCommission(db, opts...),
		CompanyCode:        newCompanyCode(db, opts...),
		ContentMap:         newContentMap(db, opts...),
		DetailTemplate:     newDetailTemplate(db, opts...),
		Harbor:             newHarbor(db, opts...),
		Node:               newNode(db, opts...),
		PerformanceGoods:   newPerformanceGoods(db, opts...),
		PerformanceInfo:    newPerformanceInfo(db, opts...),
		Route:              newRoute(db, opts...),
		ShippingConditions: newShippingConditions(db, opts...),
		ShippingSchedule:   newShippingSchedule(db, opts...),
		SurchargeTemplate:  newSurchargeTemplate(db, opts...),
		Whitelist:          newWhitelist(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	ActionLog          actionLog
	Attachment         attachment
	Bail               bail
	Commission         commission
	CompanyCode        companyCode
	ContentMap         contentMap
	DetailTemplate     detailTemplate
	Harbor             harbor
	Node               node
	PerformanceGoods   performanceGoods
	PerformanceInfo    performanceInfo
	Route              route
	ShippingConditions shippingConditions
	ShippingSchedule   shippingSchedule
	SurchargeTemplate  surchargeTemplate
	Whitelist          whitelist
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		ActionLog:          q.ActionLog.clone(db),
		Attachment:         q.Attachment.clone(db),
		Bail:               q.Bail.clone(db),
		Commission:         q.Commission.clone(db),
		CompanyCode:        q.CompanyCode.clone(db),
		ContentMap:         q.ContentMap.clone(db),
		DetailTemplate:     q.DetailTemplate.clone(db),
		Harbor:             q.Harbor.clone(db),
		Node:               q.Node.clone(db),
		PerformanceGoods:   q.PerformanceGoods.clone(db),
		PerformanceInfo:    q.PerformanceInfo.clone(db),
		Route:              q.Route.clone(db),
		ShippingConditions: q.ShippingConditions.clone(db),
		ShippingSchedule:   q.ShippingSchedule.clone(db),
		SurchargeTemplate:  q.SurchargeTemplate.clone(db),
		Whitelist:          q.Whitelist.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		ActionLog:          q.ActionLog.replaceDB(db),
		Attachment:         q.Attachment.replaceDB(db),
		Bail:               q.Bail.replaceDB(db),
		Commission:         q.Commission.replaceDB(db),
		CompanyCode:        q.CompanyCode.replaceDB(db),
		ContentMap:         q.ContentMap.replaceDB(db),
		DetailTemplate:     q.DetailTemplate.replaceDB(db),
		Harbor:             q.Harbor.replaceDB(db),
		Node:               q.Node.replaceDB(db),
		PerformanceGoods:   q.PerformanceGoods.replaceDB(db),
		PerformanceInfo:    q.PerformanceInfo.replaceDB(db),
		Route:              q.Route.replaceDB(db),
		ShippingConditions: q.ShippingConditions.replaceDB(db),
		ShippingSchedule:   q.ShippingSchedule.replaceDB(db),
		SurchargeTemplate:  q.SurchargeTemplate.replaceDB(db),
		Whitelist:          q.Whitelist.replaceDB(db),
	}
}

type queryCtx struct {
	ActionLog          *actionLogDo
	Attachment         *attachmentDo
	Bail               *bailDo
	Commission         *commissionDo
	CompanyCode        *companyCodeDo
	ContentMap         *contentMapDo
	DetailTemplate     *detailTemplateDo
	Harbor             *harborDo
	Node               *nodeDo
	PerformanceGoods   *performanceGoodsDo
	PerformanceInfo    *performanceInfoDo
	Route              *routeDo
	ShippingConditions *shippingConditionsDo
	ShippingSchedule   *shippingScheduleDo
	SurchargeTemplate  *surchargeTemplateDo
	Whitelist          *whitelistDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		ActionLog:          q.ActionLog.WithContext(ctx),
		Attachment:         q.Attachment.WithContext(ctx),
		Bail:               q.Bail.WithContext(ctx),
		Commission:         q.Commission.WithContext(ctx),
		CompanyCode:        q.CompanyCode.WithContext(ctx),
		ContentMap:         q.ContentMap.WithContext(ctx),
		DetailTemplate:     q.DetailTemplate.WithContext(ctx),
		Harbor:             q.Harbor.WithContext(ctx),
		Node:               q.Node.WithContext(ctx),
		PerformanceGoods:   q.PerformanceGoods.WithContext(ctx),
		PerformanceInfo:    q.PerformanceInfo.WithContext(ctx),
		Route:              q.Route.WithContext(ctx),
		ShippingConditions: q.ShippingConditions.WithContext(ctx),
		ShippingSchedule:   q.ShippingSchedule.WithContext(ctx),
		SurchargeTemplate:  q.SurchargeTemplate.WithContext(ctx),
		Whitelist:          q.Whitelist.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
