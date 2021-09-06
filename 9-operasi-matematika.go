package main

import "fmt"

func main() {
	var a = 10
	var b = a + 5
	fmt.Println(b)

	// augmented assigment
	var c = 10
	c +=6
	fmt.Println(c)

	// unary operator
	c++
	c++
	c--
	fmt.Println(c)

	d := -c
	fmt.Println(d)

	var e = true
	var f = !e
	fmt.Println(f)

}
