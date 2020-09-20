package server

import (
	"github.com/server/api"
	"log"
	"net/http"
)

func Route() {
	http.HandleFunc("/", api.Home)
	http.HandleFunc("/login", api.Login)
	http.HandleFunc("/add", api.AddAddress)
	http.HandleFunc("/manage", api.Manage)
	http.HandleFunc("/delete", api.Delete)
	http.HandleFunc("/web/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
