package main

import "fmt"

func main() {
	var a = 10
	fmt.Println(a)
	var b int
	b = 20
	fmt.Println(b)
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
	r := make([]int, 4)
	r[0] = 10
	fmt.Println(r)
	s := new(int) // intialized as 0
	fmt.Println(*s)
	t := new(string) // initialized as ""
	fmt.Println(*t)
	u := new([4]int) // initialized as [0,0,0,0]
	fmt.Println(*u)
}
