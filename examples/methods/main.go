package main

import "fmt"

type MyStruct struct {
	value int
}

type MyStruct2 struct {
	value int
}

func (m *MyStruct) add(n int) {
	m.value += n
	fmt.Println(m.value)
}

func (m MyStruct2) add(n int) {
	m.value += n
	fmt.Println(m.value)
}

type Square struct {
	side int
}

func (m *Square) Area() int {
	return m.side * m.side
}

func Scale(m *Square, factor int) {
	m.side = m.side * factor
}

type Rectangle struct {
	length  int
	breadth int
}

func (m *Rectangle) Area() int {
	return m.length * m.breadth
}

func main() {
	x := MyStruct{
		value: 0,
	}
	x.add(100)
	y := MyStruct2{
		value: 0,
	}
	y.add(100)
	fmt.Println(x.value)
	fmt.Println(y.value)

	sqr := Square{
		side: 10,
	}
	Scale(&sqr, 10)
	fmt.Println(sqr.Area())

	rect1 := Rectangle{
		length:  4,
		breadth: 5,
	}
	fmt.Println(rect1.Area())

	rect2 := &Rectangle{
		length:  4,
		breadth: 5,
	}
	fmt.Println(rect2.Area())
}
