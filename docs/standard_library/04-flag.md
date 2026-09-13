# 🚩 Package flag (Command Line Arguments Parsing)

---

## 📌 Pengenalan Package `flag`

Package `flag` mengimplementasikan fungsionalitas untuk melakukan *parsing* terhadap argumen baris perintah (*command line arguments / flags*).

- **Perbedaan `os.Args` vs `flag`**:
  - `os.Args`: Hanya mengambil data mentah berdasarkan urutan indeks slice (`os.Args[1]`, `os.Args[2]`).
  - `flag`: Mengambil argumen berbasis nama/opsi (*named arguments*, misal: `-username=admin -port=8080`) dengan tipe data otomatis dan nilai default (*fallback value*).
- **Fitur Otomatis Help**: Package `flag` otomatis membuatkan pesan bantuan (*help message*) ketika user menjalankan program dengan flag `-h` atau `--help`.
- **Dokumentasi Resmi**: [https://pkg.go.dev/flag](https://pkg.go.dev/flag) atau [https://golang.org/pkg/flag/](https://golang.org/pkg/flag/)

---

## ⚙️ 3 Tahap Penggunaan `package flag`

1. **Mendefinisikan Flag**:
   Gunakan fungsi seperti `flag.String()`, `flag.Int()`, `flag.Bool()`, dll.
   ```go
   var username *string = flag.String("username", "root", "database username")
   ```
   > ⚠️ **Catatan Penting**: Fungsi ini mengembalikan nilai bertipe **Pointer** (`*string`, `*int`, dsb), bukan nilai langsung.

2. **Mem-parsing Flag**:
   Panggil fungsi `flag.Parse()` setelah seluruh definisi flag selesai dibuat dan sebelum variabelnya digunakan:
   ```go
   flag.Parse()
   ```

3. **Mengambil Nilai (Dereferencing)**:
   Karena variabelnya bertipe pointer, gunakan operator asterisk (`*`) untuk mengambil nilainya:
   ```go
   fmt.Println("Username:", *username)
   ```

---

## 💻 Contoh Implementasi di Kode

File: [`standard_library/flag.go`](../../standard_library/flag.go)

```go
package standardlibrary

import (
	"flag"
	"fmt"
)

func FlagFunctions() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "", "database password")
	var address *string = flag.String("address", "localhost", "database address")
	var port *int = flag.Int("port", 0, "database port")

	// Wajib dipanggil untuk mem-parsing input terminal
	flag.Parse()

	// Mengakses nilai dengan dereference operator (*)
	fmt.Println("Username : ", *username)
	fmt.Println("Password : ", *password)
	fmt.Println("Address  : ", *address)
	fmt.Println("Port     : ", *port)
}
```

---

## 🚀 Cara Menjalankan dengan Parameter di Terminal

Kita bisa mengoper argumen saat menjalankan program dengan berbagai variasi sintaks:

```bash
# Menggunakan tanda sama dengan (=)
go run main.go -username=ahyeon -password=rahasia -port=8080

# Menggunakan spasi
go run main.go -username ahyeon -port 8080

# Menampilkan bantuan otomatis (Help)
go run main.go -help
```

---

[⬅️ Sebelumnya: Package os](03-os.md) | [📚 Indeks Standard Library](README.md) | [🏠 Hub Utama](../../LEARNING.md)
