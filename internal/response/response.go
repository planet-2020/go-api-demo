/**
 * @Description: 接口返回
 * @author zhouhongpan
 * @date 2021/5/19 11:09
 */
package response

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/internal/code"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

/**
 * @Description: 接口返回结果
 * @param c
 * @param err
 * @param data
 * @author zhouhongpan
 * @date 2021-05-19 11:23:38
 */
func ApiResponse(c *gin.Context, err error, data interface{})  {
	code, message := code.DecodeErr(err)
	c.JSON(http.StatusOK,Response{
		Code: code,
		Message: message,
		Data: data,
	})
}