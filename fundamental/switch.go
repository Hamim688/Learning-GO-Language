package fundamental

import "fmt"

func SwitchExpression() {

	//! 1. Switch Expression
	fmt.Println("\n===== 1. Switch Expression =====")
	name := "Jang Won-Young"

	switch name {
	case "Jang Won-Young":
		fmt.Println("It's My Bini ❤️")
	case "Markonah":
		fmt.Println("Halo Markonah")
	case "Budi":
		fmt.Println("Halo Budi")
	default:
		fmt.Println("Ihh siapa sih, Gak Kenal")
	}

	//! 2. Switch dengan Short Statement
	fmt.Println("\n===== 2. Switch dengan Short Statement =====")
	// Ubah nama
	name = "Markonah"

	switch panjang := len(name); panjang > 5 {
	case true:
		fmt.Println("Panjang amat nama lu")
	case false:
		fmt.Println("Nama lu Pendek")
	}

	//! 3. Switch Tanpa Kondisi
	fmt.Println("\n===== 3. Switch Tanpa Kondisi =====")
	// Ubah nama
	name = "Jang Won-Young"

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Panjang amat nama lu")
	case length > 5:
		fmt.Println("Agak Kepanjangan sihh")
	default:
		fmt.Println("Nama lu Pendek")
	}
}
