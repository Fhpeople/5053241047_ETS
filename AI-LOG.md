# AI-LOG

Catat setiap penggunaan AI yang menghasilkan kode. Jika tidak memakai AI sama sekali, tulis satu baris: **tidak memakai AI**.

Log yang jujur tidak mengurangi nilai, sebanyak apa pun AI dipakai. Yang mengurangi nilai adalah log yang tidak sesuai dengan kode, atau isi log yang tidak bisa dijelaskan.

## Kelola Sumber Daya (B1)
- Prompt: Minta bimbingan dan penjelasan langkah demi langkah untuk membuat endpoint POST /sumber-daya dan GET /sumber-daya di Go dengan net/http beserta pembagian filenya.
- Dipakai: Fungsi handleCreateSumberDaya, handleGetSumberDaya, dan handleGetSumberDayaByID pada file handler_sumberdaya.go.

## Buat dan Lihat Peminjaman (B2 & B3)
- Prompt: Minta dibuatkan fungsi handler untuk POST /peminjaman dan GET /peminjaman dengan filter query URL opsional, serta penjelasan logika penanganan tanggalnya.
- Dipakai: Fungsi handleCreatePeminjaman, handleGetPeminjaman, dan handleGetPeminjamanByID pada file handler_peminjaman.go.

## Validasi Input dan Error JSON (B4)
- Prompt: Minta arahan cara melakukan validasi format tanggal YYYY-MM-DD dan validasi rentang slot 1-8 di Go.
- Dipakai: Penggunaan time.Parse("2006-01-02", tanggal) dan fungsi helper sendError di file helper.go.