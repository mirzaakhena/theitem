package withmysqldb

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"theitem/domain_item/gateway/shared"
	"theitem/domain_item/model/entity"
	"theitem/shared/config"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
	"time"
)

type gateway struct {
	*shared.GormImpl
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	cdb := cfg.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cdb.Username, cdb.Password, cdb.Host, cdb.Port, cdb.DBName)

	log.Info(context.TODO(), "connection string : %s", dsn)

	var db *gorm.DB
	var err error

	retryConnect := 10
	for {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			retryConnect--
			if retryConnect == 0 {
				panic(err)
			}
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}

	if err = db.AutoMigrate(&entity.Item{}); err != nil {
		panic(err)
	}

	return &gateway{
		GormImpl: shared.NewGormImpl(appData, cfg, log, db),
	}
}
