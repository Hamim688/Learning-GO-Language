package standardlibrary

import (
	"fmt"
	"os"
)

func OsFunction() {
	// 1. os.Args: Berisi slice []string argumen baris perintah saat aplikasi dijalankan
	// - os.Args[0] : Path file binary/eksekusi program itu sendiri
	// - os.Args[1:] : Argumen/parameter tambahan yang diketikkan pengguna via terminal
	args := os.Args
	for _, arg := range args {
		fmt.Println("Argumen CLI:", arg)
	}

	// 2. os.Hostname(): Mengambil nama perangkat / komputer yang sedang menjalankan program
	// Mengembalikan 2 nilai: (hostname string, err error)
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname   :", hostname)
	} else {
		fmt.Println("Error hostname:", err.Error())
	}
}