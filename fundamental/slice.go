package fundamental

import "fmt"

func Slice(){
	fmt.Println("\n===== Tipe Data Slice =====")

	var names = [...]int{
		1,
		2,
		3,
		4,
		5,
	}
	
	// array[low:high] -> mengambil data array dari index low sampai sebelum index high
	slice1 := names[0:2]
	fmt.Println(slice1)

	// array[low:] -> mengambil data array dari index low sampai index akhir array
	slice2 := names[2:]
	fmt.Println(slice2)

	// array[:high] -> mengambil data array dari index 0 sampai sebelum index high
	slice3 := names[:3]
	fmt.Println(slice3)

	// array[:] -> mengambil semua data array
	slice4 := names[:]
	fmt.Println(slice4)


	fmt.Println("===== Function Slice =====")
	// kita buat array terlebih dahulu
	days := [7]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	//! len() -> panjang slice
	fmt.Println("\n===== Function Len =====")
	fmt.Println("Panjang dari array days : " ,len(days))

	//! cap() -> kapasitas slice
	fmt.Println("\n===== Function Cap =====")
	fmt.Println("Kapasitas dari array days : " , cap(days))

	//! append() -> membuat slice baru dengan menambah data ke posisi terakhir slice
	fmt.Println("\n===== Function Append =====")
	// memotong array menjadi slice, yaitu mengambil data dari index tertentu
	// Ambil data dari hari ke-5 (Sabtu) sampai akhir
	newSlice := days[5:]
	fmt.Println("Slice Baru : ", newSlice)

	// Kita ubah data pada slice baru	
	newSlice[0] = "Sabtu_libur"
	fmt.Println("Slice Baru Setelah diubah index ke 0 nya : ", newSlice)

	// Kita ubah lagi data pada slice baru
	newSlice[1] = "Minggu_libur" 
	fmt.Println("Slice Baru Setelah diubah index ke 1 nya : ", newSlice)

	// Cek apakah array aslinya ikut berubah ??? ya ternyata ikut berubah
	fmt.Println("Array Asli : ", days)

	// Coba tambahkan datanya
	newSlice2 := append(newSlice, "Senin_Libur")
	fmt.Println("Slice Baru 2 Setelah di append : ", newSlice2)

	//! make() -> cara membuat slice menggunakan make 
	fmt.Println("\n===== Function Make =====")
	colors := make([]string, 3, 5) // parameter 1 : tipe data, parameter 2 : panjang slice, parameter 3 : kapasitas slice
	
	colors[0] = "Merah"
	colors[1] = "Kuning"
	colors[2] = "Hijau"
	
	// Tampilkan hasil slicenya
	fmt.Println("Colors : ", colors)

	// Coba tambahkan data pada slice colors
	colors = append(colors, "Biru")
	fmt.Println("Colors Setelah di append : ", colors)

	//! copy() -> menyalin data dari slice satu ke slice lain
	fmt.Println("\n===== Function Copy =====")
	fromSlice := colors[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println("From Slice : ", fromSlice)
	fmt.Println("To Slice : ", toSlice)

}