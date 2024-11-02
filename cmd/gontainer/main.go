package main

import (
	"log"
	"net/http"

	"github.com/kyg9823/gontainer/api"
)

func main() {
	router := api.NewAPIRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
