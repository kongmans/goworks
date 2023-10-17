package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Username)
	fmt.Println(u.Uid)
	fmt.Println(u.Gid)
	fmt.Println(u.HomeDir)

	//error := os.Mkdir("e:/godir", os.ModeDir)
	error := os.MkdirAll("e:/godir", os.ModeDir)
	if error != nil {
		fmt.Println("file create fail", error)
		return
	}
	fmt.Println("file create success!")

}
