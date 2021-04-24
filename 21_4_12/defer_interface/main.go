package main

import "fmt"

func test() (x int) {

	defer func() {
		x++
	}()
	return 0
}

func test1() (x int) {

	x = 1
	defer func() {
		x++
	}()
	return x
}

func test2() (x int) {

	x = 1
	defer func(x int) {
		x++
	}(x)
	return x
}

func test3(x int) (y int) {

	defer func() {
		x++
	}()
	return x
}

func main() {

	fmt.Println(test()) //1
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3(3))
}
