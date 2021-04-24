package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3}
	//ar1:= [...]int{0,1,2,3,4,5,6}

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println(a) // 可以输出数组[1,2,3]

	//二维数组
	arr := [2][3]int{
		{1, 2, 3},
		{2, 3, 4},
	}

	fmt.Println(arr)

	//[n]*T表示指针数组，*[n]T表示数组指针 。

	//切片slice
	var s1 []int //定义一个存在int类型元素的切片
	var s2 []string

	fmt.Println(s1, s2)

	fmt.Println(s1 == nil) //nil 空值
	fmt.Println(s2 == nil)

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"你真的", "爱过", "我吗"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//2.由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //[1,3,5,7]
	fmt.Println(s3)
	s4 := a1[0:] //[0:len(a1)]
	fmt.Println(s4)
	s5 := a1[:] //[0:len(a1)]
	fmt.Println(s5)
	//切片指向了一个底层数组
	//切片的长度是切片元素的个数，
	//切片的容量 是底层数组从切片的第一个元素道最后一个元素的数量
	s7 := a1[3:4]
	fmt.Printf("len(a1):%d len(s7):%d cap(s7):%d\n", len(a1), len(s7), cap(s7))
	fmt.Printf("len(a1):%d len(s3):%d cap(s3):%d\n", len(a1), len(s3), cap(s3))

}
