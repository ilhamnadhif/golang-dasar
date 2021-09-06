package main

import "fmt"

func factorialLoop(angka int) int {
	hasil := 1
	for i := 1; i <= angka; i++ {
		hasil *= i
	}
	return hasil
}
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(factorialRecursive(5))
}
