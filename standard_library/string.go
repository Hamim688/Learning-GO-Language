package standardlibrary

import (
	"fmt"
	"strings"
)

func StringFunctions() {
	// 1. strings.Contains(s, substr)
	// Mengecek apakah string mengandung substring tertentu (case-sensitive) -> mengembalikan boolean (true/false)
	fmt.Println(strings.Contains("Jang Won-Young", "Won-Young")) // true

	// 2. strings.Split(s, separator)
	// Memecah/memotong string berdasarkan separator (pemisah) -> menghasilkan slice of string ([]string)
	fmt.Println(strings.Split("Jang Won-Young", " ")) // [Jang Won-Young]

	// 3. strings.ToLower(s)
	// Mengubah semua huruf menjadi huruf kecil (lowercase)
	fmt.Println(strings.ToLower("Jang Won-Young")) // jang won-young

	// 4. strings.ToUpper(s)
	// Mengubah semua huruf menjadi huruf besar (uppercase)
	fmt.Println(strings.ToUpper("Jang Won-Young")) // JANG WON-YOUNG

	// 5. strings.ReplaceAll(s, from, to)
	// Mengganti semua kemunculan karakter/kata 'from' dengan 'to'
	fmt.Println(strings.ReplaceAll("Jang Won-Young", " ", "-")) // Jang-Won-Young

	// 6. strings.Trim(s, cutset)
	// Menghapus karakter cutset (di sini spasi) yang ada di awal dan akhir string saja
	fmt.Println(strings.Trim("        Jang Won-Young   ", " ")) // Jang Won-Young
}