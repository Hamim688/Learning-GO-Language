package fundamental

import "fmt"

func endAppPanic() {
	fmt.Println("Aplikasi selesai.")
}

func runAppPanic(error bool) {
	defer endAppPanic()
	fmt.Println("Aplikasi berjalan...")
	if error {
		panic("terjadi error")
	}
}

func PanicFunction() {
	//! Panic Function
	fmt.Println("\n===== Panic Function =====")
	runAppPanic(true)
}
