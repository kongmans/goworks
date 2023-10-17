package main

import "fmt"

func t1() {
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

func t2() {
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d*%d=%d \t", i, j, i*j)
		}
		fmt.Println("")
	}
}

func main() {
	t1()
	fmt.Println("-----------------------------------")
	t2()

}
