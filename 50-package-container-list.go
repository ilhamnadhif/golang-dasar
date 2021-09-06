package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Muhammad")
	data.PushBack("Ilham")
	data.PushBack("Nadhif")
	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println("--------------")
	for e := data.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}
	fmt.Println("--------------")
	for e := data.Back(); e != nil; e = e.Prev(){
		fmt.Println(e.Value)
	}
}
