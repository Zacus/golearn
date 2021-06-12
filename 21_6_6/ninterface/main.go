package main

import "fmt"

type Sayer interface {
	say()
}

type Cat struct {
}

func (cat Cat) say() {
	fmt.Println("妙妙妙")
}

type Dog struct {
}

func (dog Dog) say() {
	fmt.Println("旺旺旺")
}

func main() {
	var x Sayer
	tom := Cat{}
	x = tom
	x.say()
	snoopy := Dog{}
	x = snoopy
	x.say()
}
