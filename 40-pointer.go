package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Pass by Value = meduplikat data
	//address1 := Address{City: "Kudus", Province: "Jawa Tengah", Country: "Indonesia"}
	//address2 := address1
	//
	//address1.City = "Demak"
	//fmt.Println(address1) // hasil {Demak Jawa Tengah Indonesia}
	//fmt.Println(address2) // hasil {Kudus Jawa Tengah Indonesia}

	var address1 Address = Address{City: "Kudus", Province: "Jawa Tengah", Country: "Indonesia"}
	//var address2 = &address1
	var address2 *Address = &address1

	address2.City = "Demak"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("=============================")

	*address2 = Address{City: "Malang", Province: "Jawa Timur", Country: "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("=============================")

	address2 = &Address{City: "Subang", Province: "Jawa Barat", Country: "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("=============================")

	//var address3 *Address = &Address{City: "Semarang"}
	var address3 *Address = new(Address)
	address3.City = "Semarang"
	fmt.Println(address3)
	fmt.Println("=============================")
	var a int = 4
	var b *int = &a
	a = 5
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(*b)

	fmt.Println("=======")
	var address4 = Address{City: "kkk"}
	address4.City = "Kudus"
	fmt.Println(address4)
}
