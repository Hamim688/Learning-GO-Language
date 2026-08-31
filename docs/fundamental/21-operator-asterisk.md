# ✳️ Operator Asterisk (`*`)

🔗 **Source Code terkait:** [asterisk.go](../../fundamental/asterisk.go)

---

## 🎯 Konsep Operator Asterisk (`*`)

- Saat kita mengubah variabel pointer tanpa operator `*` (misalnya `people2 = &People{...}`), yang berubah **hanya variabel pointer itu sendiri** karena dialihkan menunjuk ke lokasi memori baru.
- Semua variabel lain yang mengacu ke data lama tidak akan berubah.
- Namun jika kita ingin **mengubah seluruh variabel yang mengacu ke data tersebut** (menimpa nilai di alamat memori yang sama), kita dapat menggunakan **operator `*`** (*dereferencing*).

---

## 1. ❌ Tanpa Operator `*` (Hanya Mengalihkan Pointer)

- Saat kita memberikan nilai baru menggunakan `&` ke variabel pointer (seperti `people2 = &People{...}`), variabel `people2` dialihkan ke **alamat memori baru**.
- **Efeknya:** Variabel awal (`people1`) **TIDAK BERUBAH** sama sekali karena ia masih menunjuk ke alamat memori lama.

```text
Kondisi Awal: people2 := &people1
┌───────────────────────────────────────────────────────────┐
│ Alamat Memory: 0x001                                      │
│  ┌─────────────────────────────────────────────────────┐  │
│  │ Name = "Jang Won Young", Address = "Wonogiri", 22   │  │
│  └──────────────────────────▲──────────────────────────┘  │
└─────────────────────────────┼─────────────────────────────┘
               ┌──────────────┴──────────────┐
               │                             │
          ┌────┴─────┐               ┌───────┴──────┐
          │ people1  │               │   people2    │ (menunjuk ke 0x001)
          └──────────┘               └──────────────┘

Setelah: people2 = &People{"Karina", "Jakarta", 20}  (TANPA OPERATOR *)
┌───────────────────────────────────────────────────────────┐
│ Alamat Memory: 0x001 (Tetap Data Lama)                    │
│  ┌─────────────────────────────────────────────────────┐  │
│  │ Name = "Jang Won Young", Address = "Wonogiri", 22   │  │
│  └──────────────────────────▲──────────────────────────┘  │
│                                                           │
│ Alamat Memory Baru: 0x002                                 │
│  ┌─────────────────────────────────────────────────────┐  │
│  │ Name = "Karina", Address = "Jakarta", 20            │  │
│  └──────────────────────────▲──────────────────────────┘  │
└─────────────────────────────┼─────────────────────────────┘
                              │
          ┌──────────┐        │      ┌──────────────┐
          │ people1  ├────────┘      │   people2    │ (sekarang menunjuk 0x002)
          └──────────┘               └──────────────┘

👉 HASIL: 
- people1 : Jang Won Young (Tetap di Memory 0x001)
- people2 : Karina         (Pindah ke Memory Baru 0x002)
```

---

## 2. ✅ Menggunakan Operator `*` (Menimpa Memory Asli)

- Saat kita menggunakan operator `*` di depan variabel pointer (seperti `*people2 = People{...}`), kita melakukan *dereferencing*.
- Artinya, kita **memaksa menimpa nilai data di alamat memori asli** (0x001) yang sedang ditunjuk oleh pointer tersebut.
- **Efeknya:** Semua variabel yang menunjuk ke alamat memori tersebut (termasuk `people1`) **IKUT BERUBAH** secara otomatis.

```text
Setelah: *people2 = People{"Pharita", "Bandung", 20}  (MENGGUNAKAN OPERATOR *)
┌───────────────────────────────────────────────────────────┐
│ Alamat Memory: 0x001 (Nilai Data Ditimpa Langsung!)      │
│  ┌─────────────────────────────────────────────────────┐  │
│  │ Name = "Pharita", Address = "Bandung", 20           │  │
│  └──────────────────────────▲──────────────────────────┘  │
└─────────────────────────────┼─────────────────────────────┘
               ┌──────────────┴──────────────┐
               │                             │
          ┌────┴─────┐               ┌───────┴──────┐
          │ people1  │               │   people2    │ (tetap menunjuk 0x001)
          └──────────┘               └──────────────┘

👉 HASIL: 
- people1 : Pharita (Ikut Berubah karena data di 0x001 ditimpa!)
- people2 : Pharita (Ikut Berubah karena data di 0x001 ditimpa!)
```

---

[⬅️ Sebelumnya: Pointer](20-pointer.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Operator new ➡️](22-operator-new.md)
