package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("----------------------")
	for j := 0; j < 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println(j)
	}
}
