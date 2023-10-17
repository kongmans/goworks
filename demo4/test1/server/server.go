package main

import (
	"fmt"
	"net"
)

func main() {
	//1,创建服务器地址
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")

	//2,创建监听器
	lis, _ := net.ListenTCP("tcp4", addr)
	//3,通过监听器建立连接
	fmt.Println("server is working!")
	conn, _ := lis.Accept()
	//4,装换数据
	b := make([]byte, 1024)
	n, _ := conn.Read(b)
	fmt.Println("获取数据为：", string(b[:n]))
	//5,关闭连接
	conn.Close()
}
