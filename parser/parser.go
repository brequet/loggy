package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/brequet/loggy/config"
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

func NewParser(additionalParsersConf []config.ParserFormat) (*Parser, error) {
	formats := []*LogFormat{
		{
			Name:       "StandardFormat",
			DateFormat: "2006-01-02 15:04:05",
			RegexParser: regexp.MustCompile(
				`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})\s+(DEBUG|INFO|WARN|ERROR|FATAL)\s+(.+)$`,
			),
		},
	}

	for _, format := range additionalParsersConf {
		logFormat, err := parseToLogFormat(format.Name, format.DateFormat, format.RegexParser)
		if err != nil {
			return nil, fmt.Errorf("failed to parse format: %w", err)
		}
		formats = append(formats, logFormat)
	}

	return &Parser{
		formats,
	}, nil
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

func parseToLogFormat(name, dateFormat, regexPattern string) (*LogFormat, error) {
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	return &LogFormat{
		Name:        name,
		DateFormat:  dateFormat,
		RegexParser: regex,
	}, nil
}
