package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		if i == 20 {
			break
		}
		if i % 2 == 0 {
			continue;
		}
		fmt.Println(i)
	}
}
