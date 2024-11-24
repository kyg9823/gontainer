package handler

import "net/http"

type Node struct {
	ID       string `json:"id"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

func NodeRegisterHandler(w http.ResponseWriter, r *http.Request) {

}
