package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"rest-api/handlers"
	"rest-api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	expectedToken := os.Getenv("API_TOKEN")
	if expectedToken == "" {
		log.Fatal("API_TOKEN не задан")
	}

	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal("Ошибка открытия базы данных:", err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(expectedToken))

		r.Get("/count", handlers.CountItemsHandler(db))
		r.Get("/last_created_at", handlers.LastCreatedAtHandler(db))
	})

	http.ListenAndServe(":8080", r)
}
