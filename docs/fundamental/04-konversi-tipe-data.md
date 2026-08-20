# 🔄 Konversi Tipe Data

---

## 🔄 Konversi Tipe Data

Di Go-Lang, kadang kita butuh melakukan konversi dari satu tipe data ke tipe data lain (misalnya `int32` ke `int64`).

Khusus untuk konversi antara tipe data `String` dan tipe data lain (seperti angka atau boolean), kita wajib menggunakan *package* bawaan Golang yaitu `strconv` (*String Conversion*).

### 🛠️ Tabel Fungsi Package `strconv` yang Sering Dipakai

| Fungsi | Kegunaan Utama | Contoh Penggunaan |
| :--- | :--- | :--- |
| `strconv.Atoi()` | Mengubah **String ke Integer** *(ASCII to Integer)*. Sangat sering dipakai untuk membaca input user atau balikan API berupa teks angka. | `nilai, err := strconv.Atoi("123")` |
| `strconv.Itoa()` | Mengubah **Integer ke String** *(Integer to ASCII)*. Digunakan saat ingin menggabungkan angka ke dalam kalimat string. | `teks := strconv.Itoa(123)` |
| `strconv.ParseInt()` | Mirip `Atoi`, tapi lebih fleksibel. Bisa menentukan basis bilangan (base 10, hex/16, dll) dan alokasi memori (int8, int16, int32, int64). | `nilai64, err := strconv.ParseInt("123", 10, 64)` |
| `strconv.FormatInt()`| Kebalikan dari `ParseInt`. Mengubah `int64` menjadi String dengan spesifikasi basis tertentu. | `teks := strconv.FormatInt(123, 10)` |
| `strconv.ParseBool()`| Mengubah **String ke Boolean**. Mengubah teks seperti `"true"`, `"false"`, `"1"`, `"0"` menjadi `bool`. | `b, err := strconv.ParseBool("true")` |
| `strconv.FormatBool()`| Mengubah **Boolean ke String**. Menghasilkan nilai string `"true"` atau `"false"`. | `teks := strconv.FormatBool(true)` |

---

[⬅️ Sebelumnya: Variable & Constant](03-variable-dan-constant.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Type Declarations ➡️](05-type-declarations.md)
