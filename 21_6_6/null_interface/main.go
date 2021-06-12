package main

import "fmt"

type Data struct {
	name string
	sex  string
}

func test(interface{}) {
	fmt.Println("interface test ")
}

//空接口的应用
//类型断言

func assgin(arg interface{}) {
	str, ok := arg.(string)
	if !ok {
		fmt.Printf("类型断言错误！\n")
	} else {
		fmt.Printf("恭喜你！猜对了；当前字符串内容为：%s\n", str)
	}
}

func assgin2(arg interface{}) {
	fmt.Printf("你输入的内容类型为：%T,", arg)
	switch t := arg.(type) {
	case string:
		fmt.Printf("内容为：%s\n", t)
	case int:
		fmt.Printf("内容为：%d\n", t)
	case bool:
		fmt.Printf("内容为：")
		fmt.Println(t)
	}
}

func main() {
	data := Data{}
	test(data)

	var temp map[string]interface{} // 声明一个空接口变量
	temp = make(map[string]interface{}, 20)
	temp["name"] = "zhangsan"
	temp["weight"] = 65
	temp["school"] = true
	temp["hobel"] = [...]string{"打球", "跑步", "codeing"}
	fmt.Println(temp)

	assgin(666)
	assgin("interface{}")

	assgin2(false)
}
