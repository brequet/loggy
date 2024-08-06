package entity

import (
	"fmt"
	"time"
)

type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
	FATAL LogLevel = "FATAL"
)

type LogEntry struct {
	Timestamp time.Time
	Level     LogLevel
	Content   string
	AppName   string
	Filename  string
	Raw       string
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
