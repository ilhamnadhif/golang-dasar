package main

import "fmt"

func getFullName(firstName string, middleName string, lastName string) (firstname string, middlename string, lastname string) {
	firstname = firstName
	middlename = middleName
	lastname = lastName
	//return firstname, middlename, lastname
	return
}
func main() {
	a, b, c := getFullName("Muhammad", "Ilham", "Nadhif")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
