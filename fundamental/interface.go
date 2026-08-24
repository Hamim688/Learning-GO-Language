package fundamental

import "fmt"

//! Definisi Interface (Kontrak)
type HasName interface {
	GetName() string
}

func SaySayang(value HasName) {
	fmt.Println("Hello Sayang", value.GetName())
}

//! Implementasi Interface Pertama (Struct Person)
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

//! Implementasi Interface Kedua (Struct Animal)
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Interface() {

	//! 1. Implementasi Interface Implicit pada Struct Person & Animal
	fmt.Println("\n===== 1. Implementasi Interface =====")
	person := Person{
		Name: "Won Young",
	}
	animal := Animal{
		Name: "Kucing",
	}

	SaySayang(person)
	SaySayang(animal)
}
