package fundamental

import "fmt"

type Address struct {
	City, Province, Country string
}

func Pointer() {

	//! 1. Pass by Value (Copy Data)
	fmt.Println("\n===== 1. Pass by Value =====")
	var address1 Address = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}

	address2 := address1       // Copy Value (duplikasi data ke memory baru)
	address2.City = "Bandung" // Hanya merubah address2

	fmt.Println("Address 1 (Original) :", address1)
	fmt.Println("Address 2 (Copied)   :", address2)

	//! 2. Pass by Reference (Pointer &)
	fmt.Println("\n===== 2. Pointer (Pass by Reference) =====")
	var address3 *Address = &address1 // Mengambil Alamat Memory dari address1

	address3.City = "Bekasi" // Mengubah data di memory address1

	fmt.Println("Address 1 (Ikut Berubah) :", address1)
	fmt.Println("Address 3 (Pointer)      :", address3)
}