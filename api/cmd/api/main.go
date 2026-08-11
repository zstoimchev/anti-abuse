package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Anti-Abuse API listening on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
