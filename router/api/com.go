package api

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/controller/com"
)

/**
 * @Description: 社区路由
 * @param apiGroup
 */
func ComRouter(apiGroup *gin.RouterGroup) {
	group := apiGroup.Group("/com")
	// 聊天
	group.GET("/chat/list",com.ListChat)

}
