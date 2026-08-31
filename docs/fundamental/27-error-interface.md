# ⚠️ Error Interface & Membuat Error

🔗 **Source Code terkait:** [error.go](../../fundamental/error.go)

---

## 🎯 Konsep Error Interface

- Go-Lang memiliki interface yang digunakan sebagai pengabaian/penanganan error, namanya adalah **`error`**.
- Error merupakan interface bawaan Go-Lang, yang di dalamnya terdapat method `Error() string`.

---

## 🛠️ Membuat Error

- Untuk membuat error, kita tidak perlu manual.
- Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package **`errors`**.
- Kita bisa menggunakan `errors.New("pesan error")` untuk membuat error baru.

---

[⬅️ Sebelumnya: Package Initialization & Blank Identifier](26-package-initialization-dan-blank-identifier.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Membuat Custom Error ➡️](28-custom-error.md)
