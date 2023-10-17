package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//log.Println("打印日志信息")
	//log.Panicln("打印恐慌日志信息")
	//log.Fatalln("打印致命日志信息")
	f, _ := os.OpenFile("e:/godir/log.log", os.O_APPEND|os.O_CREATE, 0777)
	logger := log.New(f, "info", log.Ltime)
	logger.Println("打印日志信息\r\n")
	fmt.Println("执行成功")
}
