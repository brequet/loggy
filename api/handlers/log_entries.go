package handlers

import (
	"encoding/json"
	"fmt"
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

		start, err := parseDate(startDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		end, err := parseDate(endDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logEntryResult, err := db.GetLogEntries(page, pageSize, appNames, levels, start, end)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(logEntryResult)
	})

	return r
}

func AppsHandler(db *database.SQLiteDB) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		appNames, err := db.GetAppNames()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(appNames)
	})

	return r
}

func parseDate(dateStr string) (*time.Time, error) {
	layoutWithSeconds := "2006-01-02T15:04:05"
	if dateStr != "" {
		st, err := time.Parse(layoutWithSeconds, dateStr)
		if err == nil {
			return &st, nil
		} else {
			layoutWithSeconds := "2006-01-02T15:04"
			st, err := time.Parse(layoutWithSeconds, dateStr)
			if err == nil {
				return &st, nil
			} else {
				return nil, fmt.Errorf("invalid date format: %s", dateStr)
			}
		}
	}
	return nil, nil
}
