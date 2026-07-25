package fundamental

import "fmt"

func Map(){

	//! Cara membuat map 
	var person = map[string]string{
		"name" : "Jang Won-Young",
		"address" : "Jomokerto",
	}

	fmt.Println("Nama : ", person["name"])
	fmt.Println("Alamat : ", person["address"])
	fmt.Println(person)

	//! Function Map 
	//* 1. len(map) -> Untuk mendapatkan jumlah data yang ada di map
	fmt.Println("===== 1. Menghitung Panjang =====")
	var panjang_person = len(person)
	fmt.Println("Panjang Map Person : ", panjang_person)

	//* 2. map[key] -> Mengambil data dari map dengan key
	fmt.Println("===== 2. Mengambil data =====")
	var nama_person = person["name"]
	fmt.Println("Nama Person : ", nama_person)

	//* 3. map[key] = value -> Mengubah data dari map dengan key
	fmt.Println("===== 3. Mengubah data =====")
	person["name"] = "Wan-Young My Bini"
	fmt.Println("Nama Person Now : ", person["name"])

	//* 4. make(map[TypeKey]TypeValue) -> Membuat Map Baru
	fmt.Println("===== 4. Membuat Map Baru =====")
	devices := make(map[string]string)
	devices["laptop"] = "Lenovo LOQ"
	devices["handphone"] = "Redmi note 12"
	fmt.Println("Map Baru : ", devices)

	//* 5. delete(map, key) -> Menghapus data dari map dengan key
	fmt.Println("===== 5. Menghapus data =====")
	delete(devices, "handphone")
	fmt.Println("Map device now : ", devices)
}