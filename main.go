package main

import (
	"fmt"
	"habit-counter-htmx/views"
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()
	router.Handle("GET /public/", http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public"))))
	router.HandleFunc("GET /", handler)
	router.HandleFunc("GET /bad-habit", habitHandler)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}

func habitHandler(w http.ResponseWriter, r *http.Request) {
	isBadHabit := r.URL.Query().Get("bad-habit")
	fmt.Println(isBadHabit == "on")
	if isBadHabit == "on" {
		w.Write([]byte(nil))
	} else {
		views.IsBadHabitSelect().Render(r.Context(), w)
	}
}
