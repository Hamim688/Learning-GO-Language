package initialization

import "fmt"

var connection string

//! Function init() dipanggil secara otomatis saat package di-import
func init() {
	fmt.Println("Init: Membuka Koneksi ke Database...")
	connection = "PostgreSQL"
}

//! Public Function untuk mengakses hasil inisialisasi
func GetDatabase() string {
	return connection
}