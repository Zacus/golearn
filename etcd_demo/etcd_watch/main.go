package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// watch demo

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")

	defer cli.Close()

	for {
		// watch 设置一个哨兵去监看这个key的更新情况
		rch := cli.Watch(context.Background(), "ming") // <-chan WatchResponse
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
			}
		}
	}

}
