# ⚙️ Pointer di Method

🔗 **Source Code terkait:** [pointer_method.go](../../fundamental/pointer_method.go)

---

## 🎯 Konsep Pointer di Method

- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah **pass by value**.
- Sehingga perubahan data struct di dalam method tanpa pointer tidak akan mengubah data struct yang aslinya.

---

## 🚀 Keuntungan Menggunakan Pointer di Method

- Sangat direkomendasikan menggunakan **pointer di method**, sehingga:
  1. Method bisa mengubah data asli struct tersebut.
  2. **Tidak boros memory** karena tidak harus selalu diduplikasi data struct-nya ketika memanggil method.

---

[⬅️ Sebelumnya: Pointer di Function](23-pointer-di-function.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Access Modifier ➡️](25-access-modifier.md)
