package fundamental

import "fmt"

func Array(){

	fmt.Println("\n===== Deklarasi Array Secara Manual =====")
	var names [3]string
	names[0] = "Jang"
	names[1] = "Won"
	names[2] = "Young"	

	fmt.Println(names)
	fmt.Println(names[0])

	fmt.Println("\n===== Deklarasi Array Secara Langsung =====")

	// Membuat array tanpa deklarasi panjang
	// Tanda [...] artinya "Hitung panjangnya berdasarkan jumlah data di dalamnya"
	var values = [...]int{
		10,
		20,
		30,
		40,
	}
	fmt.Println(values)

	fmt.Println("\n===== Function Array =====")
	// Panjang Array
	panjang := len(values)
	fmt.Println("Panjang Array : ", panjang)

	// Mengambil data di Array
	data_pertama := names[0]
	fmt.Println("Nama Pertama : ", data_pertama)

	// Mengubah data di Array
	names[0] = "Sayangku"
	fmt.Println("Nama Pertama : ", names[0])
	fmt.Println(names)
}
