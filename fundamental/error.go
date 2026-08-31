package fundamental

import (
	"errors"
	"fmt"
)

//! Helper Function Pembagian dengan Error Handling
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func ErrorFunctions() {
	//! 1. Pengujian Error Handling (Pembagi = 0 -> Mengembalikan Error)
	fmt.Println("\n===== 1. Pembagian dengan Pembagi 0 =====")
	hasil, err := Pembagian(10, 0)

	if err == nil {
		fmt.Println("Hasil Pembagian :", hasil)
	} else {
		fmt.Println("Error           :", err.Error())
	}

	//! 2. Pengujian Pembagian Normal (Pembagi != 0 -> Tanpa Error)
	fmt.Println("\n===== 2. Pembagian Normal (10 / 2) =====")
	hasilSukses, errSukses := Pembagian(10, 2)

	if errSukses == nil {
		fmt.Println("Hasil Pembagian :", hasilSukses)
	} else {
		fmt.Println("Error           :", errSukses.Error())
	}
}