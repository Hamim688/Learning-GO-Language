# 🔤 Package strings (Manipulasi String)

---

## 📌 Pengenalan Package `strings`

Di Golang, tipe data `string` bersifat *immutable* (nilainya tidak bisa diubah secara langsung di alamat memori aslinya). Untuk melakukan pengolahan, pemotongan, transformasi huruf, maupun pencarian teks, Go menyediakan package bawaan bernama **`strings`**.

- **Fungsi Utama**: Menyediakan utilitas lengkap untuk memanipulasi dan memproses teks string.
- **Banyak Fungsi Siap Pakai**: Terdapat puluhan fungsi pembantu untuk menangani kasus manipulasi teks yang umum digunakan di industri.
- **Dokumentasi Resmi**: [https://pkg.go.dev/strings](https://pkg.go.dev/strings) atau [https://golang.org/pkg/strings/](https://golang.org/pkg/strings/)

---

## ⚙️ Function Populer di Package `strings`

Berikut adalah beberapa fungsi yang paling sering digunakan:

| Function | Kegunaan | Contoh Input ➡️ Output |
| :--- | :--- | :--- |
| `strings.Contains(s, search)` | Mengecek apakah suatu string mengandung teks tertentu (return `bool`) | `Contains("Golang", "lang")` ➡️ `true` |
| `strings.Split(s, separator)` | Memotong string berdasarkan pemisah (*separator*) menjadi slice `[]string` | `Split("a,b,c", ",")` ➡️ `["a", "b", "c"]` |
| `strings.ToLower(s)` | Mengubah semua karakter string menjadi huruf kecil (*lowercase*) | `ToLower("GOLANG")` ➡️ `"golang"` |
| `strings.ToUpper(s)` | Mengubah semua karakter string menjadi huruf besar (*uppercase*) | `ToUpper("golang")` ➡️ `"GOLANG"` |
| `strings.Trim(s, cutset)` | Memotong/menghapus karakter tertentu (`cutset`) di awal dan akhir string | `Trim("  Halo  ", " ")` ➡️ `"Halo"` |
| `strings.ReplaceAll(s, from, to)` | Mengganti seluruh kemunculan teks tertentu dengan teks baru | `ReplaceAll("A-B-A", "A", "C")` ➡️ `"C-B-C"` |

### 💡 Function Tambahan yang Berguna:
- `strings.TrimSpace(s)`: Menghapus spasi dan karakter whitespace (`\n`, `\t`) di awal dan akhir string.
- `strings.Join(slice, sep)`: Menggabungkan elemen-elemen slice `[]string` menjadi satu string utuh dengan pemisah tertentu (kebalikan dari `Split`).
- `strings.HasPrefix(s, prefix)`: Mengecek apakah string diawali dengan kata tertentu.
- `strings.HasSuffix(s, suffix)`: Mengecek apakah string diakhiri dengan kata tertentu.

---

## 💻 Contoh Implementasi di Kode

File: [`standard_library/string.go`](../../standard_library/string.go)

```go
package standardlibrary

import (
	"fmt"
	"strings"
)

func StringFunctions() {
	// 1. strings.Contains - Pengecekan substring (case-sensitive)
	fmt.Println(strings.Contains("Jang Won-Young", "Won-Young")) // true

	// 2. strings.Split - Memecah string menjadi slice berdasarkan spasi
	fmt.Println(strings.Split("Jang Won-Young", " "))            // [Jang Won-Young]

	// 3. strings.ToLower - Mengubah ke huruf kecil
	fmt.Println(strings.ToLower("Jang Won-Young"))               // jang won-young

	// 4. strings.ToUpper - Mengubah ke huruf besar
	fmt.Println(strings.ToUpper("Jang Won-Young"))               // JANG WON-YOUNG

	// 5. strings.ReplaceAll - Mengganti semua spasi dengan strip (-)
	fmt.Println(strings.ReplaceAll("Jang Won-Young", " ", "-")) // Jang-Won-Young

	// 6. strings.Trim - Menghapus karakter spasi di awal dan akhir string
	fmt.Println(strings.Trim("        Jang Won-Young   ", " "))  // Jang Won-Young
}
```

---

[⬅️ Sebelumnya: Package flag](04-flag.md) | [📚 Indeks Standard Library](README.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Package strconv ➡️](06-strconv.md)
