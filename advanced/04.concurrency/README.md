# Goroutines

Goroutines are light weight threads

```go
func say(s string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(s)
}

func main() {
	go say("world")
	say("hello")
}

```

## Channels

```go
func square(n int, c chan int) {
	sq := n * n

	c <- sq
}
func main() {
	c := make(chan int)

	go square(4, c)

	x := <-c
	fmt.Println(x)
}
```

## Buffered Channels

```go
func main() {
	bc := make(chan int, 1)

	bc <- 10

	y := <-bc

	fmt.Println(y)

}
```

## Closing a channel

```go
func fibonacci(n int, c chan int) {
	c <- 0
	if n == 1 {
		return
	}
	c <- 1
	if n == 2 {
		return
	}
	prev, next := 0, 1
	for i := 3; i <= n; i++ {
		temp := next
		next = next + prev
		prev = temp
		c <- next
	}
	close(c)
}

func main() {
	fibc := make(chan int)

	go fibonacci(10, fibc)

	for val := range fibc {
		fmt.Println(val)
	}
}
```