package main

//import (
//	"fmt"
//	"io/ioutil"
//	"os"
//)
//func main() {
//	// Get the file path from an environment variable
//	filePath := os.Getenv("FILE_PATH")
//	if filePath == "" {
//		return
//	}
//	// Validate the file path before opening it
//	if !isValidPath(filePath) {
//		return
//	}
//	f, _ := os.Open(filePath)
//	b, _ := ioutil.ReadAll(f)
//	fmt.Println(string(b))
//}
//// Function to validate the file path
//func isValidPath(filePath string) bool {
//	// Add your logic here
//	return true
//}

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("e:/godir/hi.txt")
	b, _ := ioutil.ReadAll(f)
	fmt.Println(string(b))

	b1, _ := ioutil.ReadFile("e:/godir/hello.txt")
	fmt.Println(string(b1))

	ioutil.WriteFile("e:/godir/hi.txt", []byte("tool2"), 0660)
	fmt.Println("程序结束")

	d, _ := ioutil.ReadDir("e:/")
	for _, dir := range d {
		fmt.Print(dir.Name(), "\t")
	}
}
