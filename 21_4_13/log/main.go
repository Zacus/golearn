package main

import (
	"log"
	"os"
)

func main() {

	log.Println("this is error")
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[info]")
	log.Println("this is a big error")

	logfile, err := os.OpenFile("golearn/21_4_13/log/1.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("open file error")
	}
	defer logfile.Close()

	log.SetOutput(logfile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[info]")
	log.Println("this is a big error")

	logger := log.New(os.Stdout, "<info>", log.Lshortfile|log.Lmicroseconds|log.Ldate)
	logger.Println("this logger info")

	//logrus zap

}
