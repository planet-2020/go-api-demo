package router

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/controller"
	"go-api-demo/internal/config"
	"go-api-demo/router/api"
)

type Server struct {
	GinEngine *gin.Engine
}

/**
 * @Description: 路由初始化
 * @param appConfig
 * @return *Server
 */
func Init(appConfig config.AppConfig) *Server {
	server := new(Server)
	var mode string
	if appConfig.Env == "pro" {
		mode = gin.ReleaseMode
	} else {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	server.GinEngine = gin.Default()
	server.GinEngine.Use(gin.Recovery())
	server.baseApi()	//基础路由
	apiGroup := server.GinEngine.Group("/api")
	api.ComRouter(apiGroup)	//注册com模块路由

	return server
}

/**
 * @Description: 基础路由
 * @receiver server
 */
func (server *Server) baseApi() {
	server.GinEngine.GET("/",controller.Index)
}
