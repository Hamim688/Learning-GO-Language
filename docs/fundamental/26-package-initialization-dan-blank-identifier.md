# 🚀 Package Initialization & Blank Identifier

🔗 **Source Code terkait:** [database.go](../../fundamental/initialization/database.go) | [internal.go](../../fundamental/blank_identifier/internal.go)

---

## 🎯 Konsep Package Initialization

- Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses.
- Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database.
- Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama **`init`**.

---

## 🕵️ Konsep Blank Identifier pada Import

- Kadang kita hanya ingin menjalankan `init` function di package tanpa harus mengeksekusi salah satu function yang ada di package.
- Secara default, Go-Lang akan komplain (error) ketika ada package yang di-import namun tidak digunakan.
- Untuk menangani hal tersebut, kita bisa menggunakan **blank identifier (`_`)** sebelum nama package ketika melakukan import.

---

[⬅️ Sebelumnya: Access Modifier](25-access-modifier.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Error Interface ➡️](27-error-interface.md)
