package main

import (
	"fmt"
)

func main() {
	counter := 0
	for counter < 5 {
		fmt.Println("ini index ke", counter+1)
		counter++
	}
	for i := 0; i < 5; i++ {
		fmt.Println("ini angka ke", i+1)
	}

	var slice = []string{"ilham", "nadhif", "yanuar"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for index, name := range slice{
		fmt.Println("index", index, "=", name)
	}
	alamatSiswa := map[string]string{
		"Ilham": "Kudus",
		"Afif": "Pati",
		"Yanuar": "Kudus",
	}
	for index, nama := range alamatSiswa{
		println(index, nama)
	}
}
