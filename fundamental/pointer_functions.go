package fundamental

import "fmt"

type Alamat struct {
	Nama, Kota, Provinsi, Negara string
}

// Function dengan Parameter Pointer (*Alamat) -> Pass by Reference
func ChangeCountryToIndonesia(alamat *Alamat) {
	alamat.Negara = "Indonesia" // Mengubah data di memory asli yang ditunjuk
}

// Function dengan Parameter Biasa (Alamat) -> Pass by Value
func ChangeCountryToMalaysia(alamat Alamat) {
	alamat.Negara = "Malaysia" // Hanya mengubah copy data di dalam scope fungsi
}

func PointerFunctions() {

	//! 1. Parameter Biasa (Pass by Value - Data Asli Tidak Berubah)
	fmt.Println("\n===== 1. Parameter Biasa (Pass by Value) =====")
	alamat1 := Alamat{
		Nama:     "Jang Won-Young",
		Kota:     "Wonogiri",
		Provinsi: "Jawa Tengah",
	}

	ChangeCountryToMalaysia(alamat1)
	fmt.Println("Alamat Asli (Negara Tidak Berubah)  :", alamat1)

	//! 2. Parameter Pointer (Pass by Reference - Data Asli Berubah)
	fmt.Println("\n===== 2. Parameter Pointer (Pass by Reference) =====")
	// Mengirimkan alamat memori menggunakan operator &
	ChangeCountryToIndonesia(&alamat1)
	fmt.Println("Alamat Asli (Negara Berhasil Berubah):", alamat1)
}