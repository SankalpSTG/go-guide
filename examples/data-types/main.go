package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 200 //defaults to int32 or int64
	var y int8 = -100
	// var z uint = -100 // invalid
	var z uint = 100

	fmt.Println(x, y, z)

	var m rune = -42
	var n byte = 37
	fmt.Println(reflect.TypeOf(m), reflect.TypeOf(n))

	var u float64 = 3.44
	fmt.Println(u)

	a := complex(1.2, 3.4)
	var b complex128 = 5.6 + 7i
	c := 8 + 7i

	fmt.Println(a, b, c)

	fmt.Println(a + b)

	fmt.Println(real(a))

	fmt.Println(reflect.TypeOf(a))

	fmt.Println(reflect.TypeOf(x))
}
