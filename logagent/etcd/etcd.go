package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	cli *clientv3.Client
)

//需要收集的日志的配置文件
type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

//初始化etcd
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr}, //地址
		DialTimeout: timeout,        //超时时间
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

//从etcd中根据key获取配置项
func GetConf(key string) (v []*LogEntry, err error) {

	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &v)
		if err != nil {
			fmt.Printf("Unmarshal conf failed,err:%v\n", err)
			return
		}
	}
	return
}

func WatchConf(key string, ch chan<- []*LogEntry) {
	for {
		// watch 设置一个哨兵去监看这个key的更新情况
		rch := cli.Watch(context.Background(), key) // <-chan WatchResponse
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
				//通知taillog.tskMgr
				//1.判断操作的类型
				var newConf []*LogEntry
				if ev.Type != clientv3.EventTypeDelete {
					//删除
					err := json.Unmarshal(ev.Kv.Value, &newConf)
					if err != nil {
						fmt.Printf("Unmarshal failed,err:%v\n", err)
						continue
					}
				}

				fmt.Printf("get new conf:%v\n", newConf)
				ch <- newConf
			}
		}
	}
}
