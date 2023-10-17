package main

import "fmt"

type Slice[T int | float64] []T

// 给泛型添加方法
func (s Slice[T]) sum() T {
	var res T
	for _, v := range s {
		res += v

	}
	return res
}

func main() {
	var myslice Slice[int] = []int{1, 2, 3, 4}
	var myslice1 Slice[float64] = []float64{1.2, 2.3, 3.3, 4.5}

	fmt.Println(myslice.sum())
	fmt.Println(myslice1.sum())
}
