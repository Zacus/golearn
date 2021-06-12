package main

import (
	"fmt"
	"golearn/logagent/conf"
	"golearn/logagent/etcd"
	"golearn/logagent/kafka"
	"golearn/logagent/taillog"
	"golearn/logagent/utils"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

// func run() {
// 	//1.读取日志
// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			//2.发送到kafka
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}

// 	}

// }

//logAgent入口程序

func main() {
	//初始化配置文件
	var err error = nil
	err = ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Print("init config failed,err:%v\n", err)
	}
	//1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Print("init Kafka failed,err:%v\n", err)
	}
	fmt.Println("init kafka sucess")

	//2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.Timeout)*time.Second)
	if err != nil {
		fmt.Print("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd sucess")

	//实现每个logagent都拉取独有的配置，建立以自身IP地址进行分区
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	//2.1从etcd拉取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Print("get etcd.conf failed,err:%v\n", err)
		return
	}
	fmt.Printf("get etcd.conf sucess:%v\n", logEntryConf)

	//2.2设置watch去监视位置收集项目的变化及时通知logAgent实现热加载配置

	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	//2.打开日志文件准备收集日志
	//3.收集日志发往Kafka
	//3.1 循环每一个日志收集项，创建TailObj
	taillog.Init(logEntryConf)
	newConfChan := taillog.PullNewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan)
	wg.Wait()
	//发往kafka
	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Printf("Init taillog failed,err:%v\n", err)
	// 	return
	// }
	// fmt.Println("打开日志文件成功")
	//3.具体业务
	// run()

}
