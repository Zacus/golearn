package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic"
)

type DataToStr struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *DataToStr
)

//初始化ES 准备接收kafka那发来的数据

func Init(address string, chanSize, nums int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return err
	}

	fmt.Println("connect to es sucess")
	ch = make(chan *DataToStr, chanSize)
	for i := 0; i < nums; i++ {
		go SendToES()
	}

	return

}

func SendToESChan(msg *DataToStr) {
	ch <- msg
}

//发送数据到ES
func SendToES() error {

	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				//panic(err)
				fmt.Println(err)
				continue
			}
			fmt.Printf("Indexed user %s to index %s,type %s\n", put1.Id, put1.Index, put1.Type)
			// return err
		default:
			time.Sleep(time.Second)
		}
	}

}
