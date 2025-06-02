# Golang Text Based Interview Questions

These are some hard level questions that could be asked in Go. These are edge questions and you should not be preparing these before covering basics of Golang.

### Goroutines & Channels
#### Q. What are the internal scheduling principles of goroutines? How does Go runtime manage them?
#### Q. What is the difference between buffered and unbuffered channels? How does it affect goroutine synchronization?
Unbuffered Channels:
```go
ch := make(chan int)
go func(){
  ch <- 1     //blocked until someone receives
}()
val := <- ch  //blocked until a value is sent
```
Channel itself has no storage. Sender is blocked until receiver is ready
Strongly synchronized: send and receive must happen at the same time.
Use case: When you want to enforce hand-off between goroutines

Buffered Channels:
```go
ch := make(chan int, 2) // buffered with size 2
ch <- 1 // doesn't block
ch <- 2 // doesn't block
// ch <- 3 would block — buffer full
```
Channel has a capacity provided while creation. Those many values could be sent before someone receives it. Once the capacity is full, sender is blocked until the receiver receives the values.
Loosely synchronized: sender is not blocked until the buffer is full
Use case: Decouple sender and receiver timing
#### Q. What happens if you close a channel and then try to send data on it?
Go runtime will panic i.e. runtime error with message
```
panic: send on closed channel
```
```go
ch := make(chan int)
close(ch)
ch <- 1 // panic: send on closed channel
```
#### Q. How does select work with multiple channel cases? What if multiple cases are ready?
```go
select {
case msg1 := <-ch1:
    // handle msg1
case msg2 := <-ch2:
    // handle msg2
default:
    // optional: runs if no case is ready
}
```
Select waits on multiple channel operations and pics one that's ready.

If multiple are ready, cases out of ready ones are chosen at random and handled one by one. 

If no cases are ready then the default one is run.

Select is often used in a for loop.
```go
for {
    select {
    case x := <-ch1:
        fmt.Println("Got", x)
    case y := <-ch2:
        fmt.Println("Got", y)
    }
}
```
#### Q. What are the guarantees provided by Go for channel send/receive ordering (if any)?
When it comes to one single channel, messages are ordered the way they are sent and Go gurantees that. However in case of multiple channels, there is no cross-channel ordering gurantee.
```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
fmt.Println(<-ch) // prints 1
fmt.Println(<-ch) // prints 2
```
#### Q. What happens if a goroutine panics? How do you safely recover from panics inside goroutines?
#### Q. How does time.After and time.Ticker work under the hood? Do they leak if not properly stopped?
If a goroutine panics, it will:
1. Stop execution of that goroutine.
2. Unwind the stack within the goroutine.
3. If not recovered, the panic will:
    - Print a stack trace.
    - Crash the whole program.

How to recover safely?

We can use defer + recover inside the goroutine itself.
```go
go func() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    // code that might panic
    panic("something went wrong")
}()
```
PS. recover works only inside a deferred function

A bad example:
```go
// ❌ Won't work
defer recover() // doesn't catch panics from other goroutines
go func() {
    panic("won't be recovered")
}()
```
The above won't work as recover is outside the goroutine.

Good example
```go
go func() {
    defer safeGuard()
    riskyFunction()
}()

func safeGuard() {
    if r := recover(); r != nil {
        log.Println("panic recovered:", r)
    }
}
```

### Memory Model & Internals
#### Q. How does escape analysis work in Go? How can you find whether a variable escapes to the heap?
#### Q. Explain the stack vs heap memory allocation in Go. When does the compiler decide to move to heap?
#### Q. What is the purpose of the sync.Pool and when should you use it?
#### Q. How does Go manage memory alignment and padding in structs?
#### Q. What are false sharing and cache line alignment in Go structs?

### Interfaces & Reflection
#### Q. How is an interface implemented internally in Go? What does its memory layout look like?
#### Q. What's the zero value of an interface? Why can comparing a nil interface to nil return false?
The zero value for an interface is `nil`.
```go
type MyInterface interface {
	Hello() string
}

type MyStruct struct {
}

func (e *MyStruct) Hello() string {
	return "World!"
}

func getStruct() MyInterface {
	var mst *MyStruct = nil
	return mst
}

func main() {
	var x MyInterface = &MyStruct{}
	var y MyInterface = nil
	z := getStruct()
	fmt.Println(x == nil) // false
	fmt.Println(y == nil) // true
	fmt.Println(z == nil) // false
}
```
In this program, we created an interface `MyInterface`. This interface has been implicitly implemented by MyStruct. Let's understand the difference between assignment for x, y, z.

