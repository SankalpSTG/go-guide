# Interfaces

Interfaces are defined as a set of method signatures.

```go
type I2DShape interface{
  Area() int
  Perimeter() int
}
```

the interface can then be implemented by any other type

```go
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
	return m.side * m.side
}

func (m *Square) Perimeter() int {
	return 4 * m.side
}

```
In the above code, we have created two structs, a rectangle and a square. both implement the two functions, `Area` and `Perimeter` of the I2DShape interface in their own way. 

Usually in other languages, you would implement an interface explicitly as shown below:
```typescript
class Rectangle implements I2DShape {}
```

However in the case of Golang, we don't need to explicitly mention that we are implementing the interface. Now look at the below code:
```go
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
}

```

In the above code, we created a variable which is of type I2DShape. However while assigning a value to it, we are able to assign both Rectangle and Square to it. Go checks if the method signatures of the interface are implemented in the value type correctly, then it won't throw an error. But in case if any of the shape had any method missing, go would throw an error. 

However extra methods are can be defined for both types Rectangle and Square.

Interfaces are like tuples
```go
(value, type)
```
where you define a type in the interface and the value has to be provided with implementation.
## Interface values with nil underlying values

```go
func main(){
  var shape2 *Square
	shape = shape2

	fmt.Println(shape.Area())
}
```

In the above case, we created a shape2 variable but didn't assign a value. And then we assigned it to shape. This is fine however when you call the below Area function
```go
func (m *Square) Area() int {
	return m.side * m.side
}
```

The square receiver i.e. variable `m` will be nil. So its not the `shape.Area()` that will cause error but the statement `return m.side * m.side` will be causing error as you are referencing a value of `nil` pointer. To fix this, we can modify the Area function for square as follows:
```go
func (m *Square) Area() int {
	if m == nil {
		return -1
	}
	return m.side * m.side
}
```

In the above case, we are checking if m is nil. If it is then we will return -1. Else we will return the area. This will not cause an error.

## Nil interface value

```go
func main(){
  var shape3 I2DShape
	fmt.Println(shape3.Area())
}
```

The above will cause a runtime error as there is no value attached to the interface. 

## Empty interface

```go
func main(){
  var shape4 interface{}
}
```

The above variable `shape4` can hold any value, any methods but will not hint for any methods as the interface signature itself does not have any methods defined. This can be used to handle values of unknown type.

```go
func main(){
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
}
```
In the above code, we assigned a string to interface{} type. To typecast data to a type, we can use `data.(type)` as we used for `int`, `float64` and `string`. To __safely__ typecast, we can expect one more variable to the left of type cast, `ok` in our case which tells if the typecast was successful or not. In case you only write one variable and the typecast fails, it will cause a panic.

You can run a switch case to handle different types of values

## Stringers

if you add one method to the square struct
```go
func (m *Square) String() string {
	return fmt.Sprintf("I am a square with side: %v", m.side)
}
```

Look at the method signature. The name is `String()` and the return type is `string`. This method is used to describe itself or the type it belongs to. In our case, we are describing what type of value it is and what is the side length.

Now if you write below line in main

```go
func main(){
	shape = &Square{
		side: 1,
	}
	fmt.Println(shape)
}
```

Notice the line `fmt.Println(shape)`, this will internally call the String method that we defined above which will return the formatted string and the same string will be printed in the terminal.

```
I am a square with side: 1
```

This is because the `fmt` package has an interface defined as follows
```go
type Stringer interface{
  String() string
} 
```

Hence if the passed arguments have the `String()` function, it will use it.

## Errors

Similar to stringer, there is an error interface built-in

```go
type error interface {
  Error() string
}
```

you can create your own errors by implementing this interface and packages like the `fmt` will use the Error() method to print your custom error.

```go
type MyError struct {
  When time.Time
  What string
}

func (e *MyError) Error() string {
  return fmt.Sprintf("Error at %v, %s", e.When, e.What)
}
```
Now all you need to do is instantiate the struct and use it.

```go
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &MyError{
			When: time.Now(),
			What: "Cannot divide by zero",
		}
	}
	return a / b, nil
}
```

When you use it

```go
func main(){
  res, err = Divide(4, 0)
	fmt.Println(res, err)
}
```
The error will be printed as defined.