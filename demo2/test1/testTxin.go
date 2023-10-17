package main

import "fmt"

func main() {
	strs := []string{"hello", "welcome"}
	ints := []int{1, 2}
	fs := []float64{2.2, 3.3}
	testArray(strs)
	testArray(ints)
	testArray(fs)
}

// any类型，等价于＇ｉｎｔｅｒｆａｃｅ｛｝＇
// comparable类型，'int,uint,float,bool,struct,指针'等
//
//	func testArray[T int | string | float64](arr []T) {
//		for _, v := range arr {
//			fmt.Println(v)
//		}
func testArray[T any](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}

}
