package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
	Age int
}

func IsValid(data interface{})bool{
	t := reflect.TypeOf(data)
	for i := 0; i< t.NumField(); i++{
		field := t.Field(i)
		if field.Tag.Get("required") == "true"{
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == ""{
				return false
			}
		}
	}
	return true
}

type ContohLagi struct {
	Name string `required:"true"`
	Description string `required:"true"`
}

func main() {
	var sample Sample = Sample{Name: "Ilham"}
	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0))
	fmt.Println(sampleType.Field(1).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	fmt.Println("---------------------")

	fmt.Println(IsValid(sample))
	sample.Name = ""
	fmt.Println(IsValid(sample))

	fmt.Println("---------------------")

	fmt.Println(IsValid(ContohLagi{Name: "Ilham", Description: ""}))
}
