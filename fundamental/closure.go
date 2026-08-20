package fundamental

import "fmt"

func Closure() {

	//! 1. Closure Sederhana
	fmt.Println("\n===== 1. Closure Sederhana =====")
	// Variabel counter & name berada di outer scope (scope luar)
	counter := 0
	name := "Jang Won-Young"

	// Anonymous function increment dapat mengakses & mengubah (mutate) 
	// data variabel yang berada di luar scope fungsinya sendiri.
	increment := func() {
		name = "Markonah" // Hati-hati: merubah nilai variabel di scope luar!
		fmt.Println("Increment dipanggil...")
		counter++
	}

	// Panggil function increment 3 kali
	increment()
	increment()
	increment()

	// Nilai counter & name di scope luar berubah karena efek closure
	fmt.Println("Counter saat ini :", counter)
	fmt.Println("Name saat ini    :", name)
}