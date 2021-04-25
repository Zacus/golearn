package main

import (
	"fmt"
	"golearn/logagent/conf"
	"golearn/logagent/kafka"
	"golearn/logagent/taillog"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	//1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2.发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}

//logAgent入口程序

func main() {
	//初始化配置文件
	var err error = nil
	err = ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Print("init config failed,err:%v\n", err)
	}
	//1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Print("init Kafka failed,err:%v\n", err)
	}
	fmt.Println("初始化成功")
	//2.打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("打开日志文件成功")
	run()

}
