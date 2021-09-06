package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice = months[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	months[5] = "Juni diubah"
	slice[0] = "Mei diubah"
	fmt.Println(months)
	fmt.Println(slice)

	fmt.Println("-------------------")

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Ilham")
	fmt.Println(slice3)
	slice3[1] = "Oke"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	fmt.Println("-----membuat slice dari awal / direkomendasikan-----")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ilham"
	newSlice[1] = "Nadhif"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("-----copy slice-----")
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
