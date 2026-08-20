# 📊 Tipe Data Dasar & String

---

## 📊 Tabel Tipe Data Dasar Golang

| Kategori | Tipe Data | Ukuran Memori | Keterangan & Jangkauan Nilai |
| :--- | :--- | :--- | :--- |
| **Boolean** | `bool` | 1 byte | Hanya memiliki dua nilai: `true` (benar) atau `false` (salah). |
| **String** | `string` | Bervariasi | Digunakan untuk teks/kumpulan karakter. Diapit dengan tanda kutip ganda `"..."` atau backtick `` `...` ``. |
| **Integer**<br>*(Bilangan Bulat bertanda)* | `int8` | 8-bit | Bilangan dari -128 sampai 127 |
| | `int16` | 16-bit | Bilangan dari -32768 sampai 32767 |
| | `int32` | 32-bit | Bilangan dari -2147483648 sampai 2147483647 |
| | `int64` | 64-bit | Bilangan dari -9223372036854775808 sampai 9223372036854775807 |
| | `int` | 32 / 64-bit | Mengikuti arsitektur OS. Di OS 32-bit sama dengan `int32`, di OS 64-bit sama dengan `int64`. |
| **Unsigned Integer**<br>*(Bilangan Bulat non-negatif)* | `uint8` | 8-bit | Bilangan dari 0 sampai 255 |
| | `uint16` | 16-bit | Bilangan dari 0 sampai 65535 |
| | `uint32` | 32-bit | Bilangan dari 0 sampai 4294967295 |
| | `uint64` | 64-bit | Bilangan dari 0 sampai 18446744073709551615 |
| | `uint` | 32 / 64-bit | Mengikuti arsitektur OS (seperti `int`, khusus 0 dan angka positif). |
| **Float**<br>*(Bilangan Desimal/Pecahan)* | `float32` | 32-bit | Bilangan desimal presisi tunggal. Contoh: 3.14 |
| | `float64` | 64-bit | Bilangan desimal presisi ganda (lebih detail/presisi dibanding `float32`). |
| **Complex**<br>*(Bilangan Imajiner)* | `complex64` | 64-bit | Bilangan kompleks dengan `float32` untuk bagian real dan imajiner. |
| | `complex128` | 128-bit | Bilangan kompleks dengan `float64` untuk bagian real dan imajiner. |
| **Alias**<br>*(Nama Lain Tipe Data)* | `byte` | 8-bit | Alias untuk `uint8`. Digunakan untuk merepresentasikan data biner/karakter ASCII. |
| | `rune` | 32-bit | Alias untuk `int32`. Digunakan untuk merepresentasikan karakter Unicode (1 karakter utuh). |
| | `uintptr` | 32 / 64-bit | Integer tanpa tanda yang ukurannya cukup untuk menampung alamat memori (pointer). |

---

## 🔤 Function untuk String

| Function | Keterangan |
| :--- | :--- |
| `len("string")` | Menghitung jumlah karakter (byte) di String. |
| `"string"[number]` | Mengambil byte/karakter pada posisi indeks yang ditentukan. |

---

[⬅️ Sebelumnya: Main Function](01-main-function.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Variable & Constant ➡️](03-variable-dan-constant.md)
