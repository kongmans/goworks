package main

import (
	"fmt"
	"time"
)

func demo(count int) {
	for i := 0; i < 100; i++ {
		fmt.Println(count, ":", i)

	}

}
func main() {
	for i := 0; i < 5; i++ {
		go demo(i)
	}
	time.Sleep(3e9)

}
