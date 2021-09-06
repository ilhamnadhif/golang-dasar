package main

import "fmt"

func main() {
	var nilai = 90

	if nilai > 90 {
		fmt.Println("A")
	} else if nilai > 75 && nilai <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	var name = "Ilham"
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
