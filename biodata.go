package main

import (
	"assignment/models"
	"assignment/repositories"
	"os"
)

func main() {
	args := os.Args[1]
	// fmt.Println("Hello World", reflect.TypeOf(args))
	var bio models.Biodata

	person := repositories.GetUserById(args)

	bio.Nama = person["nama"]
	bio.Alamat = person["alamat"]
	bio.Pekerjaan = person["pekerjaan"]
	bio.Alasan = person["alasan"]

	bio.PrintBio(args)

	// fmt.Println(person)
}
