package main

import "fmt"

func main() {
	v := operator(1, 2, add)
	v1 := operator(1, 2, sub)
	v2 := operator(4, 2, chu)
	fmt.Println(v, v1, v2)

}
func operator(a, b int, f func(int, int) int) int {
	r := f(a, b)
	return r
}
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func chu(a, b int) int {
	if b == 0 {
		fmt.Println("除数不能为０")
		return 0
	}

	return a / b

}
