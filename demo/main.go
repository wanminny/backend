package main

import (
	"fmt"
)

type LinkNode struct {
	data string
	next *LinkNode
}

type LinkList *LinkNode

//https://stackoverflow.com/questions/17902127/in-go-can-both-a-type-and-a-pointer-to-a-type-implement-an-interface

// The receiver type must be of the form T or *T where T is a type name. The type denoted by T is called the receiver base type;
// it must not be a pointer or interface type and it must be declared in the same package as the method.
//接收器类型必须是T或* T形式，其中T是类型名称。 由T表示的类型称为接收器基类型; 它不能是指针或接口类型，必须在与方法相同的包中声明。
// 针对此次 即是 它不能是指针；因为LinkList是指针；
// 必须在与方法相同的包中声明
func (l LinkList) PrintList() {
	var temp = l
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
	println("done!")
}

func main() {

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello image!\n"))
	//})
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
