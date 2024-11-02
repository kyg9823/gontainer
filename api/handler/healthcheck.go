package handler

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status bool `json:"status"`
}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	result := &Health{
		Status: true,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
