package main

import "fmt"

type Hitung interface {
	GetLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) GetLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) GetLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) GetLuas() int {
	return 1001
}

func main() {
	persegi := Hasil(Persegi{Sisi: 4})
	fmt.Println("Luas Persegi:", persegi)

	persegiPanjang := Hasil(PersegiPanjang{Panjang: 3, Lebar: 5})
	fmt.Println("Luas Persegi Panjang:", persegiPanjang)

	asal := Hasil(Asal{Angka: 3})
	fmt.Println("Asal:", asal)
}
func Hasil(hitung Hitung) int {
	return hitung.GetLuas()
}
