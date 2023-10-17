package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("go")
	time.AfterFunc(3e9, func() {
		fmt.Println("delay....")
	})
	time.Sleep(4e9)
	fmt.Println("over")
}
