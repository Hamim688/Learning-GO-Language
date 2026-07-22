<div align="center">

# 🦫 Golang Fundamental Learning Notes 🚀

![Go Language](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/Status-Learning-success?style=for-the-badge)
![Documentation](https://img.shields.io/badge/Notes-Structured-blue?style=for-the-badge)

*Catatan lengkap & terstruktur belajar pemrograman bahasa **Go (Golang)** dari dasar hingga mahir.*

---

</div>

## 📌 Daftar Isi

- [🎯 Main Function](#-main-function)
- [🚨 Multiple Main Function](#-multiple-main-function)
- [📊 Tabel Tipe Data Dasar Golang](#-tabel-tipe-data-dasar-golang)
- [🔤 Function untuk String](#-function-untuk-string)
- [🏷️ Variable](#%EF%B8%8F-variable)
- [🔐 Constant](#-constant)
- [🔄 Konversi Tipe Data](#-konversi-tipe-data)
  - [🛠️ Tabel Fungsi Package `strconv`](#%EF%B8%8F-tabel-fungsi-package-strconv-yang-sering-dipakai)
- [📛 Type Declarations](#-type-declarations)
- [🧮 Mathematics & Operators](#-mathematics--operators)
  - [🔧 Tabel Operasi Matematika Dasar](#-tabel-operasi-matematika-dasar)
  - [➕ Augmented Assignment](#-augmented-assignment-operasi-singkat)
  - [⬆️ Tabel Unary Operator](#%EF%B8%8F-tabel-unary-operator)
  - [⚖️ Operator Perbandingan](#%EF%B8%8F-operator-perbandingan)
  - [🔀 Operator Boolean](#-operator-boolean)
- [🧱 Tipe Data Array](#-tipe-data-array)
- [🍕 Tipe Data Slice](#-tipe-data-slice)
  - [🔍 3 Komponen Utama Tipe Data Slice](#-3-komponen-utama-tipe-data-slice)
  - [✂️ Cara Membuat Slice dari Array](#%EF%B8%8F-cara-membuat-slice-dari-array)
  - [🛠️ Function Utama Pada Slice](#%EF%B8%8F-function-utama-pada-slice)
  - [⚠️ Hati-Hati: Perbedaan Array vs Slice & Cara Kerja `append()`](#%EF%B8%8F-hati-hati-perbedaan-array-vs-slice--cara-kerja-append)

---

## 🎯 Main Function

Go-Lang mirip dengan bahasa pemrograman C/C++, di mana ia memerlukan fungsi `main` sebagai pintu masuk (*entry point*) eksekusi program.

- Fungsi `main` adalah fungsi utama yang akan dieksekusi secara otomatis saat program dijalankan.
- Untuk membuat fungsi di Golang, kita menggunakan kata kunci `func`.
- Fungsi `main` **wajib** berada di dalam `package main`.
- Titik koma (`;`) di Golang **tidak wajib**, sehingga kita bebas menambahkan atau tidak menambahkan titik koma di akhir baris kode.

---

## 🚨 Multiple Main Function

Di dalam satu *package* atau modul Golang, fungsi-fungsi yang ada harus bersifat unik (nama tidak boleh sama).

- Jika kita membuat file baru (misalnya `sample.go`) dan membuat fungsi dengan nama yang sama yaitu `main`, program **tidak akan bisa di-build**.
- Hal ini terjadi karena fungsi `main` baru tersebut dianggap sebagai duplikat dari fungsi `main` yang sudah ada di file lain (misalnya `helloworld.go` / `main.go`).

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

## 🏷️ Variable

- **Variable** adalah wadah/tempat untuk menyimpan data.
- Variable digunakan agar kita bisa mengakses dan mengolah data yang sama di mana pun kita mau.
- Di Go-Lang, satu variable **hanya bisa menyimpan tipe data yang konsisten**. Jika ingin menyimpan data bertipe lain, kita harus membuat variable terpisah.
- Untuk membuat variable, kita bisa menggunakan kata kunci `var` diikuti nama variable dan tipe datanya, atau menggunakan penulisan singkat `:=`.

---

## 🔐 Constant

- **Constant** adalah jenis variable yang nilainya **konstan (tidak bisa diubah lagi)** setelah pertama kali dideklarasikan.
- Pembuatannya mirip dengan variable, tetapi menggunakan kata kunci `const`.
- Saat pembuatan constant, nilai datanya **wajib langsung diinisialisasi**.

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

## 📛 Type Declarations

- **Type Declarations** adalah fitur di Golang untuk membuat tipe data baru dari tipe data yang sudah ada.
- Fitur ini biasanya digunakan untuk membuat **alias** terhadap tipe data dasar dengan nama yang lebih spesifik dan kontekstual, sehingga kode program lebih mudah dipahami dan lebih aman (*type-safe*).
- **Contoh:**
  ```go
  type NoKTP string
  type Married bool

  var ktpPengguna NoKTP = "1234567890"
  var statusNikah Married = true
  ```

---

## 🧮 Mathematics & Operators

### 🔧 Tabel Operasi Matematika Dasar

| Operasi Matematika | Operator | Keterangan |
| :--- | :--- | :--- |
| Penjumlahan | `+` | Menjumlahkan dua nilai. |
| Pengurangan | `-` | Mengurangi dua nilai. |
| Perkalian | `*` | Mengalikan dua nilai. |
| Pembagian | `/` | Membagi dua nilai. |
| Sisa Pembagian | `%` | Menghasilkan sisa (*remainder/modulo*) dari hasil pembagian. *(Hanya untuk Integer)* |

### ➕ Augmented Assignment (Operasi Singkat)

Golang menyediakan operator singkat yang menggabungkan operasi matematika langsung dengan penugasan variabel (`=`).

| Syntax Singkat | Setara Dengan |
| :--- | :--- |
| `x += y` | `x = x + y` |
| `x -= y` | `x = x - y` |
| `x *= y` | `x = x * y` |
| `x /= y` | `x = x / y` |
| `x %= y` | `x = x % y` |

### ⬆️ Tabel Unary Operator

| Operator | Keterangan |
| :--- | :--- |
| `++` | `a = a + 1` (Penambahan 1) |
| `--` | `a = a - 1` (Pengurangan 1) |
| `-` | Negative (Negatif angka) |
| `+` | Positive (Positif angka) |
| `!` | Boolean Kebalikan (NOT) |

### ⚖️ Operator Perbandingan

| Operator | Keterangan |
| :--- | :--- |
| `>` | Lebih Dari |
| `<` | Kurang Dari |
| `>=` | Lebih Dari Sama Dengan |
| `<=` | Kurang Dari Sama Dengan |
| `==` | Sama Dengan |
| `!=` | Tidak Sama Dengan |

### 🔀 Operator Boolean

| Operator | Keterangan |
| :--- | :--- |
| `&&` | DAN (Keduanya harus `true`) |
| `\|\|` | ATAU (Salah satu `true`) |
| `!` | Kebalikan (NOT) |

---

## 🧱 Tipe Data Array

- **Array** adalah tipe data yang berisikan kumpulan data dengan tipe data yang sama.  
- Ukuran (daya tampung) Array **bersifat tetap (fixed)** dan wajib ditentukan saat pertama kali dibuat. Ukuran Array tidak bisa bertambah atau berkurang setelah dibuat.

### 🔧 Function & Operasi Array

| Operasi | Keterangan |
| :--- | :--- |
| `len(array)` | Mendapatkan panjang (jumlah slot) Array. |
| `array[index]` | Mengambil data pada posisi index tertentu (mulai dari 0). |
| `array[index] = value` | Mengubah/mengisi data pada posisi index tertentu. |

---

## 🍕 Tipe Data Slice

- **Slice** adalah potongan (*slice*) dari data Array.
- **Slice** sangat mirip dengan Array, tetapi ukurannya **dinamis (bisa bertambah/berkurang)**.
- Slice dan Array **selalu terhubung**: Slice sebenarnya tidak menyimpan data sendiri, melainkan hanya "mengintip" atau mengakses Array yang ada di bawahnya (*Underlying Array*).

### 🔍 3 Komponen Utama Tipe Data Slice

1. **Pointer**: Penunjuk ke elemen pertama di Array yang diakses oleh Slice.
2. **Length (`len`)**: Jumlah elemen yang saat ini ada di dalam Slice.
3. **Capacity (`cap`)**: Kapasitas maksimal Slice (dihitung dari index pointer sampai ke ujung Array paling akhir). Nilai `len` **tidak boleh** lebih besar dari `cap`.

### ✂️ Cara Membuat Slice dari Array

| Syntax | Keterangan |
| :--- | :--- |
| `array[low:high]` | Membuat slice dari index `low` sampai sebelum index `high`. |
| `array[low:]` | Membuat slice dari index `low` sampai index paling akhir di Array. |
| `array[:high]` | Membuat slice dari index `0` sampai sebelum index `high`. |
| `array[:]` | Membuat slice yang mengambil seluruh isi Array. |

### 🛠️ Function Utama Pada Slice

| Function | Keterangan |
| :--- | :--- |
| `len(slice)` | Mendapatkan panjang (`length`) dari Slice. |
| `cap(slice)` | Mendapatkan kapasitas (`capacity`) dari Slice. |
| `append(slice, data)` | Menambahkan data baru ke posisi akhir Slice. *(Hanya bisa menerima Slice, tidak bisa Array!)* |
| `make([]TipeData, length, capacity)` | Membuat Slice baru dari awal tanpa perlu mendeklarasikan Array manual. |
| `copy(destination, source)` | Menyalin elemen dari Slice `source` ke `destination`. *(Pastikan length `destination` disiapkan dulu!)* |

---

## ⚠️ Hati-Hati: Perbedaan Array vs Slice & Cara Kerja `append()`

### 🚨 1. Cara Membedakan Penulisan Kode (Array vs Slice)

Sering kali programmer pemula terkecoh antara sintaks Array dan Slice. Perhatikan bedanya:

```go
// 🧱 INI ADALAH ARRAY (Ukurannya Tetap & Mati)
var arr1 [5]int                    // Dideklarasikan angka pasti [5]
var arr2 = [...]string{"A", "B"}    // Menggunakan titik tiga [...] (dihitung otomatis saat compile)
// ❌ append(arr1, 10) -> ERROR! Array tidak bisa di-append!

// 🍕 INI ADALAH SLICE (Ukurannya Dinamis)
var slice1 []int                   // Kurung siku KOSONG []
var slice2 = []string{"A", "B"}    // Kurung siku KOSONG []
// ✅ append(slice1, 10) -> BISA! Slice mendukung append.
```

---

### ⚡ 2. Rahasia Cara Kerja `append()` & Kapasitas (`cap`)

Mengapa kadang `append()` merubah Array asli dan kadang TIDAK merubah Array asli?

- **Kondisi 1: Kapasitas Masih Cukup (`len < cap`)**  
  `append()` akan **menumpang** di Array asal yang sama. Data baru ditaruh di slot kosong berikutnya.  
  👉 **Efek:** Array / Slice awal **IKUT BERUBAH** nilainya!

- **Kondisi 2: Kapasitas Sudah Penuh (`len == cap`)**  
  Array asal sudah tidak muat menampung data baru. Golang secara otomatis membuat **ARRAY BARU** di memori terpisah yang ukurannya dilipatgandakan ($2 \times \text{kapasitas lama}$), lalu menyalin data lama + data baru ke Array baru tersebut.  
  👉 **Efek:** Array / Slice awal **TIDAK AKAN BERUBAH**, karena Slice hasil `append` sekarang menunjuk ke Array Baru yang berbeda!