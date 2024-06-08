package modules

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	input := r.FormValue("input")
	data := struct {
		Input  string
		Option string
		Output string
	}{
		Input:  input,
		Option: "",
		Output: "",
	}
	tpl.ExecuteTemplate(w, "index.html", data)
}
