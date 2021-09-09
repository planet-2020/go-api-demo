/**
* @author zhouhongpan
* @date 2021/9/6 15:46
 */
package mq

import (
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"io/ioutil"
	"net/http"
	"net/url"
)

var topics map[string]map[string]string

func init()  {
	// 自定义topic
	topics = make(map[string]map[string]string)
	topics["pulsar-2v2pgjqzpx9n/vini/order-create"] = map[string]string{
		"vini_dist": "https://vini-dist.sumeils.com/user_api/com/user/ucard?author_uid=49984",
		"vini_mall": "https://vini-dist.sumeils.com/user_api/com/user/ucard?author_uid=49983",
	}
	topics["pulsar-2v2pgjqzpx9n/vini/order"] = map[string]string{
		"vini_dist": "http://new.mall.vn.com/user_api/tester/mq",
	}
}

func Consume() error {
	// 连接客户端
	client, err := Connect()
	if err != nil {
		return err
	}

	// 自定义消费者
	consumers := make(map[string][]string)
	for topic, value := range topics {
		for c, _ := range value{
			consumers[c] = append(consumers[c],topic)
		}
	}

	//channels := make(map[string]chan pulsar.ConsumerMessage)
	channel := make(chan pulsar.ConsumerMessage)
	for k,v := range consumers {
		// 创建订阅者
		//channels[k] = make(chan pulsar.ConsumerMessage)
		_, err := client.Subscribe(pulsar.ConsumerOptions{
			Topics:            v,
			SubscriptionName: k,
			Type:             pulsar.Shared,
			MessageChannel: channel,
		})
		if err != nil {
			return err
		}
	}

	// 循环消费
	for cm := range channel {
		msg := cm.Message
		con := cm.Consumer
		fmt.Printf("Received message  subscription: %s -- topic: %s -- producer_name: %s -- content: '%s'\n",
			con.Subscription(), msg.Topic(), msg.ProducerName(), string(msg.Payload()))
		err := request(msg.ProducerName(), con.Subscription(), string(msg.Payload()))
		if err == nil {
			println("消费成功")
			cm.Consumer.Ack(msg)
		} else {
			println("消费失败")
		}
	}

	//for {
	//	select {
	//	case cm1 := <- channels["vini_dist"]:
	//		msg := cm1.Message
	//		fmt.Printf("vini_dist Received message msgId: %#v -- content: '%s'\n",
	//			msg.ID(), string(msg.Payload()))
	//		cm1.Consumer.Ack(msg)
	//	case cm2 := <- channels["vini_mall"]:
	//		msg := cm2.Message
	//		fmt.Printf("vini_mall Received message msgId: %#v -- content: '%s'\n",
	//			msg.ID(), string(msg.Payload()))
	//		cm2.Consumer.Ack(msg)
	//	}
	//}

	return nil
}

func request(topic, consumer, message string) error {
	urlHttp := topics[topic][consumer]
	resp, err := http.PostForm(urlHttp,url.Values{"message":{message}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}