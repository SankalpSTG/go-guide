# Methods in Go

Go does not have classes. But Go allows methods on Types. Any type which is defined using the type keyword can have methods.
```go
type MyStruct struct {
	myKey int
}

func (m *MyStruct) multiply(n int) int {
	return m.myKey * n
}

func main() {
	x := MyStruct{myKey: 4}
	fmt.Println(x.multiply(6))
}
```
In the multiply function, it is receiving the holding struct type as well `MyStruct` and hence it is able to use the data stored in `MyStruct`. This is similar to `this` keyword in other languages where you self reference the class using `this`.

This pattern can be used to shape a large codebase. 

Methods can be declared on any type declared with `type` keyword.

```go
type MyNum int

func (m *MyNum) multiply(n int) int {
	return int(*m) * n
}

func main() {
	x := MyNum(100)
	fmt.Println(x.multiply(4))
}
```

In the above example we created our own int type and attached a method to it.

## Normal receivers vs Pointer receivers

```go
type MyStruct struct {
	value int
}

func (m MyStruct) add(n int) {
	m.value += n
	fmt.Println(x.value)  // will print 100
}

func main() {
	x := MyStruct{
		value: 0,
	}
	x.add(100)
	fmt.Println(x.value)  // will print 0
}

```

In the above example, you can notice the `m MyStruct` in the add function is not a pointer receiver. In that case when you use the add function to add to the value, it does not throw any error and it does not add anything to the value. This is because this method does not get the original struct but a copy of the struct and hence whatever changes you do in the struct will reflect in the add function itself.

If you try to print the value in the add function, you will notice the value will be `100`. But if you try to print the value in the main function, it will print `0`

```go
type MyStruct struct {
	value int
}

func (m *MyStruct) add(n int) {
	m.value += n
	fmt.Println(x.value)  // will print 100
}

func main() {
	x := MyStruct{
		value: 0,
	}
	x.add(100)
	fmt.Println(x.value)  // will print 100
}
```

In the above code, the add function has a pointer receiver of the struct `m *MyStruct`. Hence this add function does not get a copy of the struct but the original struct itself. Hence whatever changes are done in the struct are visible outside as well.

If you print the updated value in the add function, you will see `100` and `main` function as well will see `100`

There are more use cases of using pointer receivers for structs and hence pointer receivers should be commonly used.

# Passing pointers to functions


```go
type Rectangle struct {
	length  int
	breadth int
}

func (m *Rectangle) Area() int {
	return m.length * m.breadth
}

func Perimeter(m *Rectangle) int {
	return m.length * m.breadth
}

func main() {
	rect1 := Rectangle{
		length:  4,
		breadth: 5,
	}
	fmt.Println(rect1.Area())     // will work regardless of rect1 is a pointer or a value
	fmt.Println(Perimeter(rect1)) // will throw error as rect1 is not a pointer

	rect2 := &Rectangle{
		length:  4,
		breadth: 5,
	}
	fmt.Println(rect2.Area())     // will work regardless of rect2 is a pointer or a value
	fmt.Println(Perimeter(rect2)) // will work as rect2 is a pointer
}
```
An interesting thing here is; for your convenience, golang automatically interprets a normal variable to a pointer when a type method is called. When you call Area using `rect1.Area()`, golang interprets it as `(&rect1).Area()` and hence you don't get an error.

However, when you pass a value to a function as an argument while the function expects a pointer, you will get a compile time error as we see in the case of Perimeter function.

```go
func (m *Rectangle) Add(n *Rectangle) int {
	m.length += n.length
	m.breadth += n.breadth
}
```
in the above case, this function receives a rectangle and expects a rectangle in an argument. Now if you call it like this:
```go
func main() {
	rect1 := Rectangle{
		length:  4,
		breadth: 5,
	}
  rect1.Add(Rectangle{
		length:  4,
		breadth: 5,
	}) // This will give an error
}
```

The above will give an error as the function expects a pointer. You will have to write it like this:
```go
rect1.Add(&Rectangle{
  length:  4,
  breadth: 5,
})
```

So in short, receivers are converted implicitly but arguments aren't.

## Why choose pointer receiver over value receiver

1. Pointers aren't copied but value receivers create a copy of the original value for each method call and hence in large codebase value receivers will be less efficient.
2. Pointer receivers allow you to modify the actual value.

Best practice is to either use Pointer receiver or value receiver for all methods of a type but not both.