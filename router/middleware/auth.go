/**
 * @Description: token鉴权
 * @author zhouhongpan
 * @date 2021/5/25 10:40
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/internal/code"
	"go-api-demo/internal/config"
	"go-api-demo/internal/response"
	"go-api-demo/pkg/token"
)

/**
 * @Description: 鉴权中间件
 * @return gin.HandlerFunc
 * @author zhouhongpan
 * @date 2021-05-25 10:54:30
 */
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("authorization")
		if len(tokenString) == 0 {
			c.Abort()
			response.ApiResponse(c,code.ErrTokenNeed,nil)
			return
		}
		tokenString = tokenString[7:]	//截去开头的 Bearer
		claims, err := token.ParseToken(config.Conf.App.JwtSecret,tokenString)
		if err != nil {
			c.Abort()
			response.ApiResponse(c,code.ErrToken,nil)
			return
		}
		c.Set("uid",claims.UserID)

		c.Next()
	}
}