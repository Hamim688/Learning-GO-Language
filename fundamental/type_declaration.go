package fundamental

import "fmt"

func TypeDeclaration(){
	type KTP string
	type Umur int

	var ktp_saya KTP = "1234567890"
	var umur_saya Umur = 21

	fmt.Println("ktp saya : ", ktp_saya)
	fmt.Println("umur saya : ", umur_saya)
}