package main

import "fmt"

func terserah (name string) string{
	return "Good Bye " + name
}
func main() {
	goodBye := terserah
	fmt.Println(goodBye("Ilham"))
}
