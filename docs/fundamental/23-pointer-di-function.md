# 📍 Pointer di Function

🔗 **Source Code terkait:** [pointer_functions.go](../../fundamental/pointer_functions.go)

---

## 🎯 Konsep Pointer di Function

- Saat kita membuat parameter di function, secara default parameternya adalah **pass by value**, artinya data akan di-*copy* lalu dikirim ke function tersebut.
- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
- Hal ini membuat variabel menjadi aman, karena tidak akan bisa diubah.

---

## ⚙️ Menggunakan Parameter Pointer

- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut.
- Untuk melakukan ini, kita bisa menggunakan **pointer** di function.
- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator **`*`** di parameternya (dan mengirimkan alamat memori menggunakan operator **`&`**).

---

[⬅️ Sebelumnya: Operator new](22-operator-new.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Pointer di Method ➡️](24-pointer-di-method.md)
