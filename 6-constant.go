package main

import "fmt"

func main() {
	const firstname = "Ilham"
	const age int8 = 16
	fmt.Println(firstname, age)
	const value = 100 // bisa tidak dipakai

	const (
		nama = "Nadhif"
		umur int8 = 9
	)
	fmt.Println(nama, umur)
}
