package main

import "fmt"

type MyStruct struct {
	myKey int
}

func (m MyStruct) multiply(n int) int {
	return m.myKey * n
}

func main() {
	x := MyStruct{myKey: 4}
	fmt.Println(x.multiply(6))
}
