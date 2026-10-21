package standardlibrary

import (
	"flag"
	"fmt"
)

func FlagFunctions() {
	// Definisi Flag: flag.Type(nama_flag, default_value, deskripsi)
	// PENTING: flag.String / flag.Int mengembalikan nilai bertipe POINTER (*string, *int)!
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "", "database password")
	var address *string = flag.String("address", "localhost", "database address")
	var port *int = flag.Int("port", 0, "database port")

	// flag.Parse() WAJIB dipanggil setelah semua flag dideklarasikan
	// Berfungsi untuk membaca dan mem-parsing input argumen dari terminal (contoh: -username=admin)
	flag.Parse()

	// Karena variabel di atas berupa pointer, gunakan tanda bintang (*) untuk mengambil nilainya (dereference)
	fmt.Println("Username : ", *username)
	fmt.Println("Password : ", *password)
	fmt.Println("Address  : ", *address)
	fmt.Println("Port     : ", *port)
}