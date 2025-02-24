package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kyg9823/gontainer/service"
)

// @Summary Get Images
// @Description Get Images
// @Accept json
// @Produce json
// @Router /gontainer/api/v1/images [get]
func ImageListHandler(w http.ResponseWriter, r *http.Request) {
	result, err := service.GetImageList()
	if err != nil {
		log.Fatalf("Failed to list images: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
