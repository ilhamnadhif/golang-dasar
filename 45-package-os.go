package main

import (
	"fmt"
	"os"
)

func main() {
	var args []string = os.Args
	fmt.Print("Arguments : ")
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	halo := os.Getenv("GOPATH")
	fmt.Println(halo)
}
