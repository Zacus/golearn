package taillog

import (
	"fmt"
	"golearn/logagent/kafka"

	"github.com/hpcloud/tail"
)

//专门从日志文件收集日志的模块

// var (
// 	tailObj *tail.Tail
// 	LogChan chan string
// )

//TailTask:一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init() //根据路径去打开对应的日志
	return
}

// func Init(fileName string) (err error) {

// 	config := tail.Config{
// 		// File-specifc
// 		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // Seek to this location before tailing
// 		ReOpen:    true,                                 // Reopen recreated files (tail -F)
// 		MustExist: false,                                // Fail early if the file does not exist
// 		Poll:      true,                                 // Poll for file changes instead of using inotify
// 		Follow:    true,
// 	}
// 	tailObj, err = tail.TailFile(fileName, config)
// 	if err != nil {
// 		fmt.Println("tail file failed,err:", err)
// 		return
// 	}
// 	return

// }

func (t *TailTask) init() (err error) {

	config := tail.Config{
		// File-specifc
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // Seek to this location before tailing
		ReOpen:    true,                                 // Reopen recreated files (tail -F)
		MustExist: false,                                // Fail early if the file does not exist
		Poll:      true,                                 // Poll for file changes instead of using inotify
		Follow:    true,
	}
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	go t.run() //直接去收集日志发送到kafka

	return

}

func (t *TailTask) run() {

	for {
		select {
		case line := <-t.instance.Lines: //从tailObj的通道中一行一行的读取日志数据
			//发往kafka
			// kafka.SendToKafka(t.topic, line.Text)
			//先将日志数据发到一个通道中
			//kafka那个包中有单独的goroutine去取日志数据发送到kafka
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}

func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}

// func ReadChan() <-chan *tail.Line {
// 	return tailObj.Lines
// }

// func ReadLog() {
// 	var (
// 		line *tail.Line
// 		ok   bool
// 	)
// 	for {
// 		line, ok = <-tailObj.Lines
// 		if !ok {
// 			fmt.Printf("tail file close reopen,filename:%s\n", tails.Filename)
// 			time.Sleep(time.Second)
// 			continue
// 		}
// 		fmt.Println("line:", line.Text)
// 	}
// }
