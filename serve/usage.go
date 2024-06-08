package modules

import "net/http"

func Usage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only POST method is allowed", http.StatusBadRequest)
		return
	}
	tpl.ExecuteTemplate(w, "usage.html", nil)
}
