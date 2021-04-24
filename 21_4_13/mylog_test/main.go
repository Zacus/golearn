package main

import (
	"golearn/21_4_13/mylogger"
	"time"
)

var log mylogger.Logger

func main() {

	//log = mylogger.NewConsoleLog("debug")
	log = mylogger.NewFileLogger("Info", "./", "tt.log", 10*1024)
	for {

		name := "debug"
		s := 2
		log.Debug("%s[%d]\n", name, s)
		log.Info("%s[%d]\n", "info", s)
		time.Sleep(time.Second)

	}

}
