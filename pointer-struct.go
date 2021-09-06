package main

import "fmt"

type Siswa struct {
	Name  string
	Kelas int
}

func (siswa *Siswa) Change() {
	siswa.Name = "Hello " + siswa.Name
}
func Ubah(siswa *Siswa) {
	siswa.Name = "Nama saya " + siswa.Name
}
func main() {
	var siswa1 Siswa = Siswa{Name: "Ilham", Kelas: 12}
	siswa1.Change()
	fmt.Println(siswa1.Name)
	fmt.Println("===========")
	var siswa2 Siswa = Siswa{Name: "Nadhif", Kelas: 12}
	Ubah(&siswa2)
	fmt.Println(siswa2.Name)
}
