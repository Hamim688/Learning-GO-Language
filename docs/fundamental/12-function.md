# 🛠️ Function

---

## 🎯 Deklarasi Function

- **Function** adalah blok kode yang sengaja dibuat agar dapat digunakan berulang-ulang (*reusable*).
- Membuat function menggunakan kata kunci `func` diikuti dengan nama fungsi dan kurung `()`.
- Untuk mengeksekusi function, kita memanggil nama function tersebut diikuti kurung buka dan tutup `()`.

---

## 📥 Function Parameter

- Parameter digunakan saat function membutuhkan masukan data dari luar.
- Parameter ditulis di dalam tanda kurung `()` setelah nama function. Kita dapat membuat lebih dari satu parameter.
- Parameter bersifat opsional. Namun jika function memiliki parameter, kita **wajib** mengirimkan argumen sesuai jumlah dan tipe datanya saat memanggil function tersebut.

---

## 📤 Function Return Value

- Function dapat mengembalikan data (*return value*).
- Jika function mengembalikan data, kita wajib mendefinisikan tipe data kembaliannya setelah tanda kurung parameter.
- Kata kunci `return` digunakan di dalam function untuk mengembalikan nilai data tersebut.

---

## 🔄 Returning Multiple Values & Menghiraukan Value

- Di Golang, function dapat mengembalikan lebih dari satu nilai sekaligus (*multiple return values*).
- Semua tipe data kembalian wajib ditulis di dalam kurung `()` setelah parameter.
- Saat dipanggil, semua nilai kembalian **wajib ditangkap**. Jika ada nilai yang ingin dihiraukan, kita wajib menggunakan tanda *underscore* (`_`).

---

## 🏷️ Named Return Values

- Kita juga bisa memberikan nama variabel secara langsung pada tipe data return value di deklarasi function.
- Variabel return tersebut otomatis dibuat dan cukup dikembalikan menggunakan kata kunci `return` tanpa menyebutkan variabelnya lagi (*naked return*).

---

## 🎲 Variadic Function & Slice Parameter

- **Variadic Function** adalah function dengan parameter posisi terakhir yang dapat menerima variabel argumen (*varargs*) berupa banyak data sekaligus.
- Berbeda dengan parameter Array biasa yang mengharuskan kita membuat Array terlebih dahulu, varargs memungkinkan kita mengirim data langsung dipisahkan dengan tanda koma.
- Jika kita sudah memiliki variabel berupa **Slice**, kita bisa mengirimkannya ke varargs parameter menggunakan sintaks variadic `slice...`.

---

## 💡 Function Value

- Di Golang, **Function** adalah *first class citizen*.
- Artinya, function juga dianggap sebagai tipe data, sehingga function dapat disimpan ke dalam variabel seperti halnya nilai biasa.

---

## 🔀 Function as Parameter

- Karena function merupakan tipe data, maka function tidak hanya bisa disimpan dalam variabel, tetapi juga dapat dijadikan **parameter** untuk function lainnya.
- Fitur ini sangat berguna untuk menerapkan konsep callback atau *filtering* pada logika program.

---

## 📛 Function Type Declaration

- Jika sintaks deklarasi function sebagai parameter terlalu panjang, penulisan kode bisa menjadi rumit.
- Kita dapat memanfaatkan **Type Declaration** untuk membuat nama alias bagi tipe deklarasi function tersebut, sehingga kode menjadi lebih rapi dan mudah dibaca.

---

## 🥸 Anonymous Function

- Sebelumnya setiap membuat function, kita selalu memberikan nama pada function tersebut.
- Namun kadang ada kalanya lebih mudah membuat function secara langsung di variabel atau parameter tanpa harus mendeklarasikan nama function terlebih dahulu.
- Hal tersebut dinamakan **anonymous function** (function tanpa nama).

---

## 🔄 Recursive Function

- **Recursive Function** adalah function yang memanggil dirinya sendiri.
- Kadang dalam pekerjaan, kita sering menemui kasus di mana menggunakan *recursive function* lebih mudah dibandingkan tidak menggunakannya.
- Contoh kasus yang lebih mudah diselesaikan menggunakan *recursive function* adalah perhitungan **Factorial**.

---

[⬅️ Sebelumnya: For Loops](11-for-loops.md) | [🏠 Hub Utama](../../LEARNING.md) | [Selanjutnya: Closure ➡️](13-closure.md)
