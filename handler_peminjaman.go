package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (s *Store) handleCreatePeminjaman(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SumberDayaID int `json:"sumber_daya_id"`
		Peminjam     string `json:"peminjam"`
		Tanggal      string `json:"tanggal"`
		SlotMulai    int `json:"slot_mulai"`
		SlotSelesai  int `json:"slot_selesai"`
		Jumlah       int `json:"jumlah"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "Body request tidak valid")
		return
	}

	req.Peminjam = strings.TrimSpace(req.Peminjam)
	req.Tanggal = strings.TrimSpace(req.Tanggal)

	if req.Peminjam == "" || len([]rune(req.Peminjam)) > 50 {
		sendError(w, http.StatusBadRequest, "Nama peminjam wajib diisi maksimal 50 karakter")
		return
	}

	if _, err := time.Parse("2006-01-02", req.Tanggal); err != nil {
		sendError(w, http.StatusBadRequest, "Format tanggal tidak valid, gunakan YYYY-MM-DD")
		return
	}

	if req.SlotMulai < 1 || req.SlotMulai > 8 || req.SlotSelesai < 1 || req.SlotSelesai > 8 || req.SlotMulai > req.SlotSelesai {
		sendError(w, http.StatusBadRequest, "Slot mulai dan selesai harus antara 1-8 dan slot mulai tidak boleh lebih besar dari slot selesai")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	sd, exists := s.sumberDaya[req.SumberDayaID]
	if !exists {
		sendError(w, http.StatusBadRequest, "Sumber daya tidak ditemukan")
		return
	}

	if req.Jumlah < 1 || req.Jumlah > sd.Stok {
		sendError(w, http.StatusBadRequest, "Jumlah peminjaman harus lebih dari 0 dan tidak melebihi stok")
		return
	}

	if s.isBentrok(req.SumberDayaID, req.Tanggal, req.SlotMulai, req.SlotSelesai, req.Jumlah, 0) {
		sendError(w, http.StatusConflict, "Slot pada tanggal dan jam tersebut sudah dipesan (bentrok)")
		return
	}

	newPinjam := Peminjaman{
		ID:           s.nextPeminjamanID,
		SumberDayaID: req.SumberDayaID,
		Peminjam:     req.Peminjam,
		Tanggal:      req.Tanggal,
		SlotMulai:    req.SlotMulai,
		SlotSelesai:  req.SlotSelesai,
		Jumlah:       req.Jumlah,
		Status:       "diajukan",
	}

	s.peminjaman[s.nextPeminjamanID] = newPinjam
	s.nextPeminjamanID++

	sendJSON(w, http.StatusCreated, newPinjam)
}

func (s *Store) handleGetPeminjaman(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	sDayaIDStr := query.Get("sumber_daya_id")
	filterTanggal := strings.TrimSpace(query.Get("tanggal"))
	filterStatus := strings.TrimSpace(query.Get("status"))

	var filterSDayaID int
	if sDayaIDStr != "" {
		id, err := strconv.Atoi(sDayaIDStr)
		if err != nil {
			sendError(w, http.StatusBadRequest, "Filter sumber_daya_id harus berupa angka")
			return
		}
		filterSDayaID = id
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]Peminjaman, 0)

	for i := 1; i < s.nextPeminjamanID; i++ {
		p, exists := s.peminjaman[i]
		if !exists {
			continue
		}

		if sDayaIDStr != "" && p.SumberDayaID != filterSDayaID {
			continue
		}
		if filterTanggal != "" && p.Tanggal != filterTanggal {
			continue
		}
		if filterStatus != "" && p.Status != filterStatus {
			continue
		}

		result = append(result, p)
	}

	sendJSON(w, http.StatusOK, result)
}

func (s *Store) handleGetPeminjamanByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, http.StatusBadRequest, "ID harus berupa angka")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	p, exists := s.peminjaman[id]
	if !exists {
		sendError(w, http.StatusNotFound, "Peminjaman tidak ditemukan")
		return
	}

	sendJSON(w, http.StatusOK, p)
}

func (s *Store) isBentrok(sumberDayaID int, tanggal string, slotMulai, slotSelesai, jumlahBaru int, excludeID int) bool {
	sd, exists := s.sumberDaya[sumberDayaID]
	if !exists {
		return false
	}

	for slot := slotMulai; slot <= slotSelesai; slot++ {
		totalTerpakai := 0

		for _, p := range s.peminjaman {
			if excludeID > 0 && p.ID == excludeID {
				continue
			}

			if p.SumberDayaID == sumberDayaID && p.Tanggal == tanggal && (p.Status == "diajukan" || p.Status == "disetujui") {
				if slot >= p.SlotMulai && slot <= p.SlotSelesai {
					totalTerpakai += p.Jumlah
				}
			}
		}

		if totalTerpakai+jumlahBaru > sd.Stok {
			return true
		}
	}

	return false
}