package appcommon

import "gorm.io/gorm"

type AppContext struct {
	Db *gorm.DB
}

func (appCtx *AppContext) GetDbConnection() *gorm.DB {
	return appCtx.Db
}

func NewAppCtx(db *gorm.DB) *AppContext {
	return &AppContext{
		Db: db,
	}
}
