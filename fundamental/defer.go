package fundamental

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai.")
}

func runApp() {
	defer endApp()
	fmt.Println("Aplikasi berjalan...")
}

func DeferFunction() {
	//! Defer Function
	fmt.Println("\n===== Defer Function =====")
	runApp()
}