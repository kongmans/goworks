package main

import "fmt"

var slice []string
var arr [10]string
var m map[string]float64
var m1 map[string]string = make(map[string]string)

func main() {

	fmt.Printf("%T,%p\n", slice, slice)
	fmt.Printf("%T,%p\n", arr, arr)
	fmt.Printf("%T,%p\n", m, m)
	fmt.Printf("%T,%p\n", m1, m1)
	fmt.Println("--------------------------")
	names := []string{"jom", "jack"}
	names[0] = "rose"
	fmt.Println(names)
	fmt.Printf("%T,%p\n", names, names)

	names1 := [3]string{"jom1", "jack1"}
	fmt.Println(names1)
	fmt.Printf("%T,%p\n", names1, names1)

	fmt.Println("--------------------------")

	user := map[string]string{"name": "jim", "address": "hunan"}
	fmt.Println(user)
	user["phone"] = "1222222222"
	fmt.Println(user)
	user["phone"] = "1333333333"
	fmt.Println(user)
	delete(user, "phone")
	fmt.Println(user)
	fmt.Println(user["name"])
	value, ok := user["name"]
	fmt.Println(value, ok)

	for key, value := range user {
		fmt.Println(key, value)

	}
}
