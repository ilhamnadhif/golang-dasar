package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())

	utc := time.Date(2004, time.September, 10, 11, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	parse, _ := time.Parse("2006-01-02", "2004-09-10")
	fmt.Println(parse)
}
