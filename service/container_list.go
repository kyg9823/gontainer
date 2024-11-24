package service

import (
	"context"
	"log"

	"github.com/kyg9823/gontainer/config"
	"github.com/kyg9823/gontainer/pkg/clientutil"
	"github.com/kyg9823/gontainer/pkg/types"
)

func GetContainerList() ([]types.ContainerInfo, error) {

	address := config.GetContainerdAddress()

	client, ctx, cancel, err := clientutil.NewClient(context.Background(), address, "default")
	if err != nil {
		return nil, err
	}
	defer cancel()

	containers, err := client.Containers(ctx)
	if err != nil {
		log.Fatalf("Failed to list containers: %v", err)
	}

	log.Printf("Found %d containers\n", len(containers))

	result := []types.ContainerInfo{}
	for _, container := range containers {
		info, err := container.Info(ctx)
		if err != nil {
			log.Printf("Failed to get container info: %v", err)
			continue
		}

		task, err := container.Task(ctx, nil)
		if err != nil {
			log.Printf("Failed to get container task: %v", err)
		}

		var statusStr string = "unknown"

		if task != nil {
			status, err := task.Status(ctx)
			if err != nil {
				log.Printf("failed to get status for container %s: %v", container.ID(), err)
				continue
			}

			statusStr = string(status.Status)
		}

		log.Printf("ID: %s\n", info.ID)
		result = append(result, types.ContainerInfo{
			ID:     info.ID[:12],
			LongID: info.ID,
			Name:   info.Labels["nerdctl/name"],
			Image:  info.Image,
			Status: statusStr,
		})

	}
	return result, nil
}
