package fundamental

import "fmt"

func Variable(){

	// Tipe tipe penulisan variable di golang
	// cara 1
	var full_nama string = "Jang Won-young"

	// cara 2
	var umur int
	umur = 21

	// cara 3
	jenis_kelamin := "perempuan"

	// cara 4
	var my_bini = true

	// cara 5
	var (
		first_name = "Jang"
		last_name = "Won-young"
	)

	fmt.Println("full nama : ", full_nama)
	fmt.Println("first_name : ", first_name)
	fmt.Println("last_name : ", last_name)
	fmt.Println("umur : ", umur)
	fmt.Println("jenis kelamin : ", jenis_kelamin)
	fmt.Println("my bini : ", my_bini)	
}