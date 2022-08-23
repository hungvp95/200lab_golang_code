package component

import "gorm.io/gorm"

type AppContext interface {
	GetDBConnection() *gorm.DB
}

type appComponent struct {
	db *gorm.DB
}

func CreateAppComponent(db *gorm.DB) *appComponent {
	return &appComponent{db}
}

func (app *appComponent) GetDBConnection() *gorm.DB {
	return app.db
}