/**
* @author zhouhongpan
* @date 2021/9/6 15:13
 */
package mq

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
)

//
//  SendProducer
//  @Description: 生产者发送消息
//  @param topic
//  @param message
//  @return error
//
func SendProducer(topic, message string) error {
	// 连接客户端
	client, err := Connect()
	if err != nil {
		return err
	}

	// 创建生产者
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
		Name: topic,
	})
	if err != nil {
		return err
	}
	defer producer.Close()

	// 发送消息
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(message),
	})
	if err != nil {
		return err
	}

	return nil
}