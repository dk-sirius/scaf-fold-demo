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
		db:                         db,
		Category:                   newCategory(db, opts...),
		Item:                       newItem(db, opts...),
		ItemVersion:                newItemVersion(db, opts...),
		Property:                   newProperty(db, opts...),
		PropertyConfig:             newPropertyConfig(db, opts...),
		PropertyConfigEnumRelation: newPropertyConfigEnumRelation(db, opts...),
		PropertyEnum:               newPropertyEnum(db, opts...),
		PropertyValue:              newPropertyValue(db, opts...),
		Sku:                        newSku(db, opts...),
		SkuInventory:               newSkuInventory(db, opts...),
		SkuInventoryLog:            newSkuInventoryLog(db, opts...),
		SkuInventoryRelation:       newSkuInventoryRelation(db, opts...),
		Spu:                        newSpu(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Category                   category
	Item                       item
	ItemVersion                itemVersion
	Property                   property
	PropertyConfig             propertyConfig
	PropertyConfigEnumRelation propertyConfigEnumRelation
	PropertyEnum               propertyEnum
	PropertyValue              propertyValue
	Sku                        sku
	SkuInventory               skuInventory
	SkuInventoryLog            skuInventoryLog
	SkuInventoryRelation       skuInventoryRelation
	Spu                        spu
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                         db,
		Category:                   q.Category.clone(db),
		Item:                       q.Item.clone(db),
		ItemVersion:                q.ItemVersion.clone(db),
		Property:                   q.Property.clone(db),
		PropertyConfig:             q.PropertyConfig.clone(db),
		PropertyConfigEnumRelation: q.PropertyConfigEnumRelation.clone(db),
		PropertyEnum:               q.PropertyEnum.clone(db),
		PropertyValue:              q.PropertyValue.clone(db),
		Sku:                        q.Sku.clone(db),
		SkuInventory:               q.SkuInventory.clone(db),
		SkuInventoryLog:            q.SkuInventoryLog.clone(db),
		SkuInventoryRelation:       q.SkuInventoryRelation.clone(db),
		Spu:                        q.Spu.clone(db),
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
		db:                         db,
		Category:                   q.Category.replaceDB(db),
		Item:                       q.Item.replaceDB(db),
		ItemVersion:                q.ItemVersion.replaceDB(db),
		Property:                   q.Property.replaceDB(db),
		PropertyConfig:             q.PropertyConfig.replaceDB(db),
		PropertyConfigEnumRelation: q.PropertyConfigEnumRelation.replaceDB(db),
		PropertyEnum:               q.PropertyEnum.replaceDB(db),
		PropertyValue:              q.PropertyValue.replaceDB(db),
		Sku:                        q.Sku.replaceDB(db),
		SkuInventory:               q.SkuInventory.replaceDB(db),
		SkuInventoryLog:            q.SkuInventoryLog.replaceDB(db),
		SkuInventoryRelation:       q.SkuInventoryRelation.replaceDB(db),
		Spu:                        q.Spu.replaceDB(db),
	}
}

type queryCtx struct {
	Category                   *categoryDo
	Item                       *itemDo
	ItemVersion                *itemVersionDo
	Property                   *propertyDo
	PropertyConfig             *propertyConfigDo
	PropertyConfigEnumRelation *propertyConfigEnumRelationDo
	PropertyEnum               *propertyEnumDo
	PropertyValue              *propertyValueDo
	Sku                        *skuDo
	SkuInventory               *skuInventoryDo
	SkuInventoryLog            *skuInventoryLogDo
	SkuInventoryRelation       *skuInventoryRelationDo
	Spu                        *spuDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Category:                   q.Category.WithContext(ctx),
		Item:                       q.Item.WithContext(ctx),
		ItemVersion:                q.ItemVersion.WithContext(ctx),
		Property:                   q.Property.WithContext(ctx),
		PropertyConfig:             q.PropertyConfig.WithContext(ctx),
		PropertyConfigEnumRelation: q.PropertyConfigEnumRelation.WithContext(ctx),
		PropertyEnum:               q.PropertyEnum.WithContext(ctx),
		PropertyValue:              q.PropertyValue.WithContext(ctx),
		Sku:                        q.Sku.WithContext(ctx),
		SkuInventory:               q.SkuInventory.WithContext(ctx),
		SkuInventoryLog:            q.SkuInventoryLog.WithContext(ctx),
		SkuInventoryRelation:       q.SkuInventoryRelation.WithContext(ctx),
		Spu:                        q.Spu.WithContext(ctx),
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
