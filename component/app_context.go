package component

import (
	"food-delivery-200lab/component/uploadprovider"
	"gorm.io/gorm"
)

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}
