package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type people struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

func main() {
	peo := people{Id: 100, Name: "kongman", Address: "湖南"}
	fmt.Println(peo)
	b, _ := xml.MarshalIndent(peo, "", "	")
	bt := append([]byte(xml.Header), b...)
	ioutil.WriteFile("e:/godir/people.xml", bt, 0660)
	fmt.Println("执行结束！")
}
