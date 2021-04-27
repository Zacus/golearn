package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// etcd client put/get demo
// use etcd/clientv3

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, //地址
		DialTimeout: 5 * time.Second,            //超时时间
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"c:/tmp/nginx.log","topic":"web_log"},{"path":"d:/xxx/redis.log","topic":"redis_log"},{"path":"d:/xxx/mysql.log","topic":"mysql_log"}]`
	_, err = cli.Put(ctx, "/logagent/collect_config", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	// ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	// resp, err := cli.Get(ctx, "ming")
	// cancel()
	// if err != nil {
	// 	fmt.Printf("get from etcd failed, err:%v\n", err)
	// 	return
	// }
	// for _, ev := range resp.Kvs {
	// 	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	// }
}
