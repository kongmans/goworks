package main

import (
	"fmt"
	"net"
)

// var content ="hello!"
func main() {
	//创建服务器端地址
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	//创建连接
	conn, _ := net.DialTCP("tcp", nil, addr)
	//发送数据
	conn.Write([]byte("你好!	--from client1"))
	fmt.Println("已发送")
	//关闭连接
	conn.Close()

}
