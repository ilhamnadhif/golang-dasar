package main

import "fmt"

func main() {
	name := "IlhamNa"
	switch name {
	case "ilham":
		fmt.Println("Halo Ilham")
	case "afif":
		fmt.Println("Halo Afif")
	default:
		fmt.Println("Nama tidak diketahui")
	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	var length = len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}