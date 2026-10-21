package standardlibrary

import "fmt"

func Fmt() {
	name := "Jang Won-Young"
	age := 23

	// fmt.Printf: Mencetak teks dengan format template (format verb):
	// - %s : Format untuk tipe data string
	// - %d : Format untuk tipe data integer (angka bulat desimal)
	// - \n : Karakter baris baru (newline)
	fmt.Printf("Halo, nama saya %s dan umur saya %d tahun.\n", name, age)
}