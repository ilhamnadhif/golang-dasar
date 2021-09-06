package helper

import "fmt"

// jika nama variable atau function diawali huruf besar, maka bisa diakses dari luar
// jika nama variable atau function diawali huruf kecil, maka tidak bisa diakses dari luar

var Application string = "Belajar Golang" // bisa diakses dari luar
var version int = 1 // nggak bisa diakses dari luar

func SayHello(name string){
	fmt.Println("Hello", name)
}

func sayGoodBye(name string){
	fmt.Println("GoodBye", name)
}
