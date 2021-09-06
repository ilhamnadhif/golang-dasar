package main

import "fmt"

func sumAll(name string, numbers ...int) (user string, angka int){
	user = name
	angka = 0
	for _, value := range numbers{
		angka += value
	}
	return
}
func main() {
	nama, total := sumAll("ilham", 1,2,3,4,5)
	fmt.Println("Nama =", nama, "total =", total)

	slice := []int{1,2,3,4}
	_, b := sumAll("Nadhif", slice...)
	fmt.Println(b)
}
