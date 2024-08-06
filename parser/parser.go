package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/brequet/loggy/entity"
)

type LogFormat struct {
	Name        string
	DateFormat  string
	RegexParser *regexp.Regexp
}

type Parser struct {
	formats []*LogFormat
}

func NewParser() *Parser {
	return &Parser{
		formats: []*LogFormat{
			{
				Name:       "StandardFormat",
				DateFormat: "2006-01-02 15:04:05",
				RegexParser: regexp.MustCompile(
					`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})\s+(DEBUG|INFO|WARN|ERROR|FATAL)\s+(.+)$`,
				),
			},
			{
				Name:       "SpringBootFormat",
				DateFormat: "2006-01-02T15:04:05.000-07:00",
				RegexParser: regexp.MustCompile(
					`^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}[-+]\d{2}:\d{2})\s+(DEBUG|INFO|WARN|ERROR|FATAL)\s+\d+\s+---\s+\[.+?\]\s+.+?\s+:\s+(.+)$`,
				),
			},
			{
				Name:       "ActiveMqFormat",
				DateFormat: "2006-01-02 15:04:05.000",
				RegexParser: regexp.MustCompile(
					`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3})\s+\|\s+(DEBUG|INFO|WARN|ERROR|FATAL)\s+\|\s+(.+)\s+\|\s+.+\s+\|\s.+?$`,
				),
			},
			{
				Name:       "PostgresFormat",
				DateFormat: "2006-01-02 15:04:05 MST",
				RegexParser: regexp.MustCompile(
					`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} [A-Z]+) \[\d+\]: \[\d+-\d+\] user=\[.*\],db=\[.*\],app=\[.*\],client=\d+\.\d+\.\d+\.\d+ (LOG|ERROR|FATAL|WARNING):\s+(.+)$`,
				),
			},
		},
	}
}

func (p *Parser) ParseLogFile(filePath string) ([]entity.LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []entity.LogEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		entry, err := p.parseLine(line)
		if err != nil {
			// TODO: slog debug ?
			// fmt.Printf("Error parsing line: %v\n", err)
			continue
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func (p *Parser) parseLine(line string) (entity.LogEntry, error) {
	for _, format := range p.formats {
		matches := format.RegexParser.FindStringSubmatch(line)
		if matches != nil {
			timestamp, err := time.Parse(format.DateFormat, matches[1])
			if err != nil {
				return entity.LogEntry{}, fmt.Errorf("failed to parse date: %w", err)
			}

			return entity.LogEntry{
				Timestamp: timestamp,
				Level:     mapLevel(strings.ToUpper(matches[2])),
				Content:   matches[3],
				Raw:       line,
			}, nil
		}
	}

	return entity.LogEntry{}, fmt.Errorf("line does not match any known format: %s", line)
}

func mapLevel(level string) entity.LogLevel {
	switch level {
	case "DEBUG":
		return entity.DEBUG
	case "INFO", "LOG":
		return entity.INFO
	case "WARN", "WARNING":
		return entity.WARN
	case "ERROR":
		return entity.ERROR
	case "FATAL":
		return entity.FATAL
	default:
		return entity.INFO
	}
}

func (p *Parser) AddFormat(name, dateFormat, regexPattern string) error {
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return fmt.Errorf("invalid regex pattern: %w", err)
	}

	p.formats = append(p.formats, &LogFormat{
		Name:        name,
		DateFormat:  dateFormat,
		RegexParser: regex,
	})

	return nil
}
