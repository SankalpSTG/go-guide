package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	for i := 1; i <= 10; {
		fmt.Println(i)
		i += 2
	}
	i := 1
	for i < 10 {
		fmt.Println(i)
		i += 3
	}
	x := []string{"I", "Am", "A", "Legend"}
	for index, value := range x {
		fmt.Println(index, value)
	}
	for _, value := range x {
		fmt.Println(value)
	}
}
