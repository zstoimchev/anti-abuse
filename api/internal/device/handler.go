package device

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/devices", h.create)
	mux.HandleFunc("GET /v1/devices/{id}", h.get)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var device Device

	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if device.ID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := h.store.Create(device); err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			http.Error(w, "device already exists", http.StatusConflict)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(device)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	device, err := h.store.Get(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			http.Error(w, "device not found", http.StatusNotFound)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}
