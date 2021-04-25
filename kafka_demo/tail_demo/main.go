package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {

	fileName := "./my.log"
	config := tail.Config{
		// File-specifc
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // Seek to this location before tailing
		ReOpen:    true,                                 // Reopen recreated files (tail -F)
		MustExist: false,                                // Fail early if the file does not exist
		Poll:      true,                                 // Poll for file changes instead of using inotify
		Follow:    true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
	}

	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen,filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
