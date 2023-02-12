package withsqlitedb

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"theitem/domain_item/gateway/shared"
	"theitem/domain_item/model/entity"
	"theitem/shared/config"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
)

type gateway struct {
	*shared.GormImpl
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	db, err := gorm.Open(sqlite.Open(cfg.Database.DBName), &gorm.Config{})

	if err = db.AutoMigrate(&entity.Item{}); err != nil {
		panic(err)
	}

	return &gateway{
		GormImpl: shared.NewGormImpl(appData, cfg, log, db),
	}
}
