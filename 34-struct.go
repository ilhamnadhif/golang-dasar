package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHai(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var afif Customer
	afif.Name = "Afif"
	afif.Address = "Pati"
	afif.Age = 17
	fmt.Println(afif)

	ilham := Customer{Name: "Ilham", Address: "Kudus"}
	fmt.Println(ilham)
	fmt.Println(ilham.Name)

	ilham.sayHai("Yanuar")
}
