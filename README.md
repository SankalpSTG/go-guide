# Go Guide

This guide is a quick learning guide for Go Programming Language. Prerequisites: you should have good knowledge of other programming languages.

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