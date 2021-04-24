package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// t1, err := os.OpenFile("golearn/21_4_12/file_op/t1.text", os.O_RDONLY, 4)
	// if err != nil {
	// 	fmt.Println("file open failed")
	// 	return

	// }
	// defer t1.Close()

	/*
		tmp := make([]byte, 256)
		n, err := t1.Read(tmp)
		if err == io.EOF {
			fmt.Println("file read end")
			return
		}

		if err != nil {

			fmt.Println("read failed")
			return
		}

		fmt.Printf("read %d byte data\n", n)
		fmt.Println(string(tmp[:n]))
	*/

	//bufio读取文件
	/*
		reader := bufio.NewReader(t1)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				if len(line) != 0 {
					fmt.Println(line)
				}
				fmt.Println("file read end")
				return
			}
			if err != nil {

				fmt.Println("read failed")
				return
			}
			fmt.Println(line)
		}
	*/

	//读取整个文件
	//io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。
	content, err := ioutil.ReadFile("golearn/21_4_12/file_op/t1.text")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))

}
