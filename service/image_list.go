package service

import (
	"context"
	"log"
	"strings"

	"github.com/kyg9823/gontainer/config"
	"github.com/kyg9823/gontainer/pkg/clientutil"
	"github.com/kyg9823/gontainer/pkg/types"
)

func GetImageList() ([]types.ImageInfo, error) {
	address := config.GetContainerdAddress()

	client, ctx, cancel, err := clientutil.NewClient(context.Background(), address, "default")
	if err != nil {
		return nil, err
	}
	defer cancel()

	images, err := client.ImageService().List(ctx)
	if err != nil {
		log.Fatalf("Failed to list images: %v", err)
	}

	log.Printf("Found %d images.\n", len(images))

	result := []types.ImageInfo{}
	for _, image := range images {
		if err != nil {
			log.Fatalf("Failed to get image size: %v", err)
			continue
		}

		result = append(result, types.ImageInfo{
			ID:         (strings.Split(string(image.Target.Digest), ":")[1])[:12],
			Digest:     string(image.Target.Digest),
			Repository: strings.Split(image.Name, ":")[0],
			Tag:        strings.Split(image.Name, ":")[1],
			Path:       image.Name,
			Size:       image.Target.Size,
		})
	}

	return result, nil
}
