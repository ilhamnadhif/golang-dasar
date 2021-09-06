package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put your database host")
	var user *string = flag.String("user", "root", "put your database user")
	var password *string = flag.String("password", "root", "put your database password")
	var number *int = flag.Int("number", 100, "put your number")
	flag.Parse()
	fmt.Println(*host)
	fmt.Println(*user)
	fmt.Println(*password)
	fmt.Println(*number)
}
