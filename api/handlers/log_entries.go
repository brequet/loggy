package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

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
		appNamesStr := r.URL.Query().Get("appNames")
		levelsStr := r.URL.Query().Get("levels")
		startDate := r.URL.Query().Get("startDate")
		endDate := r.URL.Query().Get("endDate")

		if page == 0 {
			page = DEFAULT_PAGE
		}
		if pageSize == 0 {
			pageSize = DEFAULT_PAGE_SIZE
		}

		var appNames, levels []string
		if appNamesStr != "" {
			appNames = strings.Split(appNamesStr, ",")
		}
		if levelsStr != "" {
			levels = strings.Split(levelsStr, ",")
		}

		var start, end *time.Time
		layout := "2006-01-02T15:04:05"
		if startDate != "" {
			st, err := time.Parse(layout, startDate)
			if err == nil {
				start = &st
			}
		}
		if endDate != "" {
			st, err := time.Parse(layout, endDate)
			if err == nil {
				end = &st
			}
		}

		entries, err := db.GetLogEntries(page, pageSize, appNames, levels, start, end)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(entries)
	})

	return r
}
