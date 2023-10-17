package main

import "fmt"

func main() {
	a := sum(100)
	fmt.Println(a)
}
func sum(n int) int {
	if n == 1 {
		return 1
	}
	return sum(n-1) + n
}
