package main

import "fmt"

func main() {
	var p *int
	i := 100
	p = &i
	fmt.Println(p, *p)
	*p = 420
	fmt.Println(p, *p)
}
