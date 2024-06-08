package modules

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only POST method is allowed", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	tpl.ExecuteTemplate(w, "404.html", nil)
}
