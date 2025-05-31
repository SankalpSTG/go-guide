package main

import (
	"math/rand"
	"sync"
)

func main() {
	// m := make(map[int]int)

	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		key := rand.Intn(100)
	// 		m[key] = key * 2 // concurrent write
	// 	}()
	// 	go func() {
	// 		key := rand.Intn(100)
	// 		_ = m[key] // concurrent read
	// 	}()
	// }

	// time.Sleep(time.Second)
	// fmt.Println("Done")
	var mu sync.Mutex
	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			key := rand.Intn(100)
			m[key] = key * 2
			mu.Unlock()
		}()
		go func() {
			mu.Lock()
			key := rand.Intn(100)
			_ = m[key]
			mu.Unlock()
		}()
	}

}
