package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("hello,中国")
	b := make([]byte, r.Size())
	l, err := r.Read(b)
	if err != nil {
		fmt.Print("读取失败", err)
	}
	fmt.Println("读取成功,strings长度为", l)
	fmt.Println("读取成功,strings内容为", string(b))

}
