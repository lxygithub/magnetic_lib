package server

import (
	"github.com/server/pages"
	"log"
	"net/http"
)

func Route() {
	http.HandleFunc("/", pages.Home)
	http.HandleFunc("/add", pages.AddAddress)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
