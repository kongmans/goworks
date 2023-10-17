package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var rw sync.RWMutex
	wg.Add(10)

	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			rw.Lock()
			m[j] = j
			fmt.Println(m)
			wg.Done()
			rw.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println("done")

}
