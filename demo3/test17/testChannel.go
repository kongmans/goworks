package main

import "fmt"

// 同步
// 主协程和子协程通信
//func main() {
//	ch := make(chan int)
//	go func() {
//		ch <- 100
//	}()
//	a := <-ch
//	fmt.Print(a)

//}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)
	go func() {
		ch1 <- "传送数据给拧一个协程"
		ch2 <- 1
	}()
	go func() {
		content := <-ch1
		fmt.Println(content)
		ch2 <- 2
	}()
	<-ch2
	<-ch2
	fmt.Println("over")
}
