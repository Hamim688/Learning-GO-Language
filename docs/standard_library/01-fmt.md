# 🖨️ Package fmt (Format I/O)

---

## 📌 Pengenalan Package `fmt`

Package `fmt` mengimplementasikan fungsi format I/O (Input/Output) yang analog dengan fungsi `printf` dan `scanf` pada bahasa pemrograman C.

- **Sering Digunakan**: Sebelumnya kita sudah sangat sering menggunakan package `fmt`, terutama function `fmt.Println()` untuk mencetak output ke konsol.
- **Banyak Fungsi Formatting**: Selain `Println`, masih banyak function penting lain di package `fmt`, terutama function-function yang digunakan untuk melakukan formatting teks/data.
- **Dokumentasi Resmi**: [https://pkg.go.dev/fmt](https://pkg.go.dev/fmt)

---

## ⚙️ Function Populer di Package `fmt`

| Function | Keterangan | Contoh |
| :--- | :--- | :--- |
| `fmt.Print()` | Mencetak teks tanpa baris baru (*newline*) di akhir | `fmt.Print("Halo")` |
| `fmt.Println()` | Mencetak teks dan otomatis menambahkan baris baru (`\n`) di akhir | `fmt.Println("Halo Dunia")` |
| `fmt.Printf()` | Mencetak teks dengan format template (*format verbs*) tertentu | `fmt.Printf("Nama: %s", name)` |
| `fmt.Sprint()` / `fmt.Sprintf()` / `fmt.Sprintln()` | Melakukan format teks tetapi **tidak mencetak ke layar**, melainkan mengembalikan (*return*) hasilnya sebagai tipe data `string` | `hasil := fmt.Sprintf("%s berumur %d", name, age)` |

---

## 🔤 Verb Format Populer pada `fmt.Printf()`

Saat menggunakan `fmt.Printf()` atau `fmt.Sprintf()`, kita menggunakan karakter khusus (*verb*) yang diawali tanda persen (`%`):

- `%s` : Format untuk string
- `%d` : Format untuk angka integer (basis 10)
- `%f` : Format untuk angka desimal / floating point (misal: `%.2f` untuk 2 angka di belakang koma)
- `%t` : Format untuk tipe data boolean (`true` / `false`)
- `%v` : Format *default value* apa saja (sering dipakai saat debug struct/map/array)
- `%+v`: Format struct dengan menyertakan nama field-nya
- `%T` : Format untuk mencetak tipe data dari variabel tersebut
- `%%` : Mencetak tanda persen literal (`%`)

---

## 💻 Contoh Implementasi di Kode

File: [`standard_library/fmt.go`](../../standard_library/fmt.go)

```go
package standardlibrary

import "fmt"

func Fmt() {
	name := "Jang Won-Young"
	age := 23

	// Menggunakan Printf dengan format %s (string) dan %d (integer)
	fmt.Printf("Halo, nama saya %s dan umur saya %d tahun.\n", name, age)
}
```

---

[⬅️ Kembali ke Indeks Standard Library](README.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Package errors ➡️](02-errors.md)
