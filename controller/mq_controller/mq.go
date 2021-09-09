package mq_controller

import (
	"github.com/gin-gonic/gin"
	"go-api-demo/internal/code"
	"go-api-demo/internal/response"
	"go-api-demo/pkg/mq"
)

//
//  Send
//  @Description: 发送消息
//  @param c
//
func Send(c *gin.Context)  {
	// 接收参数
	topic := c.PostForm("topic")
	message := c.PostForm("message")
	if topic == "" {
		response.ApiResponse(c,code.ErrMqSendNotTopic,nil)
		return
	}
	if message == "" {
		response.ApiResponse(c,code.ErrMqSendNotMessage,nil)
		return
	}

	// 发送消息
	if err := mq.SendProducer(topic,message); err != nil {
		response.ApiResponse(c,code.ErrMqSendFail,nil)
		return
	}

	response.ApiResponse(c,nil,nil)
	return
}
