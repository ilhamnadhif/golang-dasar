package main

import "fmt"

func runApp(error bool){
	defer endApp()
	fmt.Println("Aplikasi Berjalan")
	if error{
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Proses")
}
func endApp(){
	fmt.Println("Aplikasi Berhenti")
}
func main() {
	runApp(true)
}
