package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Halo bro"
	} else {
		return "Hello " + name
	}
}
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}
func main() {
	nama := getHello("Ilham")
	fmt.Println(nama)
	fmt.Println(getHello(""))

	fmt.Println(luasPersegiPanjang(4,5))
}
