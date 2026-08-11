package main

import (
	"log"
	"net/http"

	"github.com/zstoimchev/anti-abuse/api/internal/account"
)

func main() {
	mux := http.NewServeMux()

	accountStore := account.NewStore()
	accountHandler := account.NewHandler(accountStore)

	accountHandler.RegisterRoutes(mux)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Println("Anti-Abuse API listening on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
