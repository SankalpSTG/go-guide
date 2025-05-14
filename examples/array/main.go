package main

import "fmt"

func main() {
	arr := [8]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(len(arr))
	fmt.Println(len(arr2))
	arr[1] = 50
	fmt.Println(arr)
}
