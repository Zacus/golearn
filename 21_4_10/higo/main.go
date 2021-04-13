package main

import "fmt"

// itoa 常量计数器
const (
	n1 = iota //1
	n2        //2
	n3        //3
	n4        //4
	n5 = 100
	n6 = iota
	n7
)

//iota 每新增一行常量声明将使itoa计数一次
const (
	l1, l2 = iota + 1, iota + 1 //1,1
	l3, l4 = iota + 1, iota + 1 //2,2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)

func main() {

	fmt.Println("hello")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n7)

	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)

	fmt.Printf("//占位符")
	m := 100
	//占位符
	fmt.Printf("%T\n", m)
	fmt.Printf("%v\n", m)
	fmt.Printf("%b\n", m)

	var s = "mingming"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

}
