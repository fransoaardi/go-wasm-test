package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wasm"))
	}))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	h := cors.Default().Handler(mux)

	s := http.Server{
		Addr:    ":8000",
		Handler: h,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
