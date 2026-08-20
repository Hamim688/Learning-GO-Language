package fundamental

import "fmt"

// 1. Function Sederhana
func sayHello() {
	fmt.Println("Hello Jang Won-Young!")
}

// 2. Function Parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// 3. Function Return Value
func getHello(name string) string {
	return "Hello " + name
}

// 4. Returning Multiple Values
func getFullName() (string, string) {
	return "Jang", "Won-Young"
}

// 5. Named Return Values
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Jang"
	middleName = "Won"
	lastName = "Young"
	return firstName, middleName, lastName
}

// 6. Variadic Function & Slice Parameter
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// 7. Function Value
func getGoodbye(name string) string {
	return "Good Bye " + name
}

// 8. Function Type Declaration
type Filter func(string) string

// 9. Function as Parameter
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Kasar" {
		return "..."
	} else {
		return name
	}
}

// 10. Anonymous Function
type Blacklist func(string) bool
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// 11. Recursive Function
func factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}

func Function() {

	//! 1. Function Sederhana
	fmt.Println("\n===== 1. Function Sederhana =====")
	sayHello()

	//! 2. Function Parameter
	fmt.Println("\n===== 2. Function Parameter =====")
	sayHelloTo("Jang", "Won-Young")

	//! 3. Function Return Value
	fmt.Println("\n===== 3. Function Return Value =====")
	result := getHello("My Bini ❤️")
	fmt.Println(result)

	//! 4. Returning Multiple Values & Menghiraukan Value
	fmt.Println("\n===== 4. Multiple Return Values =====")
	firstName, lastName := getFullName()
	fmt.Println("Full Name :", firstName, lastName)

	// Menghiraukan return value kedua
	firstNameOnly, _ := getFullName()
	fmt.Println("First Name Only :", firstNameOnly)

	//! 5. Named Return Values
	fmt.Println("\n===== 5. Named Return Values =====")
	first, middle, last := getCompleteName()
	fmt.Println("Named Return :", first, middle, last)

	//! 6. Variadic Function & Slice Parameter
	fmt.Println("\n===== 6. Variadic Function =====")
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println("Total Varargs :", total)

	// Mengirimkan Slice ke Variadic Function
	numbers := []int{5, 10, 15, 20}
	totalSlice := sumAll(numbers...)
	fmt.Println("Total dari Slice Parameter :", totalSlice)

	//! 7. Function Value
	fmt.Println("\n===== 7. Function Value =====")
	goodbye := getGoodbye
	fmt.Println(goodbye("Markonah"))

	//! 8. Function as Parameter & Type Declaration
	fmt.Println("\n===== 8. Function as Parameter & Type Declaration =====")
	sayHelloWithFilter("Jang Won-Young", spamFilter)
	sayHelloWithFilter("Kasar", spamFilter)

	//! 9. Anonymous Function
	fmt.Println("\n===== 9. Anonymous Function =====")
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Joko", blacklist)
	registerUser("Babi", func(name string) bool {
		if name == "Babi" {
			return true
		} else {
			return false
		}
	})

	//! 10. Recursive Function
	fmt.Println("\n===== 10. Recursive Function =====")
	fact10 := factorial(10)
	fmt.Println("Factorial 10 :", fact10)

}
