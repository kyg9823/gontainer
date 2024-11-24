package handler

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status bool `json:"status"`
}

// @Summary Get healthcheck
// @Description Get healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} Health
// @Router /gontainer/api/v1/healthcheck [get]
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	result := &Health{
		Status: true,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
