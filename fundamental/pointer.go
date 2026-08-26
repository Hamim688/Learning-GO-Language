package fundamental

import "fmt" 

type Address struct {
	City, Province, Country string
}

func Pointer() {
	var address1 Address = Address{
		City: "Jakarta",
		Province: "DKI Jakarta",
		Country: "Indonesia",
	}

	// address2 := address1 // Copy Value 

	// address2.City = "Bandung" // Tidak Merubah address1

	// fmt.Println("Address 1 :",address1) // Tidak Berubah
	// fmt.Println("Address 2 :",address2) // Berubah

	//! 2. Pointer 

	var  address3 *Address = &address1 // Mengambil Alamat Memory dari address1
	
	address3.City = "Bekasi"

	fmt.Println("Address 1 : ", address1) // Berubah
	fmt.Println("Address 3 : ", address3) // Berubah
}