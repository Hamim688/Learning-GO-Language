# 🎯 Main Function & Multiple Main Function

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

[⬅️ Kembali ke Indeks Fundamental](README.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Tipe Data & String ➡️](02-tipe-data-dan-string.md)
