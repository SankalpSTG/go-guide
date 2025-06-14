# Go Guide

This guide is a quick learning guide for Go Programming Language. Prerequisites: you should have good knowledge of other programming languages.

## Table Of Contents
1. [Hello World](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#hello-world)
2. [Variables in Go](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#variables-in-go)
3. [Data Types in Go](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#data-types-in-go)
	1. [Bool](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#bool)
	2. [Integers](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#integers)
	3. [Float](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#float)
	4. [String](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#string)
4. [Arrays in Go](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#arrays-in-go)
5. [Slice in Go](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#slice-in-go)
	1. [Adding Slice to Slice](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#adding-slice-to-slice)
	2. [Slicing a Slice](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#slicing-a-slice)
	3. [Copying a Slice](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#copying-a-slice)
6. [Make Function](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#make-function)
6. [Operators](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#operators-in-go)
6. [Conditions](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#conditions-in-go)
6. [Switch Case](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#switch-case-in-go)
6. [Loops](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#loops-in-go)
6. [Functions](https://github.com/SankalpSTG/go-guide?tab=readme-ov-file#functions-in-golang)
## Hello World

To initialize a project, you need to write the below command in your project root folder
```
go mod init <project-name>
```
Then you can create a ```main.go``` file. This will be your project entry point. In the file, write the below code.
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
In this code we have declared the file as a part of the package main. The file also has a function declared with the help of ```func``` keyword. We use the ```fmt``` package to print the message on the console.

To execute the above code, write in terminal
```
go run main.go
```
OR
```
go run .
```
To view all commands available in go, you can use below command
```
go help
```
## Variables in Go
Variables in go can be declared using either of following ways
```go
var x = 5
y := 10
```
The ```:=``` is a shortcut operator to declare variables, we will discuss more on this later. If you wish to declare variable with types you can also do that. However you cannot use ``:=`` outside of a function.
```go
var x int
```
This will declare a variable without initializing. You can also declare / initialize variables in bulk.
```go
var (
	c = 10
	d = 20
	e = 30
)
fmt.Println(c, d, e)
var (
	f int
	g int
	h int
)
fmt.Println(f, g, h)
var i, j, k int
i = 1
j = 2
k = 3
fmt.Println(i, j, k)
var l, m, n int = 4, 5, 6
fmt.Println(l, m, n)
o, p, q := 1, 2, 3
fmt.Println(o, p, q)
```
### Constants
Go also supports constants
```go
const x int = 5
```
Go does not allow the ``:=`` operator to define constants
## Data Types in Go
Unlike languages like Python and Javascript, Go is statically typed. That means once a variable type is defined, the variable will only store that type of data.

Go has 3 basic types
1. bool
2. Numeric
3. string

### Bool
To declare a variable and then initalize it, you can do it following way:
```go
var myVar bool 					//type declaration
myVar = false 					//value initialization
var myVar2 bool = false //type declaration and initialization
```

There is a special short cut operator to declare and define variables in Go:
```go
myVar := false
```
In the above code, using the ```:=``` operator, we can directly declare and assign the variable a value. The type is automatically inferred from the value.

### Integers

There are two types of integers
1. Signed integers
2. Unsigned integers

Signed integers are declared with the int prefixed keywords
```go
var x int = 200
var y int8 = -100 
```
There are five types of int:
1. int -> 32 bits for 32 bit systems and 64 bits for 64 bit systems. 
2. int8 -> 8 bit integer storing from -128 to 127
3. int16 -> 16 bit integer storing from -32768 to 32767
4. int32 -> 32 bit integer storing from +/- 2^31
5. int64 -> 64 bit integer storing from +/- 2^63

You can use these keywords to define integers.

Unsigned integers can be defined using uint prefixed key
```go
var x uint = 1200
var x uint16 = 2456
```
Similar to int, you have five types of uint
1. int -> 32 bits for 32 bit systems and 64 bits for 64 bit systems.
2. int8 -> 8 bit storing positive values from 0 to 2^8
2. int16 -> 16 bit storing positive values from 0 to 2^16
2. int32 -> 32 bit storing positive values from 0 to 2^32
2. int64 -> 64 bit storing positive values from 0 to 2^64

There are two aliases for uint8 and int32
1. btye -> alias for uint8
2. rune -> alias for int32

```go
var m rune = -42
var n byte = 37
fmt.Println(reflect.TypeOf(m), reflect.TypeOf(n))
```
``m`` will be ``int32`` whereas n will be ``uint8``

### Float
There are two types of floats
1. float32 -> 32 bit float storing from -3.4e+38 to 3.4e+38
1. float64 -> 64 bit float storing from -1.7e+308 to 1.7e+308

The default type is float64 in case you use ```:=``` for defining a variable.
```go
x := 3.44				//float64
var y float32 = 3e+38
var y float64 = 17.323
```

### String

Strings can be defined using the ```string``` keyword
```go
var x string = "Hello, World!"
y := "My name is awesome!"
```
### Complex
Go also supports complex numbers as follows
```go
a := complex(1.2, 3.4)
var b complex128 := complex(1.2, 3.4)
c := 8 + 7i
```
Complex numbers is the same mathematical concept where you had a real and an imaginary number. Both the real and imaginary part are floats either float32 or float64.

Complex numbers can be either 128bit or 64bit
1. complex128 -> 64bit f real + 64bit imaginary
1. complex64 -> 32bit real + 32bit imaginary

### Zero Values
Following are the zero values, which are given when variables are declared without initializing
1. ``0`` for numeric types
2. ``false`` for boolean type
3. ``""`` for strings


### Null
In Golang, ```null``` is replaced with ```nil```. We will talk more about ``nil`` when discussing pointers and interfaces.

## Type Conversions
You can use type functions to convert between types
1. int to float using ``float32(int)`` or ``float64(int)``
2. float to int using ``int(float)`` which is lossy, or ``math.Round(float)`` or ``math.Floor(float)`` or ``math.Ceil(float)`` 
3. float to string using strconv.FormatFloat(float, decimalFormat, precision, bits) for e.g. strconv.FormatFloat(f, "f", 2, 64). Bits can be 32 or 64. Other way is fmt.Sprintf("$.2f", float) where 2 is the precision. 
4. int to string using strconv.Itoa(int) and string to int using strconv.Atoi(string)
5. string to float using strconv.ParseFloat(string, bits) for e.g. strconv.ParseFloat(2, 64). Bits can be 32 or 64.

All the examples can be found in the ``examples/type-conversion`` folder. strconv has additional functions to parse other types like bool and complex.

There are some other type conversions for advanced types like array, struct, interface, maps which we will cover later.
## Arrays in Go

You can declare arrays in go using following syntax
```go
var myArr = [4]int{1, 2, 3, 4}                //fixed length
var myArr2 = [...]int{1, 2, 3, 4, 5}          //fixed length but dynamic length calculation
var myArr2 = [8]int{1, 2, 3, 4}               //next 4 values will be zero
var myArr2 = [5]string{"I", "am", "awesome"}  //next 2 values will be empty strings ""
var myArr2 = [2]int{1, 2, 3, 4}               //this will cause an error
```
As seen above, arrays have a fixed length which can be mentioned while declaration or can be inferred from values while initialization.

Accessing values is similar to other programming languages, using indexes. Arrays are 0-indexed.
```go
myArr := [...]int{1, 2, 3, 4}
fmt.Println(myArr[2]) //this will print 3
```
Values can also be changed using the indexes
```go
myArr := [...]int{1, 2, 3, 4}
myArr[2] = 5
fmt.Println(myArr) //[1, 2, 5, 4]
```
Initializing specific indexes of array:
```go
myArr := [5]int{2: 10, 3: 50} //[0, 0, 10, 50, 0]
```
Finding the length of the array
```go
myArr := [...]int{1, 2, 3, 4}
myArr2 := [2]int{1, 2}
fmt.Println(len(myArr)) //4
fmt.Println(len(myArr2)) //2
```
## Slice in Go
Slices in simple words are dynamic arrays. Slices provide you with modularity over arrays. Creating a slice is similar to creating an array. You just don't provide the length of the array.
```go
var myArr []int = []int{1, 2, 3}
myArr2 := []int{1, 2, 3, 4, 5}
fmt.Println(myArr)
fmt.Println(len(myArr2))
fmt.Println(cap(myArr2))
```
In the above example, you can see the cap function which returns the capacity that the slice can hold currently as per the memory alloted to it. This as well is dynamic and grows with the length of the array.

To append values to a slice, you can use the append function as follows:
```go
arr := []int(1, 2, 3, 4)
arr = append(arr, 5)
```
Below are some more examples
```go
arr := []int{1, 2, 3, 4}
fmt.Println(arr, len(arr), cap(arr))
arr = append(arr, 5)
fmt.Println(arr, len(arr), cap(arr))
arr2 := []int{1, 2, 3}
fmt.Println(arr2, len(arr2), cap(arr2))
arr2 = append(arr2, 4) // Capacity changes to 6 (2x)
// arr2 = append(arr2, 4, 5)             // Capacity changes to 6 (2x)
// arr2 = append(arr2, 4, 5, 6)          // Capacity changes to 6 (2x)
// arr2 = append(arr2, 4, 5, 6, 7)       // Capacity changes to 8 (2x + 2)
// arr2 = append(arr2, 4, 5, 6, 7, 8, 9) // Capacity changes to 8 (2x + 2 + 2)
fmt.Println(arr2, len(arr2), cap(arr2))
fmt.Println(arr2, len(arr2), cap(arr2))
```
In the above examples, I noticed one thing. When you append to an array for which the length and capacity are equal, it doubles the capacity of the array.

When the elements that are being added are more than the current length of the array and the length and capacity are equal, the capacity is incremented in a different way.

If the length and capacity of array is 3, and you are adding 6 elements, so the final length of array will be 9.

But capacity will be incremented as follows:
1. First it will double the capacity. so 3 becomes 6.
2. Can 9 elements fit in the capacity of 6, no... Hence now the capacity is further increased by 2 so the new capacity becomes 8.
3. Can 9 elements fit in the capacity of 8, no... hence now the capacity is further increased by 2 so the new capacity becomes 10.
4. Can 9 elements fit in the capacity of 10, yes. So this is the final capacity.

I observed that capacity first increases by 2 x original capacity and then is incremented by 2.

But once the append operation is done, and the length has gone from 3 to 9 and the capacity has become 10, if you now append again but just one value, then the capacity remains 10 and length becomes 10.

However instead of one value, if you increment 2 values, the length becomes 11, which cannot fit in the capacity of 10. So now the capacity is first doubled again from 10 to 20.

So the new length will be 11 and the new capacity will be 20.

### Adding Slice to Slice
```go
arr2 = append(arr2, arr...) //adding slice to another slice
fmt.Println(arr2, len(arr2), cap(arr2))
```
The above example shows how you can add a slice to a slice using the ```...``` operator.

### Slicing a Slice
```go
arr3 := arr2[5:7]
// arr3 := arr2[:7] // will slice from start upto 7
// arr3 := arr2[5:] // will slice from 5 till end
// arr3 := arr2[:] // will slice from start to end
fmt.Println(arr, arr2, arr3)
arr3[1] = 25 //Updates original array (arr2) as well but (arr) won't change
fmt.Println(arr, arr2, arr3)
```
In the above example, we created a shorter version of the slice from index 5 to 7(not inclusive of 7). However the arr3 elements have the same reference as of arr2 and hence when you make any change in arr3, arr2 gets affected.

```go
arr3 = append(arr3, 82, 83, 84)
fmt.Println(arr, arr2, arr3)
```
If you further append into arr3 like above, the effect will be as follows
1. The length of arr2 is 9 while arr3 is indexed from 5th upto 7th index of arr2. The length of the arr3 is 2 but the capacity is 7 because you sliced upto 7 even though you took only 2 elements. The memory referencing will be as follows
```
mem  0  1  2  3  4  5  6  7  8  9  10
arr2 1  2  3  4  1  2  25 4  5
arr3                2  25 -  -  -  -  -
```
2. In the above example, the ```-``` represents memory reserved for arr3 elements. So you can notice that the memory reservation also starts from the 5th index of arr2. So basically for arr3, the index 7 and 8 of arr2 are actually reserved for arr3 but empty in arr3's perspective. But for arr2, those actually hold values.
3. When you append 3 elements 82, 83 and 84... 82 and 83 take the 7th and 8th index of arr2 so the 7th and 8th index of arr2 are updated as well
4. Now arr3 has memory reserved beyond the memory of arr2. So 84 takes next position to 83. But since arr2 does not have that memory reserved for itself, 84 does not reflect in arr2.

So the initial and final state of arr2 and arr3 will be:
```
initial
arr2 = [1, 2, 3, 4, 1, 2, 25, 4, 5]
arr3 = [2, 25]

Final
arr2 = [1, 2, 3, 4, 1, 2, 25, 82, 83]
arr3 = [2, 25, 82, 83, 84]
```

### Copying a slice
If you wish that the above should not happen, then you can copy slices with their memory addresses as well using the copy function.
```go
myarr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
myarr2 := myarr1[5:7]
var myarr2Copy = make([]int, len(myarr2))
copy(myarr2Copy, myarr2)
```
Here we used 2 functions, make and copy. We will discuss more about make later. But for now, make is a way to declare variables as well as allows to declare the length and capacity as well.

We declared a ```myarr2Copy``` variable which has same length as of myarr2. Then we used the copy function to copy values from ```myarr2``` to ```myarr2Copy```. Since we have already declared the ```myarr2Copy``` variable, it's address will be different as of ```myarr2``` or ```myarr1```.

So now if you append to ```myarr2Copy```
```go
fmt.Println(myarr2Copy, myarr2, myarr1)
myarr2Copy = append(myarr2Copy, 9, 10, 11)
fmt.Println(myarr2Copy, myarr2, myarr1)
```
This should not update myarr2 or myarr1 and only myarr2Copy will be appended.

## Make Function
Make function let's you declare variables just like var. There are some differences though. Make function has 3 parameters
```
make(type, length, capacity)
```
The type is the data type to be declared, the length and capacity can be passed in case of slices while for other types, length and capacity is not needed.

Make does not work for primitive types like bool, int and string, but works for slices, struct and map.
```go
var x []int
x[0] = 10 //this will cause an error

x := make([]int, 4)
x[0] = 10 //this will work
```
In the above example, we declared x as []int however with ```var```, we cannot provide the length of the slice. And hence the length of slice will remain zero. So in that case if you index the 0th index that will cause an error.

However, if you see the ```make``` example, we are passing 4 as the length. So x will be a slice with initial length of 4. So it will be safe to access the 0th index.

### For Primitive Data types
There is an additional way to decalare variables. Look at the below example
```go
x := new(int) // intialized as 0
fmt.Println(*x)
y := new(string) // initialized as ""
fmt.Println(*y)
z := new([4]int) // initialized as [0,0,0,0]
fmt.Println(*z)
```
We can use the new keyword for allocating memory as per data type. However the variables will be pointer variables. We will discuss about pointers later.

## Operators in Go
Go supports most of the common operators that you see in other languages
```
Arithmetic:  +, -, *, /, %, ++, --
Assignment:  =, +=, -=, *=, /=, %=, &=, |=, ^=, >>=, <<=
Comparison:  ==, !=, >, <, >=, <=
Logical:     &&, ||, !
Bitwise:     &, |, ^, <<, >>
```

## Conditions in Go
Writing if else statement remains same as other languages, however we don't use the (round brackets).
```go
x= 2
if x % 2 == 0{
	fmt.Println("X is Even")
	return
}
fmt.Println("X is Odd")
```
Another example
```go
if x % 2 == 0{
	fmt.Println("X is Even")
}else{
	fmt.Println("X is Odd")
}
```
If else chaining
```go
if x == 0{
	fmt.Println("X is Zero")
} else if x % 2 == 0{
	fmt.Println("X is Even")
}else{
	fmt.Println("X is Odd")
}
```
Nested if statement
```go
if x % 2 == 0{
	if x == 0{
		fmt.Println("X is Zero")
		return
	}
	fmt.Println("X is Even")
}else{
	fmt.Println("X is Odd")
}
```
If condition can also have a short statement to execute before the function
```go
y := 4
z := 5
if x:= y * z; x % 2 == 0 {
	//code for condition
}
```
Since x is defined inside the if condition, the scope of x is limited to if condition and is not accessible outside.
## Switch Case in Go
Go has switch cases but go has made it better. Look at following example
```go
x := 1
switch x{
	case 1:
		fmt.Println("X is 1")
	case 2:
		fmt.Println("X is 2")
	default:
		fmt.Println("I don't know what X is")
}
```
Notice in the above example, we are not writing break keyword anywhere, as switch cases in go automatically break after the case is finished unlike some other languages.

Switch cases in go also support expressions and functions as well
```go
x := 2
y := 4
switch y{
	case x * 2:
		fmt.Println("Y is two times X")
	case someFunctions():
		fmt.Println("Y is two times X")
	default:
		fmt.Println("Default Case")
}
```
Another Example
```go
x := 2
y := 4
switch y{
	case "4": //this will cause error
		fmt.Println("y is 4")
	default:
		fmt.Println("Default Case")
}
```
In the above case, case "4" will cause an error as all case expressions should have the same type as of the switch expression.

Swith in go also supports multiple values
```go
day := 5
switch y{
	case 1, 3, 5: //this will cause error
		fmt.Println("Day is odd")
	case 0, 2, 4, 6:
		fmt.Println("Day is even")
}
```
You can write switch case without condition as well
```go
switch {
	case x + y == 0:
		//do something
	case x < y:
		//do something
	case x == y:
		//do something
}
```
Here, we basically mean ``switch true{}`` so whichever case expression returns true will be executed. However the sequence of execution check will remain from top to bottom.

## Loops in Go
Go has for loops that act as while loops as well. For loops takes (up to) three statements
```
for statement1; statement2; statement3{
	//write code here
}
``` 
So a normal for loop would look like
```go
for i := 1; i <= 10; i++ {
	fmt.Println(i)
}
```
A while loop would look like
```go
i := 1
for i < 10 {
	fmt.Println(i)
	i += 3
}
```
you can use ```continue``` and ```break``` statements to control the loop
```go
for i := 1; i <= 10; i++ {
	if i > 5{
		break
	} else if i % 2 == 0{
		continue
	}
	fmt.Println(i)
}
```
To iterate over arrays and slices and maps, you can use for loop as below
```
for index, value := range arrat | slice | map{
	//code
}
```
An example
```go
x := []string{"I", "Am", "A", "Legend"}
for index, value := range x {
	fmt.Println(index, value)
}
```
In case you only need values or only indexes
```go
for _, value := range x {
	fmt.Println(value)
}
for index, _ := range x {
	fmt.Println(index)
}
```
Running a loop forever
```go
for {
	//This loop will run forever unless it reaches break keyword.
}
```
## Functions in Golang
To define functions, we use the ``func`` keyword as follows
```go
func functionName(){
	//code
}
```
To call the function you would write the function name with the round brackets as you do in other languages
```go
func myFunction(){
	fmt.Println("myFunction was called")
}
func main(){
	myFunction()
}
```
Following are some rules for function names
1. A function name must start with a letter
2. A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
3. Function names are case-sensitive
4. A function name cannot contain spaces
5. If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used

### Arguments to functions
A function can ask for data, you can provide data to the function through arguments.
```
func functionName(arg1 type, arg2 type, arg3 type){
	//code
}
```
An example
```go
func greetings(greeting string, name string, age int){
	fmt.Println(greeting, name, "your age is", age)
}
func main(){
	greetings("Welcome", "Rahul", 34)
}
```
You can also write the function parameters combining the types as follows
```go
func greetings(greeting, name string, age string){
	fmt.Println(greeting, name, "your age is", age)
}
```
Both greeting and name will have type string

Functions can return value as well
```go
func add(a, b int) int {
	return a + b
}
func main(){
	result := add(4, 5)
	fmt.Println(result)
}
```
Functions can also return multiple values
```go
func moodyAdder(a, b int)(string, int){
	if a % 2 == 0 && b % 2 == 0{
		return "Happy", a + b
	}else{
		return "Sad", a - b
	}
}
func main(){
	mood, result = moodyAdder(5, 4)
	fmt.Println(mood, result)
}
```
Return values can also be named
```go
func moodyAdder(a, b int)(mood string, total int){
	if a % 2 == 0 && b % 2 == 0{
		mood = "Happy"
		total = a + b
	}else{
		mood = "Sad"
		total = a - b
	}
	return
	// OR
	// return mood, total
}
```
While the caller will still call it normally
```go
func main(){
	mood, result = moodyAdder(5, 4)
	fmt.Println(mood, result)
}
```
Whenever you don't want a particular value you can substitute it with ``_``
```go
func main(){
	mood, _ = moodyAdder(5, 4)
	fmt.Println(mood)
}
```
Go also supports recursion
```go
func factorial(n int) int {
	if n == 0 || n == 1{
		return n
	}
	return n * factorial(n - 1)
}
```
### Defer keyword
``defer`` delays the execution of a function until the surrounding function returns.
```go
func main(){
	defer fmt.Println("World!")
	fmt.Println("Hello")
}
```
Here, world! will be printed after Hello. This is useful in cases where you wish to keep a database connected until the function returns.
```go
func main(){
	defer fmt.Println("World!")
	defer fmt.Println("Hello")
}
```
``defer`` functions are pushed into a stack. When function returns, the deferred functions are called in LIFO fashion. In this case, Hello will be printed first and then World!.

### Functions are values too
You can store function in a variable and pass it to other functions
```go
func execute(a, b int, fn func(int, int) int)int{
	return fn(a, b)
}
func main(){
	add := func(a, b int) int{
		return a + b
	}
	subtract := func(a, b int) int{
		return a - b
	}
	multiply := func(a, b int) int{
		return a * b
	}
	fmt.Println(execute(4, 2, add))
	fmt.Println(execute(4, 2, subtract))
	fmt.Println(execute(4, 2, multiply))
}
``` 

### Closures
When you write a function within another function as below
```go
func testSum() func() int {
	a, b := 5, 10
	return func() int {
		return a + b
	}
}
```
The outer function returns the inner function and the outer function is done. However the inner function still keeps the scope of the outer function. That means, the inner function that is being returned, still has access to a and b, can update it and can use it as long as the function remains in use.
```go
x := testSum()
fmt.Println(x())
```

## Struct in Go
Structs are the most interesting & important part in Golang. Struct is a non-primitive Data Type which allows you to store multiple members of different data types into a single variable. Below is the syntax of defining it
```
type struct_name struct{
	member1 datatype;
	member2 datatype;
	member3 datatype;
	member4 datatype;
}
``` 
An example
```go
type Person struct {
	name string;
	age int;
	gender string
}
```
Structs can be initialized as follows
```go
func main() {
	me := Person{
		name:   "Sankalp Pol",
		age:    42,
		gender: "Male",
	}
}
```
The initialized variable can now be used as follows
```go
fmt.Println("My name is", me.name, "My age is", me.age, "My gender is", me.gender)
```
Struct values can be reinitialized
```go
func main() {
	me := Person{
		name:   "Sankalp Pol",
		age:    42,
		gender: "Male",
	}
	fmt.Println("My name is", me.name, "My age is", me.age, "My gender is", me.gender)
	
	me.name = "Kunal" // initializing a different name
	
	fmt.Println("My name is", me.name, "My age is", me.age, "My gender is", me.gender)
}
```
You can create multiple variables of same struct
```go
func main() {
	me := Person{
		name:   "Sankalp Pol",
		age:    42,
		gender: "Male",
	}

	you := Person{
		name:   "Rahul Jaykar",
		age:    64,
		gender: "Male",
	}
	fmt.Println(me, you)
}
```
Struct variables can be passed to functions as well
```go
type Person struct {
	name   string
	age    int
	gender string
}

func greetPerson(person Person) {
	fmt.Println("My name is", person.name, "My age is", person.age, "My gender is", person.gender)
}

func main() {
	me := Person{
		name:   "Sankalp Pol",
		age:    42,
		gender: "Male",
	}
	greetPerson(me)
	greetPerson(Person{
		name:   "Kaveri",
		age:    77,
		gender: "Female",
	})
}
```
## Maps in Go
A map in go is similar to map in most other languages. Its a key:value pair. The syntax to define a map is as follows
```
var varName = map[KeyType]ValueType{key1: value1, key2: value2...}
OR
varName := map[KeyType]ValueType{key1: value1, key2: value2...}
OR
var varname = make(map[KeyType]ValueType) 
```
For example
```go
var person = map[string]string{"name": "Sankalp Pol", "gender": "Male"}
```
If you use make function to create a map, by default, the map will hold ```nil```

Keys cannot be of following types
1. Slices
2. Maps
3. Functions

The key can hold following types
1. Booleans
2. Numbers
3. Strings
4. Arrays
5. Pointers
6. Structs
7. Interfaces (should not contain slices, maps and functions)

The invalid key types are invalid because you cannot compare two functions or two slices or two maps using equality operator.

However, those types can be compared to ``nil``

Accessing a map value is similar to accessing an array or slice index, we just replace index with key.
```go
var person = map[string]string{"name": "Sankalp", "gender": "Male"}
fmt.Println(person["name"])
```
You can also update or add elements to a map
```go
var person = map[string]string{"name": "Sankalp", "gender": "Male"}
person["name"] = "Kunal" // this updates name
person["city"] = "Mumbai" // this adds a key "city" with value "Mumbai"
```
To delete keys, you can use the delete function
```go
var person = map[string]string{"name": "Sankalp", "gender": "Male"}
delete(person, "name") //name will be deleted
```
To check if a key exists, you can use following way
```go
func main(){
	var person = map[string]string{"name": "Sankalp", "gender": "Male"}
	name, ok := person["name"]
	if ok != true {
		fmt.Println("Name doesn't exist")
	}else{
		fmt.Println("Name exists as", name)
	}
}
```
Iterating over a map
```go
func main(){
	var person = map[string]string{"name": "Sankalp", "gender": "Male"}
	for key, value := range person {
		fmt.Println(key, value)
	}
}
```
## Pointers in Go
Go supports pointers. Pointers are variabes which hold memory address of a value.
```go
var p *int

i := 100
*p = &i

fmt.Println(*p, p)
```
Notice the ``*`` before int. That tells that p is a pointer that points to a memory address for int data type. The ``&`` is used to reference the memory address of a variable.

A pointer variable will have a zero value of ``nil``. Go also does not support pointer arithmetic. You can say that pointers in Go are supported to allow pass by reference in Go. We will discuss more about this later.

## Importing External Packages

We will be importing the quote package that you can find at: ```rsc.io/quote```. To import this package, you can run following command
```
go get rsc.io/quote
```

This should install the package. In case you wish to directly use the package without installing, you can do that. After importing the package without installing, all you need to do is run this command
```
go mod tidy
```
This command finds all packages that aren't installed from your code and then installs it. 

Now, look at the below code.
```go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
```
In the above code we have imported the package which gives us the quote module. In the quote module, there is a Go function. We are running the function inside the Println function. So, the Go method returns a quote string which is then printed on the terminal.

Try and execute this
```
go run main.go
```