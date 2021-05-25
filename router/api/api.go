/**
 * @Description:
 * @author zhouhongpan
 * @date 2021/5/21 9:46
 */
package api

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/controller/user_controller"
	"go-api-demo/router/middleware"
)

func InitRouter(apiGroup *gin.RouterGroup) {

	//用户
	apiGroup.POST("/user/register",user_controller.Register)
	apiGroup.POST("/user/login",user_controller.Login)

	//授权分组
	apiAuthGroup(apiGroup)
}

func apiAuthGroup(authGroup *gin.RouterGroup)  {
	authGroup.Use(middleware.AuthMiddleware())
	//用户
	authGroup.GET("/user/info",user_controller.UserInfo)
	authGroup.POST("/user/update",user_controller.Update)
	authGroup.POST("/user/delete",user_controller.Delete)
}

