<div align="center">

# 🚀 Panduan Instalasi Golang

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey?style=for-the-badge)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](#)

*Panduan lengkap, cepat, dan mudah untuk mengatur environment Go (Golang) di sistem Anda.*

</div>

---

## 🌟 Daftar Isi
- [🛠️ Persiapan](#-persiapan)
- [🐧 Instalasi di Linux](#-instalasi-di-linux)
- [🪟 Instalasi di Windows](#-instalasi-di-windows)
- [🍎 Instalasi di macOS](#-instalasi-di-macos)
- [✅ Verifikasi Instalasi](#-verifikasi-instalasi)
- [🎉 Membuat Proyek Pertama](#-membuat-proyek-pertama)

---

## 🛠️ Persiapan

Sebelum memulai, pastikan Anda:
1. Memiliki **koneksi internet** yang stabil.
2. Memiliki akses ke **Terminal** (Linux/macOS) atau **Command Prompt / PowerShell** (Windows).

---

## 🐧 Instalasi di Linux

> **Catatan:** Langkah-langkah ini menggunakan terminal. Pastikan Anda memiliki hak akses `sudo`.

### 1. Hapus versi Go sebelumnya (jika ada)
```bash
sudo rm -rf /usr/local/go
```

### 2. Unduh dan ekstrak Go
Ganti `go1.21.0` dengan [versi terbaru](https://go.dev/dl/).
```bash
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
```

### 3. Tambahkan Go ke PATH
Tambahkan baris berikut ke file `~/.profile` atau `~/.bashrc` Anda:
```bash
export PATH=$PATH:/usr/local/go/bin
```
Terapkan perubahan dengan menjalankan:
```bash
source ~/.bashrc
```

---

## 🪟 Instalasi di Windows

1. Kunjungi [halaman unduhan resmi Go](https://go.dev/dl/).
2. Unduh file instalasi berformat **MSI** (contoh: `go1.21.0.windows-amd64.msi`).
3. Buka file `.msi` yang telah diunduh dan ikuti instruksi **Wizard Instalasi**.
4. Secara default, Go akan terinstal di `C:\Program Files\Go` dan secara otomatis ditambahkan ke environment variable `PATH` Anda.

> 💡 **Tip:** Jika Terminal sudah terbuka saat Anda menginstal, Anda mungkin perlu menutup dan membukanya kembali agar perubahan `PATH` dapat dimuat.

---

## 🍎 Instalasi di macOS

1. Kunjungi [halaman unduhan resmi Go](https://go.dev/dl/).
2. Unduh paket installer berformat **.pkg** (contoh: `go1.21.0.darwin-amd64.pkg`).
3. Buka file `.pkg` yang diunduh dan ikuti petunjuk instalasinya.
4. Go akan otomatis terinstal di `/usr/local/go` dan `PATH` akan dikonfigurasi secara otomatis.

> 💡 **Alternatif menggunakan Homebrew:**
> Jika Anda menggunakan Homebrew, Anda dapat menginstal Go cukup dengan satu perintah:
> ```bash
> brew install go
> ```

---

## ✅ Verifikasi Instalasi

Setelah proses instalasi selesai, mari kita pastikan Go sudah terpasang dengan benar.

Buka terminal baru dan jalankan perintah berikut:
```bash
go version
```

**Output yang diharapkan:**
```
go version go1.21.0 linux/amd64
```
*(Versi dan OS mungkin berbeda sesuai dengan sistem Anda)*

---

## 🎉 Membuat Proyek Pertama

Mari kita buat program "Hello World" untuk memastikan environment Anda siap digunakan!

1. **Buat direktori proyek:**
   ```bash
   mkdir belajar-golang
   cd belajar-golang
   ```

2. **Inisialisasi modul Go:**
   ```bash
   go mod init belajar-golang
   ```

3. **Buat file `main.go`:**
   Buat file baru bernama `main.go` dan masukkan kode berikut:
   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Halo, Golang! 🚀")
   }
   ```

4. **Jalankan program:**
   ```bash
   go run main.go
   ```

**Selamat! Anda sudah berhasil menjalankan kode Go pertama Anda!** 💻✨

---
<div align="center">
  <i>Dibuat dengan ❤️ untuk pembelajaran Golang</i>
</div>
