package main

import (
	"fmt"
	"time"
)

//struct 当作channel消息
func main() {
	done := make(chan struct{}, 1)
	go func() {
		// 做第一个任务
		time.Sleep(time.Second)
		done <- struct{}{}
	}()
	// 另外的任务
	time.Sleep(time.Millisecond * 500)

	// 等待第一个任务完成
	<-done

	fmt.Println("完成")
}

func sliceUnique(origin []byte) (unique []byte) {
	filter := map[byte]struct{}{}
	for _, b := range origin {
		if _, ok := filter[b]; !ok {
			filter[b] = struct{}{}
			unique = append(unique, b)
		}
	}
	return
}
