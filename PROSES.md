# PROSES

## 1. Persona dan masalahnya
Aplikasi ini digunakan oleh Admin Himpunan yang ingin setiap peminjaman itu mendapat persetujuan darinya terlebih dulu dan tidak sembarangan langsung terpakai. Fitur khas persona ini akan menetapkan alur peminjaman sebagai berikut:
- Permintaan peminjaman diajukan
- Admin menyetujui/menolak dengan alasan
- Status permintaan akan berubah menjadi disetujui/ditolak sesuai keputusan admin

## 2. Rancangan struktur data
Ada dua struct data yang digunakan dalam program. Yang pertama adalah 'SumberDaya' yang menyimpan barang yang dapat dipinjam seperti nama, jenis, serta stok. Selanjutnya ada 'Peminjaman' yang menyimpan data peminjaman seperti nama, tanggal, jumlah, serta status.

Selain itu, ada struct 'Store' yang menjadi wadah penyimpanan utama yang menyimpan dua struct lainnya dalam dua map, yaitu map[int]SumberDaya dan map[int]Peminjaman. Alasan pemilihan map adalah untuk memudahkan pencarian dengan ID.

Struct 'Store' sendiri bertujuan untuk mencegah kondisi race condition dan menjaga keterpusatan data saat tiap request memiliki jalurnya sendiri-sendiri. Dengan menyimpan data map dan kunci sync.Mutex dalam satu tempat, tidak akan terjadi data yang tertukar atau kunci yang hilang. Di sini juga terdapat penggunaan counter ID untuk memenuhi aturan increment pada soal. Dengan menaruh counter di wadah yang sama, penambahan ID juga akan dilakukan dengan perlindungan kunci sehingga tidak akan ada urutan ganda.

## 3. Satu keputusan desain
(Satu keputusan yang kamu ambil, pilihan lain yang kamu pertimbangkan, dan alasanmu memilih yang ini. Contoh: kenapa cek bentrok per slot, bukan per rentang.)

## 4. Bug atau kendala (minimal dua)
### Bug 1
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

### Bug 2
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

## 5. Change request
(Apa yang diminta, bagian kode mana yang kamu ubah, dan satu keputusan desain yang kamu ambil.)
