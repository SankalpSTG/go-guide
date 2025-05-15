package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 32
	var y float64 = float64(x)
	var z string = strconv.Itoa(int(y))
	fmt.Println(z)

	y = 3.23455
	z = strconv.FormatFloat(y, 'f', 2, 64)
	fmt.Println(z)

	y, err := strconv.ParseFloat(z, 64)
	if err != nil {
		fmt.Println("failed to parse float due to error,", err)
	} else {
		fmt.Println(y)
	}

	z = "1234"
	x, err = strconv.Atoi(z)
	if err != nil {
		fmt.Println("failed to parse int due to error,", err)
	} else {
		fmt.Println(x)
	}
}
