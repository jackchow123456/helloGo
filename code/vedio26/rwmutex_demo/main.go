package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	// lock.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwLock.RUnlock()
	wg.Done()
}

func write() {
	// lock.Lock()
	rwLock.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	// lock.Unlock()
	rwLock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Printf("程序耗时：%s秒\n", time.Now().Sub(start))
}
