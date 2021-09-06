package main

import "fmt"

func main() {
	var numberA int = 5
	var numberB *int = &numberA

	fmt.Println("numberA", numberA)
	fmt.Println("numberB", numberB)
	fmt.Println("numberB", *numberB)

	*numberB = 6
	fmt.Println("numberA", numberA)
	fmt.Println("numberB", *numberB)

	number := 5
	fmt.Println("Nilai Awal:", number)
	fmt.Println("Alamat Memori:", &number)
	Change(&number, 10)
	fmt.Println("Nilai Awal:", number)
	fmt.Println("Alamat Memori:", &number)
}
func Change(old *int, new int) {
	*old = new
	fmt.Println("Alamat Memori:", old)
}
