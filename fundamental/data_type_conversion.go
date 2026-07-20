package fundamental

import (
	"fmt"
	"strconv"
)

func DataTypeConversion(){
	var nilai32 uint32 = 3232423243
	var nilai64 int64 = int64(nilai32)

	fmt.Println("Nilai 32 : ", nilai32)
	fmt.Println("Nilai 64 : ", nilai64)

	var nilaiString string = "121212"
	
	// Di Golang, konversi String ke Integer HARUS menggunakan package "strconv"
	nilaiInt, err := strconv.Atoi(nilaiString)
	
	if err == nil {
		fmt.Println("Nilai String : ", nilaiString)
		fmt.Println("Nilai Int : ", nilaiInt)
	} else {
		fmt.Println("Gagal konversi:", err)
	}
	
	// Konversi Boolean ke String
	var nilaiBool bool = true
	var boolToString string = strconv.FormatBool(nilaiBool) // Mengubah bool true menjadi string "true"
	
	fmt.Println("Nilai awal (Boolean) :", nilaiBool)
	fmt.Println("Hasil konversi (String) :", boolToString)

	// Konversi String ke Boolean
	var kataBenar string = "true"
	var stringToBool, errBool = strconv.ParseBool(kataBenar) // Mengubah string "true" menjadi bool true
	
	if errBool == nil {
		fmt.Println("\nNilai awal (String) :", kataBenar)
		fmt.Println("Hasil konversi (Boolean) :", stringToBool)
	}
}