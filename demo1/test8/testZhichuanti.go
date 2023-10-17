package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println("默认数据", arr)
	update(arr)
	fmt.Println("值传递", arr) //基本数据类型，arr,struct都是值传递，复杂类型是引用传递
}

func update(arr２ [4]int) {
	fmt.Println("接受前", arr２)
	arr２[0] = 100
	fmt.Println("接受后", arr２)

}
