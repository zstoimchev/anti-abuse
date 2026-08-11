package account

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
	mux.HandleFunc("POST /v1/accounts", h.create)
	mux.HandleFunc("GET /v1/accounts/{id}", h.get)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var account Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if account.ID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := h.store.Create(account); err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			http.Error(w, "account already exists", http.StatusConflict)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(account)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	account, err := h.store.Get(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			http.Error(w, "account not found", http.StatusNotFound)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(account)
}
