package main

import "fmt"

//	func main() {
//		ch1 := make(chan string, 1)
//		ch2 := make(chan int, 1)
//		ch1 <- "kongman"
//		ch2 <- 100
//
//		select {
//		case a1 := <-ch1:
//			fmt.Println(a1)
//		case a2 := <-ch2:
//			fmt.Println(a2)
//		default:
//			fmt.Println("default")
//
//		}
//	}

// 一直发消息的循环息的循环
func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		default:
		}
	}

}
