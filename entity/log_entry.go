package entity

import (
	"fmt"
	"time"
)

type LogEntriesResult struct {
	Entries    []LogEntry `json:"entries"`
	TotalCount int        `json:"total_count"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
}

type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
	FATAL LogLevel = "FATAL"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     LogLevel  `json:"level"`
	Content   string    `json:"content"`
	AppName   string    `json:"app_name"`
	Filename  string    `json:"filename"`
	Raw       string    `json:"raw"`
}

type LogFormat struct {
	DateFormat   string
	DateRegex    string
	LevelRegex   string
	ContentRegex string
}

var DefaultLogFormat = LogFormat{
	DateFormat:   "2006-01-02 15:04:05",
	DateRegex:    `(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`,
	LevelRegex:   `(DEBUG|INFO|WARN|ERROR|FATAL)`,
	ContentRegex: `(.+)`,
}

func (le *LogEntry) String() string {
	return fmt.Sprintf("%s [%s] %s", le.Timestamp.Format(time.RFC3339), le.Level, le.Content)
}
