package main

import (
	"groupie-tracker/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/band", handlers.Band)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	log.Println("Launch server on 4000 port")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
