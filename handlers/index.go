package handlers

import (
	"chassis/dep"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {
	// handle method
	if r.Method != http.MethodGet {
		http.Error(w, "Not Implemented", 501)
		return
	}

	// return template
	if err := d.Tmp.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
	}
}
