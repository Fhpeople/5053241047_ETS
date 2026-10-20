package main

import (
	"sync"
)

type SumberDaya struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Jenis string `json:"jenis"`
	Stok  int    `json:"stok"`
}

type Peminjaman struct {
	ID           int    `json:"id"`
	SumberDayaID int    `json:"sumber_daya_id"`
	Peminjam     string `json:"peminjam"`
	Tanggal      string `json:"tanggal"`
	SlotMulai    int    `json:"slot_mulai"`
	SlotSelesai  int    `json:"slot_selesai"`
	Jumlah       int    `json:"jumlah"`
	Status       string `json:"status"`
	Alasan       string `json:"alasan,omitempty"`
}

type Store struct {
	mu               sync.Mutex
	sumberDaya       map[int]SumberDaya
	peminjaman       map[int]Peminjaman
	nextSumberDayaID int
	nextPeminjamanID int
}

func NewStore() *Store {
	return &Store{
		sumberDaya:       make(map[int]SumberDaya),
		peminjaman:       make(map[int]Peminjaman),
		nextSumberDayaID: 1,
		nextPeminjamanID: 1,
	}
}