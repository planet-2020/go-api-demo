/**
 * @Description: 默认控制器
 * @author zhouhongpan
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description: 默认首页
 * @param c
 * @author zhouhongpan
 * @date 2021-05-17 16:18:01
 */
func Index(c *gin.Context)  {
	c.String(http.StatusOK,"ok")
}
