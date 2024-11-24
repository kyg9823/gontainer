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

type ImageInfo struct {
	ID         string `json:"id"`
	Repository string `json:"repository"`
	Tag        string `json: "tag"`
}

func ImageListHandler(w http.ResponseWriter, r *http.Request) {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		log.Fatalf("Failed to create containerd client: %v", err)
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), "default")

	images, err := client.ImageService().List(ctx)
	if err != nil {
		log.Fatalf("Failed to list images: %v", err)
	}

	fmt.Printf("Found %d images.\n", len(images))
	result := []ImageInfo{}
	for _, image := range images {
		fmt.Printf("Name: %s\n", image.Name)
		result = append(result, ImageInfo{
			ID: "asdf",
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
