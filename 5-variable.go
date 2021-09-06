package main

import "fmt"

func main() {
	var name string;
	name = "Ilham";
	fmt.Println(name)
	name = "Nadhif";
	fmt.Println(name)
	// name = 100 tidak bisa

	var nama = "Muhammad Ilham Nadhif"
	var umur = 16
	var kelas int8 = 11
	fmt.Println("Nama saya", nama, ", usia saya", umur, ", kelas", kelas)

	country := "Asia" // kata kunci var bisa diganti
	fmt.Println(country)
	country = "Indonesia"
	fmt.Println(country)

	var (
		firstname = "Ilham"
		lastname = "Nadhif"
	)
	println(firstname, lastname)
}
