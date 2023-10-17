package main

import "fmt"

func main() {
	a := new(int)
	fmt.Printf("%T,%p\n", a, a)
	*a = 100
	fmt.Println(*a)

	fmt.Println("--------------------------")
	var b *int
	fmt.Printf("%T,%p\n", b, b)
	fmt.Println(b)
	//b=2  不能赋值，因为没有地址
	//指向ａ地址
	b = a
	fmt.Println(&b)
	fmt.Println(*b)

}
