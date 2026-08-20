package fundamental

import "fmt"

func ForLoops() {

	//! 1. For Loops Sederhana
	fmt.Println("\n===== 1. For Loops Sederhana =====")
	counter := 1

	for counter <= 5 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	//! 2. For dengan Init dan Post Statement
	fmt.Println("\n===== 2. For dengan Init & Post Statement =====")
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke", i)
	}

	//! 3. For Range (Iterasi Collection)
	fmt.Println("\n===== 3. For Range =====")
	names := []string{"Jang Won-Young", "Markonah", "Budi"}

	// Mengambil index dan nilainya
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	// Mengabaikan index menggunakan underscore (_)
	fmt.Println("\nMengabaikan Index:")
	for _, name := range names {
		fmt.Println("Nama :", name)
	}

	//! 4. Break & Continue
	fmt.Println("\n===== 4. Break & Continue =====")
	fmt.Println("Contoh Break (Berhenti di 5):")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("\nContoh Continue (Hanya Ganjil):")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
