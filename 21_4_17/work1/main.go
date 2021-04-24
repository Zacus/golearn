package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
	1.使用goroutine和channel实现一个计算int64随机数各位数和的程序。
		开启一个goroutine循环生成int64类型的随机数，发送到jobChan
		开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
		主goroutine从resultChan取出结果并打印到终端输出
*/

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

var wg sync.WaitGroup

type job struct {
	value int64
}

type result struct {
	jobs   *job
	result int64
}

//生成int64类型的随机数
func rand64(n chan<- *job) {
	defer wg.Done()
	//rand.Seed(time.Now().UnixNano())
	for {
		r1 := rand.Int63()
		newJob := &job{
			value: r1,
		}
		n <- newJob
	}

}

//取出随机数计算各位数的和
func randSum(n <-chan *job, results chan<- *result) {

	defer wg.Done()
	for {
		num := <-n
		var sum int64 = 0
		va := num.value
		for va > 0 {
			sum += va % 10
			va = va / 10
		}
		newResult := &result{
			jobs:   num,
			result: sum,
		}

		results <- newResult

	}

}

func main() {
	wg.Add(1)
	go rand64(jobChan)

	wg.Add(24)
	for i := 1; i <= 24; i++ {
		go randSum(jobChan, resultChan)
	}

	for resultset := range resultChan {

		fmt.Printf("value:%d sum:%d\n", resultset.jobs.value, resultset.result)
	}
	wg.Wait()
}
