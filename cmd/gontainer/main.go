package main

import (
	"log"
	"net/http"

	"github.com/kyg9823/gontainer/api"
)

// @title          Gontainer
// @version        1.1
// @description    Container manager for Go

// @contact.name Youngkook Kim
// @contact.email theboss@lgcns.com

// @license.name MIT License
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
