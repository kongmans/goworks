package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	fmt.Println("默认数据", s1)
	update(s1)
	fmt.Println("引用传递", s1) //基本数据类型，arr,struct都是值传递，切片slice，map,chen是引用传递
}

func update(s2 []int) {
	fmt.Println("接受前", s2)
	s2[0] = 100
	fmt.Println("接受后", s2)
}
