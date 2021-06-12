package main

import "fmt"

func main() {
	fmt.Printf("%s", sliceUnique([]byte("aabbcc")))
	// output: abc
}

func sliceUnique(origin []byte) (unique []byte) {
	//value是空结构体，构造集合
	filter := map[byte]struct{}{}
	for _, b := range origin {
		if _, ok := filter[b]; !ok {
			filter[b] = struct{}{}
			unique = append(unique, b)
		}
	}
	return
}
