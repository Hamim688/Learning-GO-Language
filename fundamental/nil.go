package fundamental

import "fmt"

//! Helper Function Pengecekan Data dengan Nil
func CekName(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func NilFunction() {
	//! 1. Penggunaan Data Nil pada Map & Function
	fmt.Println("\n===== 1. Data Nil (Data Kosong) =====")
	
	// Pengujian 1: Mengirim string kosong -> Mengembalikan nil
	dataKosong := CekName("")
	if dataKosong == nil {
		fmt.Println("Data Masih Kosong (Nil)")
	} else {
		fmt.Println("Nama :", dataKosong["name"])
	}

	//! Pengujian 2: Mengirim nama "Yanto" -> Mengembalikan Map
	fmt.Println("\n===== 2. Data Map (Data Yanto) =====")

	// Pengujian 2: Mengirim nama "Yanto" -> Mengembalikan Map
	dataMap := CekName("Yanto")
	if dataMap == nil {
		fmt.Println("Data Masih Kosong (Nil)")
	} else {
		fmt.Println("Nama :", dataMap["name"])
	}
}