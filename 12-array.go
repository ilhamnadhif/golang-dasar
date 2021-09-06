package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Ilham"
	names[2] = "Nadhif"
	fmt.Println(names)
	fmt.Println(names[0])

	values := [3]int{20, 30, 21}
	fmt.Println(values[2])
	fmt.Println(len(values))

	var lagi [10]int
	fmt.Println(len(lagi))

	arr := [...]int{1, 2, 3, 4, 34, 32}
	fmt.Println(len(arr))

}
