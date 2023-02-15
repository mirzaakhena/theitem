package application

import (
	"theitem/domain_item/controller/restapi"
	"theitem/domain_item/gateway/withsqlitedb"
	"theitem/domain_item/usecase/getallitem"
	"theitem/domain_item/usecase/getoneitem"
	"theitem/domain_item/usecase/runitemcreate"
	"theitem/domain_item/usecase/runitemdelete"
	"theitem/domain_item/usecase/runitempurchase"
	"theitem/domain_item/usecase/runitemupdate"
	"theitem/shared/config"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
)

type appItem struct{}

func NewAppItem() gogen.Runner {
	return &appItem{}
}

func (appItem) Run() error {

	const appName = "appItem"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	datasource := withsqlitedb.NewGateway(log, appData, cfg)
	//datasource := withmysqldb.NewGateway(log, appData, cfg)
	//datasource := withmongodb.NewGateway(log, appData, cfg)

	primaryDriver := restapi.NewController(appData, log, cfg)
	//primaryDriver := restapi2.NewController(appData, log, cfg)

	primaryDriver.AddUsecase(
		//
		getallitem.NewUsecase(datasource),
		getoneitem.NewUsecase(datasource),
		runitemcreate.NewUsecase(datasource),
		runitemdelete.NewUsecase(datasource),
		runitempurchase.NewUsecase(datasource),
		runitemupdate.NewUsecase(datasource),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
