package api

import (
	"net/http"

	"github.com/kyg9823/gontainer/api/handler"
)

func NewAPIRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/healthcheck", handler.HealthcheckHandler)

	router.HandleFunc("GET /api/v1/containers", handler.ContainerListHandler)
	router.HandleFunc("GET /api/v1/images", handler.ImageListHandler)

	return router
}
