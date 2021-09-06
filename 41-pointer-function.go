package main

import "fmt"

type Address struct {
	City, Province, Country string
}
func changeAddressToIndonesia(address *Address){
	address.Country = "Indonesia"
}
func main() {
	var address Address = Address{City: "Kudus", Province: "Jawa Tengah", Country: "Jepang"}
	changeAddressToIndonesia(&address)
	fmt.Println(address)
}
