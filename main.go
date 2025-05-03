package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest-api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	expectedToken := os.Getenv("API_TOKEN")
	if expectedToken == "" {
		log.Fatal("API_TOKEN не задан")
	}

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(expectedToken))

		r.Get("/route1", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Access granted!") })
	})

	http.ListenAndServe(":8080", r)
}
