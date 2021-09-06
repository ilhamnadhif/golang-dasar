package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func SayHeloo(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main() {
	var ilham Person
	ilham.Name = "Ilham"
	SayHeloo(ilham)

	cat := Animal{
		Name: "Kucing",
	}
	SayHeloo(cat)
}

