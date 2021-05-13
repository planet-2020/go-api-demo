package com

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListChat(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Code int `json:"code"`
		Message string `json:"message"`
		Data interface{} `json:"data"`
	}{
		Code: 0,
		Message: "测试内容",
		Data: map[string]string{"abc":"aaa","def":"ddd"},
	})
}
