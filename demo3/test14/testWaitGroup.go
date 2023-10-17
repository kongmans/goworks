package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(j int) {
			time.Sleep(3e9)
			fmt.Println("第", j, "次执行")
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Print("over")
}