Any interface can be represented as a tuple with type and value
```
(type, value)
```
In the assignment of x, we are assigning both type and value as not nil
```
(*MyStruct, non-nil)
```
In the assignment of y, we are assigning the interface as nil, so it does not know which type is it being assigned to as well.
```
(nil, nil)
```
In the third case for z, we call a function which creates a variable of type *MyStruct and assigns it a value of nil. So here, the type is *MyStruct and value is nil. When we return it, again z captures type as *MyStruct and value as nil.
```
(*MyStruct, nil)
```
`interface == nil` this operation will return true when both value and type is nil. In all other cases, it will return false.

Below is another example:
```go
func main() {
	var x MyInterface = nil
	fmt.Println(x == nil)	 	// true
	var y *MyStruct = nil
	x = y
	fmt.Println(x == nil)	 	// false
}
```
#### Q. How does reflection work in Go? What are the types returned by reflect.TypeOf and reflect.ValueOf?
```go
package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	Hello() string
}

type MyStruct struct {
}

func (e *MyStruct) Hello() string {
	return "World!"
}

func main() {
	var x MyInterface = nil
	fmt.Println(reflect.TypeOf(x))  // true
	fmt.Println(reflect.ValueOf(x)) // true
	var y *MyStruct = nil
	x = y
	fmt.Println(reflect.TypeOf(x))  // false
	fmt.Println(reflect.ValueOf(x)) // false
}
```
`reflect` can be used to identify type or value of any variable. `reflect.TypeOf` returns a value with type `reflect.Type` which has additional methods that you can use for your case. Similarly for values, you can use `reflect.ValueOf` which will return a value of type `reflect.Value`. You can run a switch case for the type:
```go
import (
	"fmt"
	"reflect"
)

func checkType(x interface{}) {
	switch reflect.TypeOf(x).Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Slice:
		fmt.Println("slice")
	case reflect.Struct:
		fmt.Println("struct")
	default:
		fmt.Println("unknown")
	}
}
```
### Concurrency Primitives
#### Q. What is a race condition? How can you detect it in Go?
A race condition is where two or more than two goroutines are trying to access the same resource and at least one of them writes to it without a proper synchronization. This can lead to unpredictable behaviors.
```go
package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	counter++
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	time.Sleep(time.Second)
	fmt.Println(counter) // Non-deterministic result
}
```
Look at the above program. We have created 1000 goroutines through the for loop and all the goroutine share the same counter data and hence when each of them is trying to update it, the others might not be in sync with the updation. 

This does not happen sequentially. There might be multiple goroutines that are incrementing the same value let's say 10 goroutines read the value as 800, when all of them update the value, the final value is't 810 but 801. This is just an example

Hence, the end result will not always be 1000 but could be anything, 1000 or less than 1000.

To detect if race conditions are happening, while running your code you can add an extra `-race` flag to your command:
```
go run -race main.go
```
This will show you how many race conditions occured.

P.S. you can use runtime library and set maximum processes as 1.
```go
package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU())
}
```
This will make sure only one OS Thread is used to run all your goroutines. This will not prevent race conditions as such but updates will be better synchronized as no goroutine is working parallely and the output should be 1000.

Note. reducing threads will slow down your application and is not a way to avoid race conditions.
#### Q. What’s the difference between sync.Mutex and sync.RWMutex? When should you use each?
sync.Mutex:
1. allows only one goroutine to access the critical section at a time.
2. It has no distinction between readers and writers — all access is exclusive.
3. Methods
    1. Lock()
    2. Unlock()

When to use:
1. Your resource is either being read or written by only one goroutine at a time.
2. There’s no performance benefit to separating reads from writes (i.e., reads are rare or very fast).

Example:
```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type MyStruct struct {
	mu   sync.Mutex
	data int
}

func (m *MyStruct) GetData(ch chan int, wg *sync.WaitGroup) {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer wg.Done()
	ch <- m.data
}

func (m *MyStruct) IncrementData() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data += 1
}

func main() {
	ms := MyStruct{}
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < 1000; i++ {
		t := rand.Int()
		if t%2 == 0 { // probably equal calls to both write and read operations
			go ms.IncrementData()
		} else {
			wg.Add(1)
			go ms.GetData(ch, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}
}

```
RWMutex:
1. Allows multiple concurrent readers OR one writer, but not both.
2. Has two sets of methods:
    1. RLock() / RUnlock() for readers
    2. Lock() / Unlock() for writers

