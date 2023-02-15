package restapi2

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"theitem/shared/config"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
)

type controller struct {
	//gogen.ControllerStarter            // built-in http that has gracefullyshutdown
	gogen.UsecaseRegisterer            // collect all the inports
	Router                  *echo.Echo // the router from preference web framework
	log                     logger.Logger
	cfg                     *config.Config
	address                 string
}

func NewController(appData gogen.ApplicationData, log logger.Logger, cfg *config.Config) gogen.ControllerRegisterer {

	router := echo.New()

	// PING API
	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, appData)
	})

	// CORS
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		ExposeHeaders: []string{"Data-Length"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowHeaders:  []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	address := cfg.Servers[appData.AppName].Address

	return &controller{
		UsecaseRegisterer: gogen.NewBaseController(),
		//ControllerStarter: NewGracefullyShutdown(log, router, address),
		Router:  router,
		log:     log,
		cfg:     cfg,
		address: address,
	}

}

func (r *controller) Start() {
	err := r.Router.Start(r.address)
	if err != nil {
		panic(err)
	}
}
