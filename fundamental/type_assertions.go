package fundamental

import "fmt"

func OK() any {
	return "OK"
}

func TypeAssertions() {
	//! 1. Type Assertion Manual
	fmt.Println("\n===== 1. Type Assertion Manual =====")
	data := OK()

	// Mengonversi data dari any/interface{} menjadi string
	var dataString string = data.(string)
	fmt.Println("Data String :", dataString)

	//! 2. Type Assertion Menggunakan Switch Expression (Aman dari Panic)
	fmt.Println("\n===== 2. Type Assertion dengan Switch =====")
	switch value := data.(type) {
	case string:
		fmt.Println("String :", value)
	case int:
		fmt.Println("Integer :", value)
	case bool:
		fmt.Println("Boolean :", value)
	default:
		fmt.Println("Unknown :", value)
	}
}