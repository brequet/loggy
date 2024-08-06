package ingester

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/brequet/loggy/database"
	"github.com/brequet/loggy/entity"
	"github.com/brequet/loggy/parser"
)

const BATCH_SIZE = 100

type Ingester struct {
	db     *database.SQLiteDB
	parser *parser.Parser
	logger *slog.Logger
}

func NewIngester(db *database.SQLiteDB, parser *parser.Parser, logger *slog.Logger) *Ingester {
	return &Ingester{
		db:     db,
		parser: parser,
		logger: logger,
	}
}

func (s *Ingester) IngestLogs(inputDir string) error {
	s.logger.Debug("Ingesting logs", "inputDir", inputDir)
	apps, err := listSubdirectories(inputDir)
	if err != nil {
		return fmt.Errorf("failed to list subdirectories: %w", err)
	}
	s.logger.Debug("Found apps", "apps count", len(apps))

	for _, app := range apps {
		s.logger.Debug("Ingesting app", "app", app)

		appDir := filepath.Join(inputDir, app)
		logFiles, err := listLogFiles(appDir)
		if err != nil {
			return fmt.Errorf("failed to list log files for %s: %w", app, err)
		}
		s.logger.Debug("Found log files", "logFiles count", len(logFiles))

		for _, logFile := range logFiles {
			s.logger.Debug("Ingesting log file", "logFile", logFile)

			if err := s.ingestLogFile(app, logFile); err != nil {
				return fmt.Errorf("failed to ingest log file %s: %w", logFile, err)
			}
		}
	}

	return nil
}

func (s *Ingester) ingestLogFile(appName, filePath string) error {
	entries, err := s.parser.ParseLogFile(filePath)
	if err != nil {
		return err
	}

	batch := make([]entity.LogEntry, 0, BATCH_SIZE)

	for _, entry := range entries {
		dbEntry := entity.LogEntry{
			Timestamp: entry.Timestamp,
			AppName:   appName,
			Filename:  filepath.Base(filePath),
			Level:     entry.Level,
			Content:   entry.Content,
			Raw:       entry.Raw,
		}
		batch = append(batch, dbEntry)

		if len(batch) == BATCH_SIZE {
			if err := s.db.InsertLogEntries(batch); err != nil {
				return fmt.Errorf("failed to insert log entries batch: %w", err)
			}
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		if err := s.db.InsertLogEntries(batch); err != nil {
			return fmt.Errorf("failed to insert remaining log entries: %w", err)
		}
	}

	return nil
}

func listSubdirectories(dir string) ([]string, error) {
	var subDirs []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDirs = append(subDirs, entry.Name())
		}
	}

	return subDirs, nil
}

func listLogFiles(dir string) ([]string, error) {
	var logFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		filename := filepath.Base(path)
		if !info.IsDir() && strings.Contains(filename, ".log") {
			logFiles = append(logFiles, path)
		}
		return nil
	})

	return logFiles, err
}
