package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}
func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	res := 5 / value
	fmt.Println("result", res)
}
func main() {
	runApplication(0)
}
