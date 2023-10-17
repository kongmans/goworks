package main

import "fmt"

type Slice[T int | float64] []T

// 泛型方法
func Add[T int | float64 | string](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add[int](1, 2))
	fmt.Println(Add[string]("hello", "250"))

}
