/**
* @author zhouhongpan
* @date 2021/9/6 15:13
 */
package mq

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"sync"
	"time"
)

var client pulsar.Client
var once sync.Once

//
//  Connect
//  @Description: 连接客户端（单例）
//  @return pulsar.Client
//  @return error
//
func Connect() (pulsar.Client, error) {
	var err error
	once.Do(func() {
		println("单例连接TDMQ客户端")
		client, err = pulsar.NewClient(pulsar.ClientOptions{
			URL:               "http://pulsar-2v2pgjqzpx9n.tdmq.ap-gz.public.tencenttdmq.com:8080", //更换为接入点地址（控制台集群管理页完整复制）
			Authentication:    pulsar.NewAuthenticationToken("eyJrZXlJZCI6InB1bHNhci0ydjJwZ2pxenB4OW4iLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwdWxzYXItMnYycGdqcXpweDluX3Rlc3RfYWRtaW4ifQ.BByceoq9r6Q7jEMJcv0xisl8WlvU5jzXGiOrkYNk5-U"), //更换为密钥
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})
	})
	if err != nil {
		return nil,err
	}
	return client,nil
}

