package main

import (
	"fmt"
	"net/http"
)

// http server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HelloWorld")
}

func main() {

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
