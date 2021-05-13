package api

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/controller/com"
)

func ComRouter(apiGroup *gin.RouterGroup) {
	group := apiGroup.Group("/com")
	// 聊天
	group.GET("/chat/list",com.ListChat)

}
