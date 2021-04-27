package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

//向kafka写日志的模块

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer //声明一个连接kafka的生产者client
	logDataChan chan *logData
)

//Init 初始化client
func Init(addrs []string, maxSize int) (err error) {

	config := sarama.NewConfig()
	//tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要leader和follow确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在 sucess channel返回

	//连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed,err", err)
		return
	}
	//初始化logDataChan
	logDataChan = make(chan *logData, maxSize)
	//开启后台的goroutine从通道中拉取数据发往kafka
	go sendToKafka()
	return
}

//外部暴露的一个函数,该函数只把日志数据发送到一个内部的channel中
func SendToChan(topic, data string) {

	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg

}

//真正往kafka发送日志的函数
func sendToKafka() {

	for {
		select {
		case ld := <-logDataChan:
			//构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)

			//发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed,err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}

}

// func SendToKafka(topic, data string) {
// 	msg := &sarama.ProducerMessage{}
// 	msg.Topic = topic
// 	msg.Value = sarama.StringEncoder(data)

// 	//发送消息
// 	pid, offset, err := client.SendMessage(msg)
// 	if err != nil {
// 		fmt.Println("send msg failed,err:", err)
// 		return
// 	}
// 	fmt.Printf("pid:%v offset:%v\n", pid, offset)
// }
