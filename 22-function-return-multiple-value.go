package main

import "fmt"

func getBiodata(firstname string, lastname string, age int) (string, string, int) {
	return firstname, lastname, age
}
func main() {
	fmt.Println(getBiodata("ilham", "nadhif", 16))
	var a, b, c = getBiodata("afif", "arifin", 17)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var _, _, umur = getBiodata("oke", "oke", 56)
	fmt.Println(umur)
}
