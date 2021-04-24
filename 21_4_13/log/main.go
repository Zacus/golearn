package main

import (
	"log"
	"os"
	"time"
)

func main() {

	log.Println("this is error")
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[info]")
	log.Println("this is a big error")

	tmp := "syslog"
	tmp += time.Now().Format("20060102150405")
	tmp += ".log"

	logfile, err := os.OpenFile("golearn/21_4_13/log/"+tmp, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("open file error")
	}

	//f, err := logfile.Stat()
	//defer logfile.Close()

	log.SetOutput(logfile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[info]")
	log.Println("this is a big error")

	logger := log.New(os.Stdout, "<info>", log.Lshortfile|log.Lmicroseconds|log.Ldate)
	logger.Println("this logger info")

	logfile.Close()

	os.Rename("golearn/21_4_13/log/"+tmp, "golearn/21_4_13/log/"+tmp+"bak")

}
