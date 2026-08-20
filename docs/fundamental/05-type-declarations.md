# 📛 Type Declarations

---

## 📛 Type Declarations

- **Type Declarations** adalah fitur di Golang untuk membuat tipe data baru dari tipe data yang sudah ada.
- Fitur ini biasanya digunakan untuk membuat **alias** terhadap tipe data dasar dengan nama yang lebih spesifik dan kontekstual, sehingga kode program lebih mudah dipahami dan lebih aman (*type-safe*).
- **Contoh:**
  ```go
  type NoKTP string
  type Married bool

  var ktpPengguna NoKTP = "1234567890"
  var statusNikah Married = true
  ```

---

[⬅️ Sebelumnya: Konversi Tipe Data](04-konversi-tipe-data.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Matematika & Operator ➡️](06-matematika-dan-operator.md)
