package fundamental

import "fmt"

//! Penggunaan Recover yang salah
/*
func endAppRecover() {
	fmt.Println("Aplikasi selesai.")
}

func runAppRecover(error bool) {
	fmt.Println("Aplikasi berjalan...")
	if error {
		panic("terjadi error")
	}
	message := recover()
	fmt.Println("Error: ", message)
	defer endAppRecover()
}
*/

//! Penggunaan Recover yang benar
func endAppRecover() {
	fmt.Println("Aplikasi selesai.")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
}

func runAppRecover(error bool) {
	defer endAppRecover()
	fmt.Println("Aplikasi berjalan...")
	if error {
		panic("Upss Error")
	}
}

func RecoverFunction() {
	//! Recover Function
	fmt.Println("\n===== Recover Function =====")
	runAppRecover(true)
	fmt.Println("Jang Won-Young (Program tetap berjalan setelah recover)")
}