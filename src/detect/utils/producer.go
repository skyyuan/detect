package utils

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
	"github.com/astaxie/beego"
	"log"
)

var producer *nsq.Producer

// 初始化生产者
func InitProducer() {
	str := beego.AppConfig.String("nsqProducerAddress")
	var err error
	fmt.Println("address: ", str)
	producer, err = nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		fmt.Println("Can't new producer")
	}
}

// 关闭
func CloseProducer(){
	producer.Stop()
}

//发布消息
func ProducerPublish(topic string, message string) error {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Publish(topic, []byte(message))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
	//var err error
	//if producer != nil {
	//	if message == "" { //不能发布空串，否则会导致error
	//		return nil
	//	}
	//	fmt.Println(message)
	//	err = producer.Publish(topic, []byte(message)) // 发布消息
	//	return err
	//}
	return fmt.Errorf("producer is nil", err)
}

//发布消息
func ProducerPublishDelay(topic string, message string, delay time.Duration) error {
	str := beego.AppConfig.String("nsqProducerAddress")
	var err error
	fmt.Println("address: ", str)
	producer, err = nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		fmt.Println("Can't new producer")
	}
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		err = producer.DeferredPublish(topic, delay, []byte(message)) // 发布消息
		if err != nil {
			fmt.Println(11111)
			fmt.Println(err)
		}
		producer.Stop()
		return err

	}

	return fmt.Errorf("producer is nil", err)
}