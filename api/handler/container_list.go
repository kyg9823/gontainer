package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kyg9823/gontainer/service"
)

type ContainerFilter struct {
}

// @Summary Get Containers
// @Description Get Containers
// @Accept json
// @Produce json
// @Success 200 {object} []ContainerInfo
// @Router /gontainer/api/v1/containers [get]
func ContainerListHandler(w http.ResponseWriter, r *http.Request) {
	result, err := service.GetContainerList()
	if err != nil {
		log.Fatalf("Failed to list containers: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
