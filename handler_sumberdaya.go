package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (s *Store) handleCreateSumberDaya(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nama string `json:"nama"`
		Jenis string `json:"jenis"`
		Stok int `json:"stok"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "Body request tidak valid")
		return
	}

	req.Nama = strings.TrimSpace(req.Nama)
	req.Jenis = strings.TrimSpace(req.Jenis)

	if req.Nama == "" || len(req.Nama) > 50 {
		sendError(w, http.StatusBadRequest, "Nama wajib diisi maksimal 50 karakter")
		return
	}

	if req.Jenis != "ruang" && req.Jenis != "alat" {
		sendError(w, http.StatusBadRequest, "Jenis harus 'ruang' atau 'alat'")
		return
	}

	if req.Stok < 1 {
		sendError(w, http.StatusBadRequest, "Stok harus lebih dari 0")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	newSumberDaya := SumberDaya{
		ID: s.nextSumberDayaID,
		Nama: req.Nama,
		Jenis: req.Jenis,
		Stok: req.Stok,
	}

	s.sumberDaya[s.nextSumberDayaID] = newSumberDaya
	s.nextSumberDayaID++

	sendJSON(w, http.StatusCreated, newSumberDaya)
}

func (s *Store) handleGetSumberDaya(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]SumberDaya, 0)
	for i := 1; i < s.nextSumberDayaID; i++ {
		if val, exists := s.sumberDaya[i]; exists {
			result = append(result, val)
		}
	}

	sendJSON(w, http.StatusOK, result)
}

func (s *Store) handleGetSumberDayaByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	val, exists := s.sumberDaya[id]
	if !exists {
		sendError(w, http.StatusNotFound, "Sumber daya tidak ditemukan")
		return
	}

	sendJSON(w, http.StatusOK, val)
}