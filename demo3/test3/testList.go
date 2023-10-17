package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()
	fmt.Println(myList)

	myList.PushFront("a")
	myList.PushBack("b")
	myList.InsertBefore("c", myList.Back())
	myList.InsertAfter("d", myList.Front())

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, "|")
	}
	fmt.Println()
	fmt.Println(myList.Front().Value, myList.Back().Value)
}
