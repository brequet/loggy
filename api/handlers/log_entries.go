package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/brequet/loggy/database"
	"github.com/go-chi/chi/v5"
)

const (
	DEFAULT_PAGE      = 1
	DEFAULT_PAGE_SIZE = 10
)

func LogEntriesHandler(db *database.SQLiteDB) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

		if page == 0 {
			page = DEFAULT_PAGE
		}
		if pageSize == 0 {
			pageSize = DEFAULT_PAGE_SIZE
		}

		entries, err := db.GetLogEntries(page, pageSize)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(entries)
	})

	return r
}
