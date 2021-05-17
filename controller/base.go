package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description: 默认首页
 * @param c
 */
func Index(c *gin.Context)  {
	c.String(http.StatusOK,"ok")
}
