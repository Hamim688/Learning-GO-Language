package fundamental

import (
	"fmt"
)

func Mathematics() {
	var a  = 20
	var b  = 25
	var c = 5


	fmt.Println("\n===== OPERATOR MATEMATIKA DASAR =====")
	// Penjumlahan
	var hasil_penjumlahan = a + b
	fmt.Println("Hasil Penjumlahan : ", hasil_penjumlahan)

	// Pengurangan 
	var hasil_pengurangan = a - b
	fmt.Println("Hasil Pengurangan : ", hasil_pengurangan)

	// Perkalian
	var hasil_perkalian = a * b
	fmt.Println("Hasil Perkalian : ", hasil_perkalian)

	// Pembagian
	var hasil_pembagian = a / c
	fmt.Println("Hasil Pembagian : ", hasil_pembagian)

	// Modulus (Sisa Bagi)
	var hasil_modulus = a % c
	fmt.Println("Hasil Modulus : ", hasil_modulus)



	fmt.Println("\n===== AUGMENTED ASSIGNMENT =====")

	// Penjumlahan
	a += b // sama dengan a = a + b
	fmt.Println("Hasil Augmented Assignment Penjumlahan : ", a)

	// Pengurangan
	a -= b // sama dengan a = a - b
	fmt.Println("Hasil Augmented Assignment Pengurangan : ", a)

	// Perkalian
	a *= b // sama dengan a = a * b
	fmt.Println("Hasil Augmented Assignment Perkalian : ", a)

	// Pembagian
	a /= b // sama dengan a = a / b
	fmt.Println("Hasil Augmented Assignment Pembagian : ", a)

	fmt.Println("\n===== UNARY OPERATOR =====")
	// Unary Operator
	a++ // sama dengan a = a + 1
	fmt.Println("Hasil Unary Operator Penjumlahan : ", a)

	a-- // sama dengan a = a - 1
	fmt.Println("Hasil Unary Operator Pengurangan : ", a)

	fmt.Println("\n===== OPERATOR PERBANDINGAN =====")
	
	// Lebih Dari
	fmt.Println("a > b : ", a > b)

	// Kurang Dari
	fmt.Println("a < b : ", a < b)

	// Lebih Dari Sama Dengan
	fmt.Println("a >= b : ", a >= b)

	// Kurang Dari Sama Dengan
	fmt.Println("a <= b : ", a <= b)

	// Sama Dengan
	fmt.Println("a == b : ", a == b)

	// Tidak Sama Dengan
	fmt.Println("a != b : ", a != b)

	fmt.Println("\n===== OPERATOR BOOLEAN =====")
	
	var nilai_uas = 90
	var absensi = 30

	var lulus_uas = nilai_uas >= 80
	var lulus_absensi = absensi >= 30

	var lulus_semuanya = lulus_uas && lulus_absensi
	
	fmt.Println("Nilai UAS : ", nilai_uas)
	fmt.Println("Absensi : ", absensi)
	fmt.Println("Syarat lulus UAS : ", lulus_uas)
	fmt.Println("Syarat lulus Absensi : ", lulus_absensi)
	fmt.Println("Kesimpulan : ", lulus_semuanya)
}