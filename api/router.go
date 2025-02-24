package api

import (
	"net/http"

	_ "github.com/kyg9823/gontainer/api/docs"
	"github.com/kyg9823/gontainer/api/handler"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewAPIRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/gontainer/api/v1/healthcheck", handler.HealthcheckHandler)

	router.Get("/gontainer/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/gontainer/swagger/doc.json"), //The url pointing to API definition
	))

	fs := http.FileServer(http.Dir("../../web/static"))
	router.Handle("/gontainer/*", http.StripPrefix("/gontainer/", fs))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/gontainer/", http.StatusFound)
	})
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/gontainer/vite.svg", http.StatusFound)
	})

	router.Get("/gontainer/api/v1/containers", handler.ContainerListHandler)
	router.Get("/gontainer/api/v1/images", handler.ImageListHandler)

	router.Post("/gontainer/api/v1/node", handler.NodeRegisterHandler)

	return router
}
