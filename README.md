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

## Creating Modules in Go



## Concurrency in Golang


## Channels in Golang