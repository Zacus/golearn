package main

import (
	"fmt"
	"golearn/log_transfer/conf"
	"golearn/log_transfer/es"
	"golearn/log_transfer/kafka"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.LogTransferCfg)
)

//log transfer
//将日志数据从kafka取出来发往ES

func main() {
	//0.加载配置文件
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Print("init config failed,err:%v\n", err)
	}

	//初始化ES
	//初始化一个ES连接的client
	//对外提供一个往ES写入数据 的一个函数
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.CoreNums)
	if err != nil {
		fmt.Printf("init ES failed,err:%v\n", err)
	}

	//1.初始化kafka
	//连接kafka,创建分区的消费者
	//每个分区的消费者分别取出数据 通过SendToES发送给ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Print("init Kafka failed,err:%v\n", err)
	}
	fmt.Println("init kafka sucess")

	//2.从kafka取日志数据

	//3.发往ES
}
