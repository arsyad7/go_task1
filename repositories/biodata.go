package repositories

func GetUserById(id string) map[string]string {
	var people = []map[string]string{
		{
			"id":        "1",
			"nama":      "M Arsyad Ramadhan",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "2",
			"nama":      "Tiara Dewangga",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "3",
			"nama":      "Juni Dio Kasandra",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "4",
			"nama":      "Tasrifin",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "5",
			"nama":      "Adhitya Febhiakbar",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "6",
			"nama":      "Esra Delima Manurung",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "7",
			"nama":      "Muhammad Avtara",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "8",
			"nama":      "Hamonangan Sitorus",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "9",
			"nama":      "Julius Martogi",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "10",
			"nama":      "Indra Bayu Sudirman",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "11",
			"nama":      "Philip Bryan Halim",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "12",
			"nama":      "Teguh Ainul Darajat",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "13",
			"nama":      "Saut Raja Marihot Tua",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "14",
			"nama":      "Putra Irawan",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "15",
			"nama":      "Albert",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
		{
			"id":        "16",
			"nama":      "Dhany Putra",
			"alamat":    "Jakarta",
			"pekerjaan": "Backend Engineer",
			"alasan":    "Mencari pengalaman",
		},
	}
	// fmt.Println(id, "<><><", people[0])
	// idInt := strToInt(id)
	// fmt.Println(reflect.TypeOf(idInt))

	result := map[string]string{}

	for _, v := range people {
		if v["id"] == id {
			result = v
			break
			// fmt.Println(v)
		}
	}
	// fmt.Println(result)
	return result
}

// func strToInt(str string) int {
// 	i, err := strconv.Atoi(str)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(i, "ini sudah jadi number")
// 	return i
// }
