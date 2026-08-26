# 🗂️ Interface Kosong (`any`)

---

## 🎯 Konsep Interface Kosong

- Go-Lang bukanlah bahasa pemrograman yang berorientasi objek (*Object-Oriented Programming*).
- Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut (contoh di Java ada `java.lang.Object`).
- Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan **interface kosong** (`interface{}`).
- Interface kosong adalah interface yang tidak memiliki deklarasi method satupun. Hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya.
- Interface kosong juga memiliki type alias bernama **`any`**.

---

## 💡 Penggunaan Interface Kosong di Go-Lang

- Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti:
  - `fmt.Println(a ...interface{})` atau `fmt.Println(a ...any)`
  - `panic(v interface{})`
  - `recover() interface{}`
  - Dan lain-lain.

---

[⬅️ Sebelumnya: Interface](16-interface.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Nil ➡️](18-nil.md)
