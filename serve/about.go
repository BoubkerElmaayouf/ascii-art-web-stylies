package modules

import "net/http"

func About(w http.ResponseWriter, r *http.Request) {
	if  r.Method != "GET" {
		http.Error(w, "Bad Request: Only POST method is allowed", http.StatusBadRequest)
		return
	}
	tpl.ExecuteTemplate(w, "about.html", nil)
}
