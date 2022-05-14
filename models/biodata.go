package models

import "fmt"

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (b *Biodata) PrintBio(id string) {
	fmt.Println("Biodata : ")
	fmt.Printf("No\t :\t%s\nNama\t :\t%s\nAlamat\t :\t%s\nPekerjaan:\t%s\nAlasan\t :\t%s\n", id, b.Nama, b.Alamat, b.Pekerjaan, b.Alasan)
}
