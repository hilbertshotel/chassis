package handlers

import (
	"chassis/dep"
	"net/http"
)

func Mux(d *dep.Dependencies) *http.ServeMux {
	var mux http.ServeMux

	// handle static files
	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/static/", static)

	// handle index
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r, d)
	})

	// handle websocket server
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws(w, r, d)
	})

	return &mux
}
