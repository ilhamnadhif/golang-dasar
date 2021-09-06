package main

import "fmt"

func main() {
	type NoKTP string
	type merried bool

	const ilhamKTP NoKTP = "12345"
	var ilhamMarried merried = false

	fmt.Println(ilhamKTP, ilhamMarried)
}
