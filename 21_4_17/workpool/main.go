package main

import (
	"fmt"
	"runtime"
	"time"
)

func work(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {

		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {

	runtime.GOMAXPROCS(2) //指定开启的核心数
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go work(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}

}
