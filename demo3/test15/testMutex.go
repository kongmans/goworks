package main

import (
	"fmt"
	"sync"
)

var (
	wg     sync.WaitGroup
	m      sync.Mutex
	number = 100
)

func demo() {
	m.Lock()
	for i := 0; i < 10; i++ {
		number = number - 1
	}
	m.Unlock()
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go demo()
	}
	wg.Wait()
	fmt.Println("done")
}
