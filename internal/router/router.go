package router

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/config"
	"os"
	"runtime"
)

func RouteSetup() *gin.Engine {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if os.Getenv(config.ENVkEY) == config.PROD {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	api := r.Group("/api")
	uploadRouter(api)
	return r
}
