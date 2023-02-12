package restapi

import (
	"io/fs"
	"net/http"
	"theitem/shared/config"
	"theitem/shared/gogen"
	"theitem/shared/infrastructure/logger"
	"theitem/web"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type controller struct {
	gogen.ControllerStarter             // built-in http that has gracefullyshutdown
	gogen.UsecaseRegisterer             // collect all the inports
	Router                  *gin.Engine // the router from preference web framework
	log                     logger.Logger
	cfg                     *config.Config
}

func NewController(appData gogen.ApplicationData, log logger.Logger, cfg *config.Config) gogen.ControllerRegisterer {

	router := gin.Default()

	// PING API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, appData)
	})

	// used for web ui development
	contentStatic, _ := fs.Sub(web.StaticFiles, "dist")
	router.StaticFS("/web", http.FS(contentStatic))

	// CORS
	router.Use(cors.New(cors.Config{
		ExposeHeaders:   []string{"Data-Length"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))

	address := cfg.Servers[appData.AppName].Address

	return &controller{
		UsecaseRegisterer: gogen.NewBaseController(),
		ControllerStarter: NewGracefullyShutdown(log, router, address),
		Router:            router,
		log:               log,
		cfg:               cfg,
	}

}
