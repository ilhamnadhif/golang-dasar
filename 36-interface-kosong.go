package main

import "fmt"

func Ups(i int) interface{} {
	switch i {
	case 1:
		return 1
	case 2:
		return true
	default:
		return "Ups"
	}
}
func main() {
	var data interface{} = Ups(3)
	fmt.Println(data)
	tes := Ups(1)
	fmt.Println(tes)
}
