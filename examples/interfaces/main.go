package main

import (
	"fmt"
	"time"
)

type I2DShape interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	length  int
	breadth int
}

func (m *Rectangle) Area() int {
	return m.length * m.breadth
}

func (m *Rectangle) Perimeter() int {
	return 2*m.length + 2*m.breadth
}

type Square struct {
	side int
}

func (m *Square) Area() int {
	if m == nil {
		return -1
	}
	return m.side * m.side
}

func (m *Square) Perimeter() int {
	return 4 * m.side
}

func handleInterface(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("This was an int")
	case string:
		fmt.Println("This was an int")
	case float64:
		fmt.Println("This was a float")
	default:
		fmt.Println("This was an unknown")
	}
}

func (m *Square) String() string {
	return fmt.Sprintf("I am a square with side: %v", m.side)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error at %v, %v", e.When, e.What)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &MyError{
			When: time.Now(),
			What: "Cannot divide by zero",
		}
	}
	return a / b, nil
}

func main() {

	var shape I2DShape

	shape = &Rectangle{
		length:  2,
		breadth: 4,
	}

	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())

	shape = &Square{
		side: 4,
	}

	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())

	var shape2 *Square
	shape = shape2

	fmt.Println(shape.Area())

	// var shape3 I2DShape
	// fmt.Println(shape3.Area()) // will panic

	var data interface{} = "hello"

	s := data.(string)
	fmt.Println(s)

	i, ok := data.(int)
	fmt.Println(i, ok)

	f, ok := data.(float64)
	fmt.Println(f, ok)

	// i2 := data.(int) // will throw an error
	// fmt.Println(i2)
	// f2 := data.(float64) // will throw an error
	// fmt.Println(f2)

	data = 40

	s, ok = data.(string)
	fmt.Println(s, ok)

	handleInterface("Hello")
	handleInterface(4)
	handleInterface(2.6)
	shape = &Square{
		side: 1,
	}
	handleInterface(&Square{
		side: 1,
	})
	fmt.Println(shape)

	res, err := Divide(4, 2)
	fmt.Println(res, err)

	res, err = Divide(4, 0)
	fmt.Println(res, err)
}
