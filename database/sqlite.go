package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/brequet/loggy/entity"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB(dbPath string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTable(db); err != nil {
		return nil, err
	}

	return &SQLiteDB{db: db}, nil
}

func (s *SQLiteDB) Close() error {
	return s.db.Close()
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS log_entries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			timestamp DATETIME,
			app_name TEXT,
			filename TEXT,
			level TEXT,
			content TEXT,
			raw TEXT
		);
		CREATE INDEX IF NOT EXISTS idx_timestamp ON log_entries(timestamp);
	`)
	return err
}

func (s *SQLiteDB) CleanLogEntries() error {
	_, err := s.db.Exec(`
		DELETE FROM log_entries;
	`)
	return err
}

func (s *SQLiteDB) InsertLogEntry(entry entity.LogEntry) error {
	_, err := s.db.Exec(`
		INSERT INTO log_entries (timestamp, app_name, filename, level, content, raw)
		VALUES (?, ?, ?, ?, ?, ?)
	`, entry.Timestamp, entry.AppName, entry.Filename, entry.Level, entry.Content, entry.Raw)
	return err
}

func (s *SQLiteDB) InsertLogEntries(entries []entity.LogEntry) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO log_entries (timestamp, app_name, filename, level, content, raw)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, entry := range entries {
		_, err = stmt.Exec(entry.Timestamp, entry.AppName, entry.Filename, entry.Level, entry.Content, entry.Raw)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *SQLiteDB) GetLogEntries(page int, pageSize int, appNames []string, levels []string, startDate *time.Time, endDate *time.Time) ([]entity.LogEntry, error) {
	query := `
		SELECT timestamp, app_name, filename, level, content, raw
		FROM log_entries
		WHERE 1 = 1
	`

	var args []interface{}
	if len(appNames) > 0 {
		query += ` AND app_name IN (?` + strings.Repeat(",?", len(appNames)-1) + `)`
		for _, name := range appNames {
			args = append(args, name)
		}
	}

	if len(levels) > 0 {
		query += ` AND level IN (?` + strings.Repeat(",?", len(levels)-1) + `)`
		for _, level := range levels {
			args = append(args, level)
		}
	}

	if startDate != nil {
		query += ` AND timestamp >= ?`
		args = append(args, *startDate)
	}

	if endDate != nil {
		query += ` AND timestamp <= ?`
		args = append(args, *endDate)
	}

	query += ` ORDER BY timestamp ASC LIMIT ? OFFSET ?`
	args = append(args, pageSize, (page-1)*pageSize)

	fmt.Printf("query: %s\n", query)
	fmt.Printf("args: %v\n", args)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	entries := make([]entity.LogEntry, 0)
	for rows.Next() {
		var entry entity.LogEntry
		if err := rows.Scan(&entry.Timestamp, &entry.AppName, &entry.Filename, &entry.Level, &entry.Content, &entry.Raw); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