Use when:
1. You have many reads and few writes.
2. You want to optimize performance by allowing multiple goroutines to read at the same time.

Example:
```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type MyStruct struct {
	mu   sync.RWMutex
	data int
}

func (m *MyStruct) GetData(ch chan int, wg *sync.WaitGroup) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	defer wg.Done()
	ch <- m.data
}

func (m *MyStruct) IncrementData(wg *sync.WaitGroup) {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer wg.Done()
	m.data += 1
}

func main() {
	ms := MyStruct{}
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < 1000; i++ {
		t := rand.Int()
		if t%10 == 0 { // probably higher calls to read
			wg.Add(1)
			go ms.IncrementData(&wg)
		} else {
			wg.Add(1)
			go ms.GetData(ch, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}
	fmt.Println(ms.data) // will have higher value than data received through channel.
}
```
#### Q. How is context.Context designed to help with goroutine control and cancellation?
There are diffent types of context that help in different ways
1. context.Background()
    1. Always non-nil
    2. Does not cancel or times out.
    3. Better to be used at top / root level.
2. context.TODO()
    1. Use when you are unsure which context to use
    2. Same as context.Background()
3. context.WithCancel(parent)
    1. Manually cancellable context
    2. If passed to another context, upon cancelling the parent context will cancel the child context as well.
```go
context, cancel := context.WithCancel(context.Background())
defer cancel()
```
4. context.WithTimeout(parent, timeout)
    1. Gets automatically cancelled within given timeout.
    2. Can be manually cancelled as well.
    3. Good for enforcing deadlines in http calls or db operations.
```go
context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
defer cancel()
```
5. context.WithDeadline(parent, deadline)
    1. Gets automatically cancelled once the deadline is passed.
    2. Can be manually cancelled as well.
    3. Precise control over timing, aligning with external event.
```go
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
defer cancel()
```
6. context.WithValue(parent, key, value)
    1. Allows additional data to be passed with context.
```go
ctx := context.WithValue(context.Background(), "user_id", "newirtvnewybir733ivf4g")
```
If a function is triggering 1000 goroutines, you can pass a context to the goroutines and the goroutines can decide to work until the context is not cancelled.

Cancellation of a context does not automatically trigger all goroutines to stop but it has to be mnually checked if the context is still alive.

A context can be used to create another context
```go
package main

import (
	"context"
	"fmt"
	"time"
)

type KeyType string

type ValueType struct {
	data int
}

func (m *ValueType) Print() {
	fmt.Println(m.data)
}
func main() {
	ctx := context.Background()
	tctx, cancel := context.WithTimeout(ctx, time.Minute)
	var key KeyType = "user_id"
	var Value ValueType = ValueType{
		data: 1234,
	}
	vctx := context.WithValue(tctx, key, Value)
	data := vctx.Value(key)

	val, ok := data.(ValueType)
	if !ok {
		fmt.Println("Invalid Value Type")
	} else {
		val.Print()
	}

	cancel() // we cancel timeout context here

	if err := vctx.Err(); err != nil {
		// value context was created using timeout context and hence
		// after timeout context gets cancelled, value context also gets cancelled.
		// But we have to check this manually or else you will be continuing code execution.
		fmt.Println(err)
	}
}
```

Context drilling creates a tree structure where a node gets cancelled then the entire subtree gets cancelled.
### Language Features
#### Q. Why does Go not support generics (older versions)? How has the introduction of generics in Go 1.18 changed things?
#### Q. What are the restrictions when embedding structs? How does method resolution work in embedded structs?
Restrictions:

1. Only one field with same name. 

    If two parents have same names, you cannot implicitly use them and it will cause selector ambiguity error.

