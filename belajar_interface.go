package main

import (
	"fmt"
)

type Hitung interface {
	GetLuas() int
	GetKeliling() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) GetLuas() int {
	return persegi.Sisi * persegi.Sisi
}

func (persegi Persegi) GetKeliling() int {
	return persegi.Sisi * 4
}
func (persegi Persegi) GetVolume() int {
	return persegi.Sisi * 4
}

func Hasil(hitung Hitung) {
	fmt.Println("Luas =", hitung.GetLuas())
	fmt.Println("Keliling =", hitung.GetKeliling())
}

func main() {
	//var persegi Persegi = Persegi{Sisi: 8}
	//Hasil(persegi)
	var persegi Hitung = Persegi{Sisi: 8} // tipe data interface
	fmt.Println(persegi.GetLuas()) // GetVolume tidak bisa
}
