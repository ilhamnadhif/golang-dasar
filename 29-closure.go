package main

import "fmt"

func main() {
	name := "ilham"
	counter := 0
	increment := func() {
		name := "nadhif"
		counter++
		fmt.Println(name)
	}
	increment()
	fmt.Println(name)
	fmt.Println(counter)
}
