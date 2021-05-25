/**
 * @Description:
 * @author zhouhongpan
 * @date 2021/5/20 17:22
 */
package user_controller

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/internal/code"
	"go-api-demo/internal/response"
	"go-api-demo/internal/validate"
	"go-api-demo/service/user_service"
)

/**
 * @Description: 用户注册
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

/**
 * @Description: 用户登录
 * @param c
 * @author zhouhongpan
 * @date 2021-05-21 15:31:47
 */
func Login(c *gin.Context)  {
	var loginService user_service.LoginService
	if err := c.ShouldBind(&loginService); err != nil {
		response.ApiResponse(c,validate.GetValidateErr(err),nil)
		return
	}
	token, err := loginService.Login()
	if  err != nil {
		response.ApiResponse(c,err,nil)
		return
	}
	response.ApiResponse(c,nil,token)
}

/**
 * @Description: 获取用户信息
 * @param c
 * @author zhouhongpan
 * @date 2021-05-25 11:44:21
 */
func UserInfo(c *gin.Context)  {
	uid, ok := c.Keys["uid"]
	if !ok {
		response.ApiResponse(c,code.ErrTokenNeed,nil)
		return
	}
	var userInfo user_service.UserInfo
	err := userInfo.Get(uid)
	if err != nil {
		response.ApiResponse(c,err,nil)
		return
	}
	response.ApiResponse(c,nil,userInfo)
}

/**
 * @Description: 更新用户
 * @param c
 * @author zhouhongpan
 * @date 2021-05-25 14:05:25
 */
func Update(c *gin.Context)  {
	uid, ok := c.Keys["uid"]
	if !ok {
		response.ApiResponse(c,code.ErrTokenNeed,nil)
		return
	}
	var updateService user_service.UpdateService
	if err := c.ShouldBind(&updateService); err != nil {
		response.ApiResponse(c,validate.GetValidateErr(err),nil)
		return
	}
	userInfo, err := updateService.Update(uid)
	if err != nil {
		response.ApiResponse(c,err,nil)
		return
	}
	response.ApiResponse(c,nil,userInfo)
}

/**
 * @Description: 注销用户
 * @param c
 * @author zhouhongpan
 * @date 2021-05-25 14:22:05
 */
func Delete(c *gin.Context)  {
	uid, ok := c.Keys["uid"]
	if !ok {
		response.ApiResponse(c,code.ErrTokenNeed,nil)
		return
	}
	if err := user_service.Delete(uid); err != nil {
		response.ApiResponse(c,err,nil)
		return
	}
	response.ApiResponse(c,nil,nil)
}