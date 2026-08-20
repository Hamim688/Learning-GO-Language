# 🔒 Closure

---

## 🎯 Konsep Closure

- **Closure** adalah kemampuan sebuah function untuk mengakses dan berinteraksi dengan data-data (variabel) di sekitarnya dalam scope yang sama (*outer/parent scope*).
- Ketika function dibuat, function tersebut secara otomatis mengikat (*bind*) variabel-variabel yang ada di scope tempat function tersebut dideklarasikan.

---

## ⚠️ Penggunaan Closure

- Variabel yang ada di scope luar dapat secara tidak sengaja terubah (*mutate*) oleh logika di dalam closure function.
- Harap gunakan fitur closure ini dengan bijak saat membuat aplikasi agar tidak menyebabkan efek samping (*side effect*) yang tidak diinginkan pada variabel di scope luar.

---

[⬅️ Sebelumnya: Function](12-function.md) | [🏠 Hub Utama](../../LEARNING.md)
