package component

import "gorm.io/gorm"

type appCtx struct {
	db *gorm.DB
}

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}
