# 📍 Pointer & Pass by Reference

🔗 **Source Code terkait:** [pointer.go](../../fundamental/pointer.go)

---

## 1. 📦 Pass by Value

- Secara default di Go-Lang semua variable itu di *passing by value*, bukan *by reference*.
- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah **duplikasi value-nya**.

```text
┌───────────────────────────────────────────────────────────┐
│ Memory                                                    │
│                                                           │
│  ┌──────────────────────┐   copied   ┌──────────────────┐ │
│  │ Address1             │ ─────────► │ Address2         │ │
│  │ -------------------- │            │ ---------------- │ │
│  │ City = Subang        │            │ City = Subang    │ │
│  │ Province = Jawa B.   │            │ Province = Jawa B│ │
│  │ Country = Indonesia  │            │ Country = Ind.   │ │
│  └──────────────────────┘            └──────────────────┘ │
└──────────────▲─────────────────────────────────▲──────────┘
               │                                 │
          ┌────┴─────┐                      ┌────┴─────┐
          │ address1 │                      │ address2 │
          └──────────┘                      └──────────┘
```

---

## 2. 📍 Pointer (Pass by Reference)

- **Pointer** adalah kemampuan membuat *reference* ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada.
- Sederhananya, dengan kemampuan pointer, kita bisa membuat **pass by reference**.
- Menggunakan operator `&` di depan variabel untuk mengambil alamat memorinya.

```text
┌───────────────────────────────────────────────────────────┐
│ Memory                                                    │
│                                                           │
│           ┌──────────────────────────────────┐            │
│           │ Address1                         │            │
│           │ -------------------------------- │            │
│           │ City = Subang                    │            │
│           │ Province = Jawa Barat            │            │
│           │ Country = Indonesia              │            │
│           └──────────────────────────────────┘            │
└─────────────────────────────▲─────────────────────────────┘
                              │
               ┌──────────────┴──────────────┐
               │                             │
          ┌────┴─────┐               ┌───────┴──────┐
          │ address1 │               │   address3   │
          └──────────┘               │  (&address1) │
                                     └──────────────┘
```

---

[⬅️ Sebelumnya: Type Assertions](19-type-assertions.md) | [🏠 Hub Utama](../../LEARNING.md) | [📘 Daftar Isi Fundamental](README.md)
