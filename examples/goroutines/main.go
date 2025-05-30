package main

import (
	"fmt"
	"time"
)

func say(s string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(s)
}

func square(n int, c chan int) {
	sq := n * n

	c <- sq
}
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

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}
func main() {
	go say("world")
	say("hello")

	c := make(chan int)

	go square(4, c)

	x := <-c
	fmt.Println(x)

	bc := make(chan int, 1)

	bc <- 10

	y := <-bc

	fmt.Println(y)

	fibc := make(chan int)

	go fibonacci(10, fibc)

	for val := range fibc {
		fmt.Println(val)
	}

	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacci2(ch, quit)
}
