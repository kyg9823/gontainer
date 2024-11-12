package api

import (
	_ "github.com/kyg9823/gontainer/api/docs"
	"github.com/kyg9823/gontainer/api/handler"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewAPIRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/gontainer/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/gontainer/swagger/doc.json"), //The url pointing to API definition
	))

	router.Get("/gontainer/api/v1/healthcheck", handler.HealthcheckHandler)

	router.Get("/gontainer/api/v1/containers", handler.ContainerListHandler)
	router.Get("/gontainer/api/v1/images", handler.ImageListHandler)

	return router
}
