package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

//专门从日志文件收集日志的模块

var (
	tailObj *tail.Tail
	LogChan chan string
)

func Init(fileName string) (err error) {

	config := tail.Config{
		// File-specifc
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // Seek to this location before tailing
		ReOpen:    true,                                 // Reopen recreated files (tail -F)
		MustExist: false,                                // Fail early if the file does not exist
		Poll:      true,                                 // Poll for file changes instead of using inotify
		Follow:    true,
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	return

}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}

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
