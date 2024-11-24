package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

type ContainerInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type ContainerFilter struct {
}

// @Summary Get Containers
// @Description Get Containers
// @Accept json
// @Produce json
// @Success 200 {object} ContainerInfo
// @Router /gontainer/api/v1/containers [get]
func ContainerListHandler(w http.ResponseWriter, r *http.Request) {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		log.Fatalf("Failed to create containerd client: %v", err)
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), "default")

	containers, err := client.Containers(ctx)
	if err != nil {
		log.Fatalf("Failed to list containers: %v", err)
	}

	fmt.Printf("Found %d containers\n", len(containers))

	result := []ContainerInfo{}
	for _, container := range containers {
		info, err := container.Info(ctx)
		if err != nil {
			log.Printf("Failed to get container info: %v", err)
			continue
		}

		fmt.Printf("ID: %s\n", info.ID)
		result = append(result, ContainerInfo{
			ID:    info.ID,
			Name:  info.Labels["nerdctl/name"],
			Image: info.Image,
		})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
