package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

var m = make(map[int]int)
var m2 = sync.Map{}

func get(x int) (y int) {
	return m[x]
}

func set(x int, value int) {
	m[x] = value
}

// func main() {
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			set(i, i+100)
// 			fmt.Printf("key:%v,value:%v\n", i, get(i))
// 			wg.Done()
// 		}(i)
// 	}

// 	wg.Wait()
// }

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+100)
			value, _ := m2.Load(i)
			fmt.Printf("key:%v,value:%v\n", i, value)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
