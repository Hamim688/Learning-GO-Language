package fundamental

import "fmt"

type People struct {
	Name    string
	Address string
	Age     int
}

func Asterisk() {

	//! 1. Mengubah Field Melalui Pointer
	fmt.Println("\n===== 1. Pointer Biasa (Mengubah Field) =====")
	people1 := People{"Jang Won Young", "Wonogiri", 22}
	people2 := &people1

	// Mengubah field Name dari pointer people2
	people2.Name = "Ahyeon"

	// Karena people2 menunjuk memory people1, field Name pada people1 ikut berubah
	fmt.Println("people1 :", people1)
	fmt.Println("people2 :", people2)

	//! 2. Mengubah Variable Pointer TANPA Operator * (Membuat Memory Baru)
	fmt.Println("\n===== 2. Tanpa Operator * (Hanya Mengalihkan Pointer) =====")
	// people2 dialihkan menunjuk ke lokasi memory BARU
	people2 = &People{"Karina", "Jakarta", 20}

	// Efek: people1 TIDAK BERUBAH, hanya people2 yang menunjuk ke alamat memory baru
	fmt.Println("people1 (Tetap Memory Lama)  :", people1)
	fmt.Println("people2 (Menunjuk Memory Baru):", people2)

	//! 3. Mengubah Seluruh Data Memory MENGGUNAKAN Operator *
	fmt.Println("\n===== 3. Menggunakan Operator * (Menimpa Memory Asli) =====")
	// Kita alihkan kembali people2 agar menunjuk ke memory people1
	people2 = &people1

	// Menggunakan operator * untuk menimpa nilai data di memory yang ditunjuk oleh people2
	*people2 = People{"Pharita", "Bandung", 20}

	// Efek: people1 & people2 DUA-DUANYA IKUT BERUBAH menjadi Pharita!
	fmt.Println("people1 (Ikut Berubah) :", people1)
	fmt.Println("people2 (Ikut Berubah) :", people2)
}