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