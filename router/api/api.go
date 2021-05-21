/**
 * @Description:
 * @author zhouhongpan
 * @date 2021/5/21 9:46
 */
package api

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/controller/user_controller"
)

func InitRouter(apiGroup *gin.RouterGroup) {

	//用户
	apiGroup.POST("/user/register",user_controller.Register)

}