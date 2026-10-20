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
Pengecekan bentrok slot dilakukan dengan cara mengiterasi setiap slot jam satu per satu (slot 1 hingga 8), bukan memeriksa bentrok berdasarkan rentang awal dan akhir saja. Pilihan lain yang dipertimbangkan adalah menggunakan rumus irisan rentang matematika. Alasan memilih pengecekan per slot adalah karena setiap sumber daya memiliki jumlah kapasitas stok unit tertentu. Rumus irisan rentang biasa hanya mendeteksi apakah ada peminjaman lain di jam yang sama, tetapi tidak bisa menghitung akumulasi total unit yang terpakai jika terdapat beberapa peminjaman sekaligus di slot jam tersebut. Pengecekan per slot memastikan total unit terpakai di setiap jam tidak melebihi stok yang ada.

## 4. Bug atau kendala (minimal dua)
### Bug 1
- Gejala: Validasi nama sumber daya dan nama peminjam gagal mendeteksi batas panjang karakter dengan tepat saat menggunakan input karakter khusus.
- Cara menemukan: Diperiksa saat penulisan logika validasi batas maksimal 50 karakter.
- Penyebab dan solusi: Penggunaan fungsi bawaan len(string) di Go menghitung jumlah byte, bukan jumlah karakter. Pada karakter tertentu, 1 karakter bisa dihitung lebih dari 1 byte. Solusinya adalah mengonversi string ke slice rune lebih dulu dengan len([]rune(text)) agar jumlah karakter dihitung dengan benar sesuai spesifikasi kontrak API.

### Bug 2
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

## 5. Change request
(Apa yang diminta, bagian kode mana yang kamu ubah, dan satu keputusan desain yang kamu ambil.)
