# 🛠️ Membuat Custom Error

🔗 **Source Code terkait:** [custom_error.go](../../fundamental/custom_error.go)

---

## 🎯 Konsep Custom Error

- Karena error adalah sebuah interface, jadi jika kita ingin membuat error sendiri, kita bisa membuat struct yang mengikuti kontrak dari interface error.
- Cukup dengan menambahkan method `Error() string` pada struct yang kita buat, maka struct tersebut otomatis dianggap sebagai implementasi dari interface error.

---

## 🔍 Pengecekan Custom Error

- Dengan membuat custom error, kita bisa menambahkan informasi/field detail tambahan di dalam struct error tersebut.
- Untuk mengecek jenis custom error mana yang terjadi saat program berjalan, kita bisa menggunakan **Type Assertions** atau **Type Switch**.

---

[⬅️ Sebelumnya: Error Interface](27-error-interface.md) | [🏠 Hub Utama](../../LEARNING.md)
