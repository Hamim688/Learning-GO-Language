# 🔀 Type Assertions

---

## 🎯 Konsep Type Assertions

- **Type Assertions** merupakan kemampuan untuk mengubah tipe data `interface{}` / `any` menjadi tipe data yang diinginkan.
- Fitur ini sering sekali digunakan ketika kita bertemu dengan data *interface kosong*.

---

## 🛑 Type Assertions Menggunakan Switch

- Saat salah menggunakan *type assertions*, maka bisa berakibat terjadi **panic** di aplikasi kita.
- Jika panic dan tidak ter-*recover*, maka otomatis program kita akan mati / *crash*.
- Agar lebih aman, sebaiknya kita menggunakan **Switch Expression** untuk melakukan *type assertions*.

---

[⬅️ Sebelumnya: Nil](18-nil.md) | [🏠 Hub Utama](../../LEARNING.md) | [➡️ Selanjutnya: Pointer](20-pointer.md)
