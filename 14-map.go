package main

import "fmt"

func main() {
	var person = map[string]string{
		"Ilham" : "Kudus",
		"Fahmi" : "Demak",
		"Afif" : "Pati",
	}
	person["Rafa"] = "Jakarta"
	fmt.Println(person["Ilham"])
	fmt.Println(person)

	var book = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Eko Kurniawan"
	book["salah"] = "Ups"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "salah")
	fmt.Println(book)
	fmt.Println(len(book))

}