```go
package main

import "fmt"

type A struct{ Name string }
type B struct{ Name string }
type C struct {
	A
	B
}

func main() {
	x := &C{
		A: A{
			Name: "Hello",
		},
		B: B{
			Name: "World",
		},
	}
	// fmt.Println(x.Name)   // compile time error
	fmt.Println(x.A.Name) // Hello
	fmt.Println(x.B.Name) // World
}
```
2. You can embed only exported types from other packages (starting with upper case)
3. Field promotion is recursively shallow
    
    If B includes A and C includes B, A's inclusion will not be present in C. This is shallow field promotion.
    
    In Go's context, exposed attributes of A will still be accessible directly since Golang will recursively resolve the fileds by first looking into C, then B and then into A. This is recursively shallow field promotion.
```go
package main

import "fmt"

type A struct{ FirstName string }
type B struct {
	A
	MiddleName string
}
type C struct {
	B
	LastName string
}

func main() {
	x := &C{
		B: B{
			A: A{
				FirstName: "Hello",
			},
			MiddleName: "Ello",
		},
		LastName: "Llo",
	}
	fmt.Println(x.FirstName)
	fmt.Println(x.MiddleName)
	fmt.Println(x.LastName)
}
```
In the above example, `x.FirstName` is possible because Go looks for `FirstName` in `C`, then in `B`, then in `A` where it is found.

Method Resolution Rules:

Method resolution in embedded structs works by promotion and delegation.

1. Promoted methods are accessible as if they belong to the outer struct.
```go
type A struct{}
func (A) Hello() { fmt.Println("Hello from A") }

type B struct {
    A
}

b := B{}
b.Hello() // Calls A.Hello — promoted method
```
2. If a method has a pointer receiver, the outer struct must also be a pointer for that method to be accessible. However Go for your convenience interprets value and pointer receivers as per needed.
```go
type A struct{}

func (a *A) Speak() { fmt.Println("Pointer method") }

type B struct {
	A
}

func main() {
	b := B{}
	b.Speak() // here b is converted to (&b).Speak() by Go and no error is thrown
}
```
3. If multiple embedded structs define the same method, the outer struct gets an ambiguity error unless the method is explicitly accessed.
```go
type A struct{}
func (A) Talk() {}

type B struct{}
func (B) Talk() {}

type C struct {
    A
    B
}

c := C{}
c.Talk() // Error: Talk is ambiguous
c.A.Talk() // OK
```
#### Q. How does defer work? When is the expression evaluated and when is it executed?
In Go, defer schedules a function call to be executed after the surrounding function returns.
1. When the defer statement is encountered, the arguments of the deferred function call are evaluated immediately.
2. The actual deferred function call is executed later, just before the function returns, in last-in-first-out (LIFO) order if there are multiple deferred calls.
```go
func example() {
    x := 10
    defer fmt.Println(x)  // 'x' is evaluated now (10), but printed later
    x = 20
    fmt.Println("In function:", x) // prints 20
}
```
Output will be
```
In function: 20
10
```
### Package System & Compilation
#### Q. What is the difference between init() and main()? How are init() functions ordered in execution?
`init()` function is used to initialize package level initialization before the program starts running and `init()` function is automatically called before `main()` function.

We avoid writing complex setup in `init()`. Maybe you can set up some package level variables and things that will surely not cause error / panic.

Everything else that is complex like Database connections and route config should go into `main()` function itself as these are things that can cause errors and are needed to be handled.

You can write multiple `init()` functions in one package while you can only write one `main()` function.
```go
package main

import "fmt"

func init() {
    fmt.Println("Init 1")
}

func init() {
    fmt.Println("Init 2")
}

func main() {
    fmt.Println("Main function")
}
```
Output:
```
Init 1
Init 2
Main function
```
#### Q. How does Go handle package dependency resolution and import cycles?
#### Q. What’s the difference between go run, go build, and go install?

### Language Semantics & Behavior
#### Q. What are blank identifiers (_) used for in Go? Name at least 3 use cases.
#### Q. Why is Go’s range over a slice sometimes dangerous when using the loop variable inside goroutines?
#### Q. What happens when you append to a slice and then mutate the original? Does the new slice share memory?
#### Q. What is the difference between new() and make()? When should you use one over the other?
#### Q. How does Go handle map iteration order? Is it deterministic? Why or why not?

### Type System & Methods Sets
#### Q. What is a method set in Go? How does it determine interface satisfaction?
#### Q. What’s the difference between a pointer receiver and a value receiver in Go methods? What are the rules for method sets on each?
#### Q. Can you implement an interface using a type alias? Why or why not?
#### Q. Why can't Go structs have default values like in constructors? How does the language promote immutability differently?
#### Q. Explain the differences between type embedding and inheritance. How does Go simulate inheritance?

