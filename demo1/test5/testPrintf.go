package main

import (
	"fmt"
)

func main() {
	var a int = 1
	str1 := "hello,welcome!"
	//fmt.Printf(a)   Printf不能输出ｉｎｔ
	fmt.Println(a)
	fmt.Println(str1)
	fmt.Println(len(str1))
	fmt.Printf("%c \n", str1[0])
	fmt.Println(str1[0])

	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i])
	}
	fmt.Println()
	for i, v := range str1 {
		fmt.Print(i)
		fmt.Printf("%c \t", v)
	}

}
