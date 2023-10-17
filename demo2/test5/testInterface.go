package main

import "fmt"

//"~"表示支持后面的衍生类型
type int8a int8
type MyInt interface {
	int | ~int8 | int16 | int32 | int64 | string | float32 | float64
}

func GetMax[T MyInt](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func main() {
	var a int8a = 10
	var b int8a = 12
	fmt.Println(GetMax[int8a](a, b))

	fmt.Println("--------------------------------")

	n := GetMax[int](100, 2)
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	n1 := GetMax[float64](2.3, 6.6)
	fmt.Println(n1)
	fmt.Printf("%T\n", n1)
}
