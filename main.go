package main

import (
	"log"
	"net/http"

	help "modules/serve"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", help.Index)
	http.HandleFunc("/ascii-art", help.GenerateHandler)
	http.HandleFunc("/about", help.About)
	http.HandleFunc("/usage", help.Usage)
	http.HandleFunc("/404", help.NotFound)

	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
