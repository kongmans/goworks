package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")

	for i := 0; i < 5; i++ {
		conn, _ := net.DialTCP("tcp", nil, addr)
		conn.Write([]byte("hello!	--from client2" + strconv.Itoa(i)))
		b := make([]byte, 1024)
		count, _ := conn.Read(b)
		fmt.Println(string(b[:count]))
		conn.Close()
	}
	fmt.Println("已发送")

}
