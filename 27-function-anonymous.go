package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
func main() {
	toxicFilter := func(name string) bool {
		return name == "anjing"
	}
	registerUser("anjing", toxicFilter)
	registerUser("ilham", func(name string) bool {
		return name == "kasar"
	})
}
