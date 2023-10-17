package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	lis, _ := net.ListenTCP("tcp4", addr)

	for {
		conn, _ := lis.Accept()
		go func() {
			b := make([]byte, 1024)
			count, _ := conn.Read(b)
			fmt.Println("服务器接收数据为：", string(b[:count]))
			conn.Write(append([]byte("服务器成功接受："), b[:count]...))
			conn.Close()
		}()

	}

}
