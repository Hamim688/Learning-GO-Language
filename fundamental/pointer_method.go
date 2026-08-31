package fundamental

import "fmt"

type Woman struct {
	Name string
}

//! Method dengan Value Receiver (Pass by Value)
func (woman Woman) Married() {
	woman.Name = "Mrs. " + woman.Name // Hanya mengubah copy data di dalam scope method
}

//! Method dengan Pointer Receiver (Pass by Reference)
func (woman *Woman) ChangeName(name string) {
	woman.Name = name // Mengubah data asli pada struct
}

func PointerMethod() {
	//! 1. Pengujian Method Biasa (Value Receiver)
	fmt.Println("\n===== 1. Method Biasa (Pass by Value) =====")
	woman := Woman{"Jang Won-Young"}
	woman.Married()
	fmt.Println("Method Biasa (Data Asli Tidak Berubah)  :", woman.Name)

	//! 2. Pengujian Method Pointer (Pointer Receiver)
	fmt.Println("\n===== 2. Method Pointer (Pass by Reference) =====")
	woman.ChangeName("Ahyeon")
	fmt.Println("Method Pointer (Data Asli Berhasil Berubah):", woman.Name)
}