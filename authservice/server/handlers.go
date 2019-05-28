package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type healthResponse struct {
	AuthServiceStatus string
}

func apiHealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("auth health check hit")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&healthResponse{AuthServiceStatus: "HEALTHY"})
}

func apiLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("auth login check hit")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := struct {
		Login string
	}{
		"successful",
	}
	json.NewEncoder(w).Encode(&resp)
}
