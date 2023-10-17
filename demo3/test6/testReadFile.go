package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("e:/godir/hello.txt")
	if err != nil {
		fmt.Println("file open fail", err)
		return
	}
	fmt.Println("file open success!")
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	fmt.Println("读取成功,内容为", string(b))

}
