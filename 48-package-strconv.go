package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("bool", boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	float, err := strconv.ParseFloat("245.3", 64)
	if err == nil {
		fmt.Println("float", float)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("245", 10, 64)
	if err == nil {
		fmt.Println("int", number)
	} else {
		fmt.Println("Error", err.Error())
	}

	value := strconv.FormatInt(12345, 10)
	fmt.Println("string", value)


	valueString, _ := strconv.Atoi("10000")
	fmt.Println("int", valueString)
}
