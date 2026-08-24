package fundamental

import "fmt"

//! 1. Interface Kosong / any sebagai Return Value Serbaguna
func Ups() (any, any, interface{}) {
	// Mengembalikan integer, string, dan boolean sekaligus 
	// (karena tipe data any / interface{} dapat menampung semua tipe data)
	return 1, "Won Young", true
}

func AnyFunction() {
	//! 1. Penggunaan Interface Kosong (any)
	fmt.Println("\n===== 1. Interface Kosong (any) =====")
	data1, data2, data3 := Ups()
	fmt.Println("Data 1 (int)     :", data1)
	fmt.Println("Data 2 (string)  :", data2)
	fmt.Println("Data 3 (boolean) :", data3)
}