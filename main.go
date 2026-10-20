package main

import (
	"net/http"
)

func NewHandler() http.Handler {
	store := NewStore()
	mux := http.NewServeMux()

	mux.HandleFunc("POST /sumber-daya", store.handleCreateSumberDaya)
	mux.HandleFunc("GET /sumber-daya", store.handleGetSumberDaya)
	mux.HandleFunc("GET /sumber-daya/{id}", store.handleGetSumberDayaByID)

	mux.HandleFunc("POST /peminjaman", store.handleCreatePeminjaman)
	mux.HandleFunc("GET /peminjaman", store.handleGetPeminjaman)
	mux.HandleFunc("GET /peminjaman/{id}", store.handleGetPeminjamanByID)

	return mux
}

func main() {
	http.ListenAndServe(":8080", NewHandler())
}