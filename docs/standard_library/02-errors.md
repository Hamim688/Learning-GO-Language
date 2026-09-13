# ⚠️ Package errors & Error Checking

---

## 📌 Pengenalan Package `errors`

Di Go, penanganan kesalahan (*error handling*) tidak menggunakan mekanisme `try-catch` seperti di bahasa pemrograman lain, melainkan memanfaatkan interface bawaan `error`.

- **Materi Sebelumnya**: Kita sudah mengenal fungsi `errors.New("pesan error")` untuk membuat instance error sederhana.
- **Kebutuhan Kompleks**: Di dalam aplikasi nyata, kita sering kali perlu membedakan beberapa jenis error yang terjadi (misal: apakah datanya tidak ditemukan, atau inputnya yang tidak valid?) agar aplikasi bisa mengambil keputusan yang tepat.
- **Dokumentasi Resmi**: [https://pkg.go.dev/errors](https://pkg.go.dev/errors)

---

## 🎯 Konsep *Sentinel Error*

**Sentinel Error** adalah pola (*pattern*) di mana kita mendefinisikan variabel-variabel error di tingkat package (*package-level variables*) menggunakan `var`:

```go
var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)
```

### Mengapa Dibuat Variabel Terpisah?
1. **Identitas Unik**: Variabel error ini memiliki alamat memori unik yang dapat dijadikan patokan pembanding.
2. **Reusability**: Tidak perlu membuat string error berulang-ulang di berbagai tempat.
3. **Pembeda Logika**: Memudahkan pemanggil fungsi untuk mendeteksi jenis kesalahan yang terjadi.

---

## 🔍 Mengecek Jenis Error dengan `errors.Is()`

Untuk memeriksa apakah error yang dihasilkan sama dengan jenis error tertentu, gunakan fungsi bawaan:

```go
errors.Is(err, targetError)
```

- Mengembalikan nilai boolean (`true` / `false`).
- **Lebih Aman**: Dibandingkan operator perbandingan langsung (`err == targetError`), `errors.Is()` lebih direkomendasikan karena tetap bekerja meskipun error sudah dibungkus (*wrapped error*).

---

## 💻 Contoh Implementasi di Kode

File: [`standard_library/errors.go`](../../standard_library/errors.go)

```go
package standardlibrary

import (
	"errors"
	"fmt"
)

// Definisi Sentinel Error di level package
var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Won-Young" {
		return NotFoundError
	}

	// Berhasil / Sukses
	return nil
}

func ErrorFunctions() {
	err := GetById("ahyeon")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
```

---

## 💡 Tips Penting: Mencetak Isi Error Asli

Selain menulis pesan custom di dalam blok `if`, kamu juga bisa langsung mengambil pesan dari variabel error-nya:

```go
if errors.Is(err, ValidationError) {
	fmt.Println("Pesan error:", err)        // Output: Pesan error: validation error
	// atau
	fmt.Println("Pesan error:", err.Error())
}
```

---

[⬅️ Sebelumnya: Package fmt](01-fmt.md) | [📚 Indeks Standard Library](README.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Package os ➡️](03-os.md)
