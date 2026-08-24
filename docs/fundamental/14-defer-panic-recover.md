# 🛡️ Defer, Panic & Recover

---

## ⏱️ Defer Function

- **Defer Function** adalah function yang dapat kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi.
- Defer function akan **selalu dieksekusi** di akhir function, walaupun terjadi error pada function yang sedang dieksekusi tersebut.

---

## 🚨 Panic Function

- **Panic Function** adalah function yang bisa kita gunakan untuk menghentikan program.
- Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan.
- Saat *panic function* dipanggil, program akan terhenti, namun **defer function** tetap akan dieksekusi.

---

## 🩹 Recover Function

- **Recover** adalah function yang bisa kita gunakan untuk menangkap data panic.
- Dengan *recover*, proses panic akan terhenti sehingga program akan tetap berjalan.
- Fungsi *recover* wajib dipanggil di dalam **defer function** agar dapat menangkap pesan panic dengan benar.

---

[⬅️ Sebelumnya: Closure](13-closure.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Struct & Struct Method ➡️](15-struct.md)
