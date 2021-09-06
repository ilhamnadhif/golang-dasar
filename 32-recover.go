package main

import "fmt"

func runApp2(error bool) {
	defer endApp2()
	fmt.Println("Aplikasi Berjalan")
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Proses")
}
func endApp2() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi Berhenti")
}
func main() {
	runApp2(true)
}
