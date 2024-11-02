package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

type ContainerInfo struct {
	ID string `json:"id"`
}

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
	for _, container := range containers {
		info, err := container.Info(ctx)
		if err != nil {
			log.Printf("Failed to get container info: %v", err)
			continue
		}

		fmt.Printf("ID: %s\n", info.ID)
	}
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
	for _, image := range images {
		fmt.Printf("Name: %s\n", image.Name)
	}
}
