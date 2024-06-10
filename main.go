package main

import (
	"habit-counter-htmx/views"
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()
	router.Handle("GET /public/", http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public"))))
	router.HandleFunc("GET /", handler)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
