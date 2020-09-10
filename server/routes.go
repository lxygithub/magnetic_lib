package server

import (
	"github.com/server/pages"
	"log"
	"net/http"
)

func Route() {
	http.HandleFunc("/", pages.Home)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
