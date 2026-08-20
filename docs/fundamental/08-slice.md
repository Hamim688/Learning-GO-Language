# 🍕 Tipe Data Slice

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

---

[⬅️ Sebelumnya: Tipe Data Array](07-array.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: If Expression ➡️](09-if-expression.md)
