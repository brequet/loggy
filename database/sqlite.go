package database

import (
	"database/sql"

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
