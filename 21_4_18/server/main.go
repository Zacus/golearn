package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {

	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			panic("Read failed")
		}

		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}

}

func main() {

	//创建socket
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		panic("create socket failed")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("accept failed")
		}
		go process(conn)

	}

}
