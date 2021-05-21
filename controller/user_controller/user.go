/**
 * @Description:
 * @author zhouhongpan
 * @date 2021/5/20 17:22
 */
package user_controller

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/internal/response"
	"go-api-demo/internal/validate"
	"go-api-demo/service/user_service"
)

/**
 * @Description: 会员注册
 * @param c
 * @author zhouhongpan
 * @date 2021-05-21 09:45:33
 */
func Register(c *gin.Context)  {
	var registerService user_service.RegisterService
	if err := c.ShouldBind(&registerService); err != nil {
		response.ApiResponse(c, validate.GetValidateErr(err),nil)
		return
	}
	if err := registerService.Register(); err != nil {
		response.ApiResponse(c,err,nil)
		return
	}

	response.ApiResponse(c,nil,registerService)
}