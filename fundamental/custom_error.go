package fundamental

import (
	"fmt"
)

//! 1. Custom Struct Error Pertama (ValidationError)
type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

//! 2. Custom Struct Error Kedua (NotFoundError)
type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

//! Function yang Mengembalikan Custom Error Berdasarkan Kondisi
func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"ID is required"}
	}

	if id != "1" {
		return &NotFoundError{"ID not found"}
	}

	return nil
}

func ErrorCustom() {
	//! 1. Pengujian Pengecekan Custom Error Menggunakan Type Assertions (if-else ok)
	fmt.Println("\n===== 1. Custom Error (Validation Error) =====")
	errValidation := SaveData("", nil)
	if errValidation != nil {
		if validationError, ok := errValidation.(*ValidationError); ok {
			fmt.Println("Validation Error :", validationError.Message)
		} else if notFoundError, ok := errValidation.(*NotFoundError); ok {
			fmt.Println("Not Found Error  :", notFoundError.Message)
		} else {
			fmt.Println("Unknown Error    :", errValidation.Error())
		}
	}

	//! 2. Pengujian Pengecekan Custom Error Menggunakan Type Switch
	fmt.Println("\n===== 2. Custom Error (Not Found Error via Type Switch) =====")
	errNotFound := SaveData("2", nil)
	if errNotFound != nil {
		switch e := errNotFound.(type) {
		case *ValidationError:
			fmt.Println("Validation Error :", e.Message)
		case *NotFoundError:
			fmt.Println("Not Found Error  :", e.Message)
		default:
			fmt.Println("Unknown Error    :", e.Error())
		}
	}

	//! 3. Pengujian Skenario Sukses (Tanpa Error)
	fmt.Println("\n===== 3. Save Data Sukses (Tanpa Error) =====")
	errSukses := SaveData("1", "Data Berhasil Disimpan")
	if errSukses == nil {
		fmt.Println("Status           : Data saved successfully")
	}
}