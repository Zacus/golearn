package main

import "fmt"

func main() {

	//goto+label
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {

			if j == 'Q' {
				goto gotoTag
			}
			fmt.Printf("%v-%c\n", i, j)
		}

	}
gotoTag:
	fmt.Println("over")

}
