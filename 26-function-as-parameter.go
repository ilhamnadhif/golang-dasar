package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello " + filter(name))
}

//func sayHelloWithFilter(name string, filter func(string)string){
//	fmt.Println("Hello "+filter(name))
//}

func toxicFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	filter := toxicFilter
	sayHelloWithFilter("Ilham", filter)

	sayHelloWithFilter("Anjing", toxicFilter)
}
