# 📚 Golang Learning Notes

## 🚀 Main Function

Go-Lang is similar to C/C++ programming languages, in that it requires a `main` function.

- The `main` function is a function that will be executed when the program runs.
- To create a function, we can use the `func` keyword.
- The `main` function **must** be located within the `main` package.
- Semicolons (`;`) in Golang are not mandatory, meaning we can choose whether or not to add a semicolon at the end of our program code.

---

## ⚠️ Multiple Main Function

In Golang, functions within a module/project must be unique, meaning we cannot create functions with the same name.

Therefore, if we create a new file, for example `sample.go`, and create a function with the same name (`main`), we will not be able to build the module. This is because the new `main` function would be a duplicate of the `main` function in another file (e.g., `helloworld.go`).

---

## 📊 Tabel Tipe Data Dasar Golang

| Kategori | Tipe Data | Ukuran Memori | Keterangan & Jangkauan Nilai |
| :--- | :--- | :--- | :--- |
| **Boolean** | `bool` | 1 byte | Hanya memiliki dua nilai: `true` (benar) atau `false` (salah). |
| **String** | `string` | Bervariasi | Digunakan untuk teks/kumpulan karakter. Diapit dengan tanda kutip ganda `"..."` atau backtick `` `...` ``. |
| **Integer**<br>*(Bilangan Bulat yang bisa negatif/positif)* | `int8` | 8-bit | Bilangan dari -128 sampai 127 |
| | `int16` | 16-bit | Bilangan dari -32768 sampai 32767 |
| | `int32` | 32-bit | Bilangan dari -2147483648 sampai 2147483647 |
| | `int64` | 64-bit | Bilangan dari -9223372036854775808 sampai 9223372036854775807 |
| | `int` | 32 / 64-bit | Mengikuti arsitektur OS. Di OS 32-bit sama dengan `int32`, di OS 64-bit sama dengan `int64`. |
| **Unsigned Integer**<br>*(Bilangan Bulat yang HANYA positif/nol)* | `uint8` | 8-bit | Bilangan dari 0 sampai 255 |
| | `uint16` | 16-bit | Bilangan dari 0 sampai 65535 |
| | `uint32` | 32-bit | Bilangan dari 0 sampai 4294967295 |
| | `uint64` | 64-bit | Bilangan dari 0 sampai 18446744073709551615 |
| | `uint` | 32 / 64-bit | Mengikuti arsitektur OS (seperti `int`, tapi khusus non-negatif). |
| **Float**<br>*(Bilangan Desimal/Pecahan)* | `float32` | 32-bit | Bilangan desimal presisi tunggal. Contoh: 3.14 |
| | `float64` | 64-bit | Bilangan desimal presisi ganda (lebih detail/presisi di belakang koma dibanding `float32`). |
| **Complex**<br>*(Bilangan Imajiner)* | `complex64` | 64-bit | Bilangan kompleks dengan `float32` untuk real dan imajinernya. |
| | `complex128` | 128-bit | Bilangan kompleks dengan `float64` untuk real dan imajinernya. |
| **Alias**<br>*(Nama Lain Tipe Data)* | `byte` | 8-bit | Alias untuk `uint8`. Sering dipakai untuk merepresentasikan data biner/karakter ASCII. |
| | `rune` | 32-bit | Alias untuk `int32`. Digunakan untuk merepresentasikan karakter Unicode (1 karakter utuh). |
| | `uintptr` | 32 / 64-bit | Integer tanpa tanda yang ukurannya cukup besar untuk menampung alamat memori (pointer). |

---

## 🔤 Function untuk String

| Function | Keterangan |
| :--- | :--- |
| `len("string")` | Menghitung jumlah karakter di String |
| `"string"[number]` | Mengambil karakter pada posisi yang ditentukan |


---

## 📦 Variable

- **Variable** adalah tempat untuk menyimpan data.
- Variable digunakan agar kita bisa mengakses data yang sama di mana pun kita mau.
- Di Go-Lang, variable hanya bisa menyimpan tipe data yang sama. Jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable.
- Untuk membuat variable, kita bisa menggunakan kata kunci `var`, lalu diikuti dengan nama variable dan tipe datanya.

---

## 🔒 Constant

- **Constant** adalah variable yang nilainya **tidak bisa diubah lagi** setelah pertama kali diberi nilai.
- Cara pembuatan constant mirip dengan variable. Yang membedakan hanya kata kunci yang digunakan adalah `const`, bukan `var`.
- Saat pembuatan constant, kita **wajib** langsung menginisialisasikan (mengisi) datanya.

---

## 🔄 Konversi Tipe Data

Di Go-Lang, kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain.
Misalnya kita ingin mengkonversi tipe data `int32` ke `int64`, dan lain-lain.

Khusus untuk tipe data `String` ke tipe data lain (atau sebaliknya), kita harus menggunakan *package* bawaan Golang yaitu `strconv` (String Conversion). Berikut adalah fungsi-fungsi `strconv` yang paling sering dipakai di dunia industri:

### 🛠️ Tabel Fungsi Package `strconv` yang Sering Dipakai

| Fungsi | Kegunaan Utama | Contoh Penggunaan |
| :--- | :--- | :--- |
| `strconv.Atoi()` | Mengubah **String ke Integer** *(ASCII to Integer)*. Sangat sering dipakai untuk membaca input user atau balikan API yang berbentuk teks angka. | `nilai, err := strconv.Atoi("123")` |
| `strconv.Itoa()` | Mengubah **Integer ke String** *(Integer to ASCII)*. Sering dipakai kalau mau menggabungkan atau nge-*print* angka bersama dengan teks. | `teks := strconv.Itoa(123)` |
| `strconv.ParseInt()` | Mirip `Atoi`, tapi lebih advanced. Bisa menentukan basis bilangan (base 10, hex/16, dll) dan alokasi memori (int8, int16, int32, int64). | `nilai64, err := strconv.ParseInt("123", 10, 64)` |
| `strconv.FormatInt()`| Kebalikan dari `ParseInt`. Mengubah `int64` menjadi String dengan spesifikasi basis (base) tertentu. | `teks := strconv.FormatInt(123, 10)` |
| `strconv.ParseBool()`| Mengubah **String ke Boolean**. Hanya mengenali input spesifik (seperti `"true"`, `"false"`, `"1"`, `"0"`). | `b, err := strconv.ParseBool("true")` |
| `strconv.FormatBool()`| Mengubah **Boolean ke String**. Akan selalu menghasilkan string `"true"` atau `"false"`. | `teks := strconv.FormatBool(true)` |