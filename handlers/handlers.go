package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "modernc.org/sqlite"
)

func CountItemsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM items").Scan(&count)
		if err != nil {
			http.Error(w, "Ошибка при запросе количества", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]int{"count": count})
	}
}

func LastCreatedAtHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ts sql.NullString
		err := db.QueryRow("SELECT MAX(created_at) FROM items").Scan(&ts)
		if err != nil {
			http.Error(w, "Ошибка при получении даты", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"last_created_at": ts.String,
		})
	}
}
