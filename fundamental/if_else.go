package fundamental

import "fmt"

func IfElse() {
	//! 1. if Expression
	fmt.Println("\n===== 1. If Expression =====")
	name := "Jang Won-Young"

	if name == "Jang Won-Young" {
		fmt.Println("It's My Bini")
	}

	//! 2. Else Expression
	fmt.Println("\n===== 2. Else Expression =====")
	// Ubah nama
	name = "Markonah"

	if name == "Jang Won-Young" {
		fmt.Println("It's My Bini")
	} else {
		fmt.Println("Ihh siapa sih, Gak Kenal")
	}

	//! 3. Else IF Expression
	fmt.Println("\n===== 3. Else If Expression =====")
	// Ubah nama
	name = "Markonah"

	lenght := len(name)
	if lenght > 10 {
		fmt.Println("Panjang amat nama lu")
	} else if lenght > 5{
		fmt.Println("Agak Kepanjangan sihh")
	} else {
		fmt.Println("Nama lu Pendek")
	}

	//! 4. IF dengan Short Statement
	fmt.Println("\n===== 4. IF dengan Short Statement =====")
	// Ubah nama
	name = "Markonah"

	if panjang := len(name); panjang > 10 {
		fmt.Println("Panjang amat nama lu")
	} else {
		fmt.Println("Nama lu Pendek")
	}
}