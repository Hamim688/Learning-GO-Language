package standardlibrary

import (
	"fmt"
	"strconv"
)

func StrconvFunctions() {
	// 1. strconv.Atoi(string): ASCII to Integer (string -> int)
	// Mengembalikan 2 nilai: (int, error) jika string bukan angka valid
	value, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Atoi (string -> int)      :", value)
	}

	// 2. strconv.Itoa(int): Integer to ASCII (int -> string)
	// Kebalikan dari Atoi, tidak menghasilkan error
	valuestr := strconv.Itoa(100)
	fmt.Println("Itoa (int -> string)      :", valuestr+"00") // Penggabungan string "100" + "00" = "10000"

	// 3. strconv.ParseBool(string): Mengubah string ("true"/"false"/"1"/"0") ke tipe data boolean
	valuebool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("ParseBool (string -> bool):", valuebool)
	}

	// 4. strconv.ParseFloat(s, bitSize): Mengubah string ke float (biasanya bitSize 64 untuk float64)
	valueFloat, err := strconv.ParseFloat("10.5", 64)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("ParseFloat                :", valueFloat)
	}

	// 5. strconv.ParseInt(s, base, bitSize): Mengubah string angka ke int64
	// base 10 = desimal, bitSize 64 = int64
	valueInt, err := strconv.ParseInt("10", 10, 64)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("ParseInt                  :", valueInt)
	}

	// 6. strconv.FormatBool(bool): Mengubah tipe boolean ke bentuk string ("true"/"false")
	fmt.Println("FormatBool (bool -> string):", strconv.FormatBool(true))

	// 7. strconv.FormatFloat(f, fmt, prec, bitSize): Mengubah float ke string
	// 'f' = tanpa eksponen, 2 = 2 angka di belakang koma, 64 = float64
	fmt.Println("FormatFloat               :", strconv.FormatFloat(10.5, 'f', 2, 64)) // "10.50"

	// 8. strconv.FormatInt(i, base): Mengubah int64 ke bentuk string dengan basis tertentu (10 = desimal)
	fmt.Println("FormatInt                 :", strconv.FormatInt(10, 10))
}
	