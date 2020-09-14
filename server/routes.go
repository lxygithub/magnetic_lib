package server

import (
	"github.com/server/pages"
	"log"
	"net/http"
)

func Route() {
	http.HandleFunc("/", pages.Home)
	http.HandleFunc("/add", pages.AddAddress)
	http.HandleFunc("/delete", pages.Delete)
	http.HandleFunc("/web/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
