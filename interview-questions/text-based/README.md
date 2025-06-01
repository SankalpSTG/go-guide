# Golang Text Based Interview Q.estions

These are some hard level questions that could be asked in Go. These are edge questions and you should not be preparing these before covering basics of Golang.

## Goroutines & Channels
### Q. What are the internal scheduling principles of goroutines? How does Go runtime manage them?
### Q. What is the difference between buffered and unbuffered channels? How does it affect goroutine synchronization?
### Q. What happens if you close a channel and then try to send data on it?
### Q. How does select work with multiple channel cases? What if multiple cases are ready?
### Q. What are the guarantees provided by Go for channel send/receive ordering (if any)?
### Q. What happens if a goroutine panics? How do you safely recover from panics inside goroutines?
### Q. How does time.After and time.Ticker work under the hood? Do they leak if not properly stopped?

## Memory Model & Internals
### Q. How does escape analysis work in Go? How can you find whether a variable escapes to the heap?
### Q. Explain the stack vs heap memory allocation in Go. When does the compiler decide to move to heap?
### Q. What is the purpose of the sync.Pool and when should you use it?
### Q. How does Go manage memory alignment and padding in structs?
### Q. What are false sharing and cache line alignment in Go structs?

## Interfaces & Reflection
### Q. How is an interface implemented internally in Go? What does its memory layout look like?
### Q. What's the zero value of an interface? Why can comparing a nil interface to nil return false?
### Q. How does reflection work in Go? What are the types returned by reflect.TypeOf and reflect.ValueOf?

## Concurrency Primitives
### Q. What is a race condition? How can you detect it in Go?
### Q. What’s the difference between sync.Mutex and sync.RWMutex? When should you use each?
### Q. How is context.Context designed to help with goroutine control and cancellation?

## Language Features
### Q. Why does Go not support generics (older versions)? How has the introduction of generics in Go 1.18 changed things?
### Q. What are the restrictions when embedding structs? How does method resolution work in embedded structs?
### Q. How does defer work? When is the expression evaluated and when is it executed?

## Package System & Compilation
### Q. What is the difference between init() and main()? How are init() functions ordered in execution?
### Q. How does Go handle package dependency resolution and import cycles?
### Q. What’s the difference between go run, go build, and go install?

## Language Semantics & Behavior
### Q. What are blank identifiers (_) used for in Go? Name at least 3 use cases.
### Q. Why is Go’s range over a slice sometimes dangerous when using the loop variable inside goroutines?
### Q. What happens when you append to a slice and then mutate the original? Does the new slice share memory?
### Q. What is the difference between new() and make()? When should you use one over the other?
### Q. How does Go handle map iteration order? Is it deterministic? Why or why not?

## Type System & Methods Sets
### Q. What is a method set in Go? How does it determine interface satisfaction?
### Q. What’s the difference between a pointer receiver and a value receiver in Go methods? What are the rules for method sets on each?
### Q. Can you implement an interface using a type alias? Why or why not?
### Q. Why can't Go structs have default values like in constructors? How does the language promote immutability differently?
### Q. Explain the differences between type embedding and inheritance. How does Go simulate inheritance?

## Compilation & Runtime Details
### Q. How does Go achieve cross-compilation? What makes it easier in Go than in languages like C++?
### Q. What does the go:generate directive do? How is it used in real projects?
### Q. What happens when a panic is not recovered? How does it differ from os.Exit()?
### Q. How is garbage collection implemented in Go (e.g., tri-color marking)?
### Q. What is the size of an empty struct (struct{}) and why is it used in maps and channels?

## Modules & Build System
### Q. What’s the difference between $GOPATH and Go Modules? Why was the switch necessary?
### Q. What are replace and exclude directives in go.mod and how are they useful?
### Q. What happens when you run go mod tidy? What about go mod vendor?
### Q. How does Go ensure module versioning works correctly across teams and CI pipelines?
### Q. Explain import path resolution in module mode. How does Go locate remote packages?
### Q. How does go.sum work? What security role does it play in module downloads?

## Tooling / Ecosystem
### Q. How does go test discover and run tests? What are table-driven tests in Go?
### Q. What is go vet and how is it different from golint or staticcheck?

## Performance & Runtime
### Q. How does Go manage goroutine scheduling (M:N model)? What are G, M, and P?
### Q. What are finalizers in Go’s GC? When should you (not) use them?

## Testing & Mocking
### Q. How do you mock interfaces in Go for unit testing? What are tools like gomock or testify?
### Q. How does Go handle benchmarking with testing.B?

## Security / Production
### Q. How do you write race-free concurrent code with channels vs. mutexes? Tradeoffs?
### Q. What are common Go anti-patterns in large codebases (e.g., overusing interface{}, global state)?

# Compiler & Optimization
### Q. What are inlining and escape analysis optimizations? How can you confirm if a function was inlined?
### Q. What is //go:nosplit and when would you use it?

# Generics
### Q. How does Go’s type inference work in generic functions? When does it fail?

# Web/Cloud/Practical Focus
### Q. How do you gracefully shut down a Go HTTP server?
### Q. What is the purpose of http.Server's IdleTimeout, ReadTimeout, WriteTimeout?

# Pragmatic Go (Team Scaling & Prod)
### Q. How do you structure large Go projects for modularity and testing?
### Q. How do you handle shared config/state across packages without global variables?