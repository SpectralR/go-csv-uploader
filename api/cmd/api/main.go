package main

import (
	"encoding/json"
	"go-async-training-api/internal/job"
	"log"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/upload", job.Upload)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
	})
	mux.HandleFunc("/status", job.Status)


	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}