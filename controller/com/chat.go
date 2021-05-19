package com

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/internal/response"
)

func ListChat(c *gin.Context) {

	response.ApiResponse(c,nil,"测试response")
}
