package main

import "fmt"

func main() {
	type Slice[T int | float32 | float64] []T

	var a Slice[int] = []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("a type %T", a)
	var b Slice[float32] = []float32{1, 2, 3}
	fmt.Println(b)
	fmt.Printf("b type %T", b)
	var c Slice[float64] = []float64{1, 2, 3}
	fmt.Println(c)
	fmt.Printf("c type %T\n", c)

	type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE
	var m1 MyMap[string, float32] = map[string]float32{
		"java":   9.9,
		"go":     10.0,
		"python": 8.0,
	}
	fmt.Println(m1)
}
func a() {

}
