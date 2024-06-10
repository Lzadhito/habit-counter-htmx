package main

import (
	"embed"
	"habit-counter-htmx/views"
	"log/slog"
	"net/http"
	"os"
)

var publicFS embed.FS

func main() {
	router := http.NewServeMux()
	router.Handle("/public/", http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public"))))
	router.HandleFunc("GET /foo", handler)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	server.ListenAndServe()
}

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
