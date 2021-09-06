package main

import "fmt"

type Man struct {
	Name string
}
func (man *Man) Married(){
	man.Name = "Mr. "+ man.Name
}
func main() {
	var ilham Man = Man{Name: "Ilham"}
	ilham.Married()
	fmt.Println(ilham.Name)
}
