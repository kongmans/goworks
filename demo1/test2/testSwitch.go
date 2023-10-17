package main

import "fmt"

func main() {
	var score int = 80
	switch score {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
		//fallthrough
	case 70:
		fmt.Println("C")
	case 60:
		fmt.Println("D")

	}
}