### Compilation & Runtime Details
#### Q. How does Go achieve cross-compilation? What makes it easier in Go than in languages like C++?
#### Q. What does the go:generate directive do? How is it used in real projects?
#### Q. What happens when a panic is not recovered? How does it differ from os.Exit()?
#### Q. How is garbage collection implemented in Go (e.g., tri-color marking)?
#### Q. What is the size of an empty struct (struct{}) and why is it used in maps and channels?

### Modules & Build System
#### Q. What’s the difference between $GOPATH and Go Modules? Why was the switch necessary?
#### Q. What are replace and exclude directives in go.mod and how are they useful?
#### Q. What happens when you run go mod tidy? What about go mod vendor?
#### Q. How does Go ensure module versioning works correctly across teams and CI pipelines?
#### Q. Explain import path resolution in module mode. How does Go locate remote packages?
#### Q. How does go.sum work? What security role does it play in module downloads?

### Tooling / Ecosystem
#### Q. How does go test discover and run tests? What are table-driven tests in Go?
#### Q. What is go vet and how is it different from golint or staticcheck?

### Performance & Runtime
#### Q. How does Go manage goroutine scheduling (M:N model)? What are G, M, and P?
#### Q. What are finalizers in Go’s GC? When should you (not) use them?

### Testing & Mocking
#### Q. How do you mock interfaces in Go for unit testing? What are tools like gomock or testify?
#### Q. How does Go handle benchmarking with testing.B?

### Security / Production
#### Q. How do you write race-free concurrent code with channels vs. mutexes? Tradeoffs?
#### Q. What are common Go anti-patterns in large codebases (e.g., overusing interface{}, global state)?

### Compiler & Optimization
#### Q. What are inlining and escape analysis optimizations? How can you confirm if a function was inlined?
#### Q. What is //go:nosplit and when would you use it?

### Generics
#### Q. How does Go’s type inference work in generic functions? When does it fail?

### Web/Cloud/Practical Focus
#### Q. How do you gracefully shut down a Go HTTP server?
#### Q. What is the purpose of http.Server's IdleTimeout, ReadTimeout, WriteTimeout?

### Pragmatic Go (Team Scaling & Prod)
#### Q. How do you structure large Go projects for modularity and testing?
```
/project-root
│
├── /cmd/                # Entry points (main packages)
│   └── app/             # e.g., main.go
│
├── /internal/           # Private app logic
│   ├── /user/           # A module/domain
│   │   ├── handler.go   # HTTP handlers
│   │   ├── service.go   # Business logic
│   │   ├── repository.go # Interface + impl for data access
│   │   ├── model.go     # Domain models
│   │   └── service_test.go
│   └── /auth/
│
├── /pkg/                # Public utility packages
│   └── /logger/
│
├── /api/                # OpenAPI/proto definitions
│
├── /config/             # App config (YAML/ENV loaders)
│
└── go.mod
```
1. The application is broken into features / modules where each module has it's own folder. 
2. In that folder, you will have services and route handlers.
3. We will use interfaces for abstractions
4. dependencies are decoupled using dependency injection
5. Tight coupling between framework and services is avoided. For e.g. we won't be writing gin route handlers in services.
#### Q. How do you handle shared config/state across packages without global variables?
In Go, to share config or state across packages without using global variables, you typically use dependency injection via structs and interfaces.

```go
//config.go
package config

type AppConfig struct {
    DBURL      string
    RedisHost  string
    SecretKey  string
}
```
The config can be used in other packages while the other packages receive it in their Init methods
```go
// service/user.go
package service

import "myapp/config"

type UserService struct {
    cfg *config.AppConfig
}

func InitUserService(cfg *config.AppConfig) *UserService {
    return &UserService{cfg: cfg}
}

func (u *UserService) DoSomething() {
    // access config
    println(u.cfg.DBURL)
}
```
You will be passing the config to the service in the root level. The struct can be instantiated and passed on to other packages as follows
```go
//main.go
package main

import (
    "myapp/config"
    "myapp/service"
)

func main() {
    cfg := &config.AppConfig{
        DBURL:     "postgres://...",
        RedisHost: "localhost:6379",
        SecretKey: "supersecret",
    }

    svc := service.NewUserService(cfg)
    svc.DoSomething()
}
```