package fundamental

import "fmt"

type Mahasiswa struct {
	Nama string
}

func NewFunction() {
	//! 1. Penggunaan Operator new() untuk Membuat Pointer ke Data Kosong
	fmt.Println("\n===== 1. Operator new =====")
	
	// Function new(TipeData) membuat pointer yang menunjuk ke memori baru berisi zero-value
	var Mahasiswa1 *Mahasiswa = new(Mahasiswa)
	var Mahasiswa2 *Mahasiswa = Mahasiswa1

	// Mengubah data melalui Mahasiswa2
	Mahasiswa2.Nama = "Jang Won-Young"

	// Karena Mahasiswa2 & Mahasiswa1 menunjuk ke alamat memori yang sama, Mahasiswa1 ikut berubah
	fmt.Println("Mahasiswa 1 (Ikut Berubah) :", Mahasiswa1.Nama)
	fmt.Println("Mahasiswa 2 (Pointer)      :", Mahasiswa2.Nama)
}