package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "e:/godir/hi.txt"
	f, err := os.OpenFile(filePath, os.O_APPEND, 0660)
	defer f.Close()
	if err != nil {
		_, err := os.Create("e:/godir/hi.txt")
		fmt.Println("文件不存在，已经新建空白文件", err)
		return
	}

	//f.Write([]byte("输入内容"))
	f.WriteString("hello!!!\r\n")
	fmt.Println("写入成功，程序结束！")

}
