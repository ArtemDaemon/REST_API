package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "modernc.org/sqlite"
)

type Item struct {
	IndicatorId     string `json:"indicator_id"`
	IndicatorValue  string `json:"indicator_value"`
	CountryId       string `json:"country_id"`
	CountryValue    string `json:"country_value"`
	CountryISO3Code string `json:"country_iso3_code"`
	Date            string `json:"date"`
	Value           uint   `json:"value"`
	Unit            string `json:"unit"`
	ObsStatus       string `json:"obs_status"`
	Decimal         uint   `json:"decimal"`
	CreatedAt       string `json:"created_at"`
}

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

func GetItemByDateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		year := r.URL.Query().Get("year")
		if year == "" {
			http.Error(w, "Параметр 'year' обязателен", http.StatusBadRequest)
			return
		}

		var item Item
		err := db.QueryRow(
			`SELECT indicator_id, indicator_value, country_id, country_value, country_iso3_code, date, value, unit, obs_status, decimal, created_at FROM items WHERE date = ?`, year,
		).Scan(&item.IndicatorId, &item.IndicatorValue, &item.CountryId, &item.CountryValue,
			&item.CountryISO3Code, &item.Date, &item.Value, &item.Unit, &item.ObsStatus, &item.Decimal, &item.CreatedAt)

		if err == sql.ErrNoRows {
			http.Error(w, "Элемент не найден", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, "Ошибка при запросе: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(item)
	}
}
