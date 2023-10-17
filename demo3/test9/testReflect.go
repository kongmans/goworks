package main

import (
	"fmt"
	"reflect"
)

type Peopel struct {
	Name    string
	Address string
}

func main() {
	//a := 1.5
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.ValueOf(a))
	//
	peo := Peopel{"kongman", "Hunan"}
	fmt.Println(peo)
	//v := reflect.ValueOf(peo)
	//fmt.Println(v.NumField())
	//fmt.Println(v.FieldByIndex([]int{0}))

	content := "Name"
	//fmt.Println(v.FieldByName(content))
	//fmt.Println(v.FieldByName(content).CanSet())

	peo1 := new(Peopel)
	v1 := reflect.ValueOf(peo1).Elem()
	fmt.Println(v1.FieldByName(content).CanSet())
	v1.FieldByName(content).SetString("jim")
	v1.FieldByName("Address").SetString("hubei")

	fmt.Println(peo1)
}
