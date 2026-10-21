package standardlibrary

import (
	"errors"
	"fmt"
)

// Sentinel Error: Variabel error yang dibuat di tingkat package (global)
// Berfungsi sebagai "identitas unik" untuk membedakan jenis kesalahan yang terjadi
var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

// Fungsi yang mengembalikan tipe interface 'error'
func GetById(id string) error {
	// Jika id kosong, kembalikan jenis error ValidationError
	if id == "" {
		return ValidationError
	}

	// Jika id tidak cocok, kembalikan jenis error NotFoundError
	if id != "Won-Young" {
		return NotFoundError
	}

	// Kembalikan nil jika tidak ada error (berhasil/sukses)
	return nil
}

func ErrorFunctions() {
	err := GetById("ahyeon")
	if err != nil {
		// errors.Is(err, target): Mengecek apakah error yang terjadi cocok dengan jenis Sentinel Error tertentu
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error:", err.Error())
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error :", err.Error())
		} else {
			fmt.Println("Unknown error   :", err.Error())
		}
	}
}
	

