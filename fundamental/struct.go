package fundamental

import "fmt"

//! Definisi Struct
type Customer struct {
	Name, Address string
	Age           int
}

//! Struct Method
func (customer Customer) SayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func Struct() Customer {

	//! 1. Deklarasi & Pengisian Field Struct
	fmt.Println("\n===== 1. Deklarasi Struct =====")
	var WonYoung Customer
	WonYoung.Name = "Won Young"
	WonYoung.Address = "Seoul"
	WonYoung.Age = 22

	fmt.Println("Name    :", WonYoung.Name)
	fmt.Println("Address :", WonYoung.Address)
	fmt.Println("Age     :", WonYoung.Age)

	//! 2. Membuat Struct Literal (Dengan Key)
	fmt.Println("\n===== 2. Struct Literal (Dengan Key) =====")
	ahyeon := Customer{
		Name:    "Ahyeon",
		Address: "Seoul",
		Age:     22,
	}
	fmt.Println(ahyeon)

	//! 3. Membuat Struct Literal (Tanpa Key / Position Base)
	fmt.Println("\n===== 3. Struct Literal (Tanpa Key) =====")
	karina := Customer{
		"Karina", "Seoul", 22,
	}
	fmt.Println(karina)

	//! 4. Memanggil Struct Method
	fmt.Println("\n===== 4. Struct Method =====")
	WonYoung.SayHello("Sayangku")

	return WonYoung
}