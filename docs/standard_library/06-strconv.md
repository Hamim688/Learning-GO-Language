# 🔄 Package strconv (String Conversion)

---

## 📌 Pengenalan Package `strconv`

Sebelumnya di materi Fundamental, kita sudah mempelajari konversi antar tipe data yang sejenis (misalnya dari `int32` ke `int64`, atau `float32` ke `float64`). Namun, bagaimana jika kita ingin mengonversi antar tipe data yang berbeda, seperti mengubah angka `int` menjadi `string`, atau sebaliknya dari teks `string` menjadi angka?

Di Go, hal tersebut tidak bisa dilakukan dengan *casting* biasa `int("100")`, melainkan membutuhkan bantuan package bawaan bernama **`strconv`** (*String Conversion*).

- **Dokumentasi Resmi**: [https://pkg.go.dev/strconv](https://pkg.go.dev/strconv) atau [https://golang.org/pkg/strconv/](https://golang.org/pkg/strconv/)

---

## ⚙️ 2 Kategori Utama Fungsi di `strconv`

Fungsi di dalam package `strconv` secara umum terbagi menjadi dua kelompok besar:

### 1. Fungsi `Parse...` (String ➡️ Tipe Data Lain)
Digunakan untuk mengubah data dari `string` menjadi tipe data primitif (`bool`, `int`, `float`).
> ⚠️ **Catatan Penting**: Semua fungsi `Parse...` selalu mengembalikan **2 nilai** yaitu `(hasil, error)`. Jika string yang dimasukkan tidak valid (misal: `"halo"` diubah ke angka), maka variabel `error` tidak akan bernilai `nil`.

| Function | Kegunaan | Return Value |
| :--- | :--- | :--- |
| `strconv.ParseBool(s)` | Mengubah string (`"true"`, `"false"`, `"1"`, `"0"`) ke `bool` | `(bool, error)` |
| `strconv.ParseFloat(s, bitSize)` | Mengubah string desimal ke `float` (`bitSize`: 32 / 64) | `(float64, error)` |
| `strconv.ParseInt(s, base, bitSize)` | Mengubah string angka ke `int64` (`base`: 10 untuk desimal, 2 biner, 16 heksa) | `(int64, error)` |
| `strconv.Atoi(s)` | *ASCII to Integer*: Shortcut instan untuk konversi string langsung ke `int` | `(int, error)` |

---

### 2. Fungsi `Format...` (Tipe Data Lain ➡️ String)
Digunakan untuk mengubah tipe data primitif menjadi representasi `string`.
> 💡 **Info**: Fungsi `Format...` **tidak mengembalikan error** karena tipe data primitif selalu valid untuk diubah menjadi string.

| Function | Kegunaan | Parameter Penting |
| :--- | :--- | :--- |
| `strconv.FormatBool(b)` | Mengubah nilai `bool` menjadi `"true"` atau `"false"` | `b bool` |
| `strconv.FormatFloat(f, fmt, prec, bitSize)` | Mengubah float ke string | `fmt`: format ('f' desimal biasa), `prec`: jumlah angka di belakang koma, `bitSize`: 64 |
| `strconv.FormatInt(i, base)` | Mengubah `int64` ke string berbasis tertentu | `base`: 10 (desimal) |
| `strconv.Itoa(i)` | *Integer to ASCII*: Shortcut instan untuk konversi `int` langsung ke `string` | `i int` |

---

## 💻 Contoh Implementasi di Kode

File: [`standard_library/strconv.go`](../../standard_library/strconv.go)

```go
package standardlibrary

import (
	"fmt"
	"strconv"
)

func StrconvFunctions() {
	// 1. Atoi (string -> int)
	value, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Atoi (string -> int)      :", value)
	}

	// 2. Itoa (int -> string)
	valuestr := strconv.Itoa(100)
	fmt.Println("Itoa (int -> string)      :", valuestr+"00") // "100" + "00" = "10000"

	// 3. ParseBool (string -> bool)
	valuebool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("ParseBool (string -> bool):", valuebool)
	}

	// 4. ParseFloat (string -> float64)
	valueFloat, err := strconv.ParseFloat("10.5", 64)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("ParseFloat                :", valueFloat)
	}

	// 5. ParseInt (string -> int64 dengan basis 10)
	valueInt, err := strconv.ParseInt("10", 10, 64)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("ParseInt                  :", valueInt)
	}

	// 6. FormatBool (bool -> string)
	fmt.Println("FormatBool (bool -> string):", strconv.FormatBool(true))

	// 7. FormatFloat (float64 -> string)
	fmt.Println("FormatFloat               :", strconv.FormatFloat(10.5, 'f', 2, 64)) // "10.50"

	// 8. FormatInt (int64 -> string)
	fmt.Println("FormatInt                 :", strconv.FormatInt(10, 10))
}
```

---

[⬅️ Sebelumnya: Package strings](05-strings.md) | [📚 Indeks Standard Library](README.md) | [🏠 Hub Utama](../../LEARNING.md)
