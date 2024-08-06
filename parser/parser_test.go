package parser

import (
	"os"
	"testing"
	"time"

	"github.com/brequet/loggy/entity"
)

func TestParseLine(t *testing.T) {
	s := NewParser()

	tests := []struct {
		name    string
		input   string
		want    entity.LogEntry
		wantErr bool
	}{
		{
			name:  "Valid log entry",
			input: "2023-04-10 15:04:05 INFO This is a log message",
			want: entity.LogEntry{
				Timestamp: time.Date(2023, 4, 10, 15, 4, 5, 0, time.UTC),
				Level:     entity.INFO,
				Content:   "This is a log message",
			},
			wantErr: false,
		},
		{
			name:    "Invalid date format",
			input:   "2023/04/10 15:04:05 INFO This is a log message",
			want:    entity.LogEntry{},
			wantErr: true,
		},
		{
			name:    "No date in log entry",
			input:   "This is a log message without a date",
			want:    entity.LogEntry{},
			wantErr: true,
		},
		{
			name:  "SpringBoot log entry",
			input: "2024-05-16T09:24:06.067+02:00  WARN 15400 --- [restartedMain] o.f.c.internal.database.base.Database    : Flyway upgrade recommended: PostgreSQL 16.0 is newer than this version of Flyway and support has not been tested. The latest supported version of PostgreSQL is 15.",
			want: entity.LogEntry{
				Timestamp: time.Date(2024, 5, 16, 9, 24, 6, 67000000, time.FixedZone("CEST", 2*60*60)),
				Level:     entity.WARN,
				Content:   "Flyway upgrade recommended: PostgreSQL 16.0 is newer than this version of Flyway and support has not been tested. The latest supported version of PostgreSQL is 15.",
			},
			wantErr: false,
		},
		{
			name:  "Activemq log entry",
			input: "2024-08-06 08:35:13,785 | INFO  | BMXBROKER bridge to horusAdminBroker stopped | org.apache.activemq.network.DemandForwardingBridgeSupport | ActiveMQ BrokerService[BMXBROKER] Task-71632",
			want: entity.LogEntry{
				Timestamp: time.Date(2024, 8, 6, 8, 35, 13, 785000000, time.UTC),
				Level:     entity.INFO,
				Content:   "BMXBROKER bridge to horusAdminBroker stopped",
			},
			wantErr: false,
		},
		{
			name:  "Postgres log entry",
			input: "2024-07-24 13:09:16 UTC [15412]: [1-1] user=[unknown],db=[unknown],app=[unknown],client=127.0.0.1 LOG:  connection received: host=127.0.0.1 port=53138",
			want: entity.LogEntry{
				Timestamp: time.Date(2024, 7, 24, 13, 9, 16, 0, time.UTC),
				Level:     entity.INFO,
				Content:   "connection received: host=127.0.0.1 port=53138",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.parseLine(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Level != tt.want.Level {
				t.Errorf("parseLine() level = %v, want %v", got.Level, tt.want.Level)
			}
			if got.Content != tt.want.Content {
				t.Errorf("parseLine() content = %v, want %v", got.Content, tt.want.Content)
			}
			if got.Timestamp.Unix() != tt.want.Timestamp.Unix() {
				t.Errorf("parseLine() timestamp = %v, want %v", got.Timestamp, tt.want.Timestamp)
			}
		})
	}
}

func TestParseLogFile(t *testing.T) {
	s := NewParser()

	// Create a temporary log file for testing
	tmpfile, err := os.CreateTemp("", "test*.log")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Write test log entries to the file
	testLogs := []string{
		"2023-04-10 15:04:05 INFO First log entry",
		"2023-04-10 15:04:06 ERROR Second log entry",
		"Invalid log entry",
		"2023-04-10 15:04:07 DEBUG Third log entry",
	}
	for _, log := range testLogs {
		if _, err := tmpfile.WriteString(log + "\n"); err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}
	}
	tmpfile.Close()

	// Test ParseLogFile
	entries, err := s.ParseLogFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("ParseLogFile() error = %v", err)
	}

	// Check the number of valid entries
	expectedEntries := 3 // We expect 3 valid entries out of 4 lines
	if len(entries) != expectedEntries {
		t.Errorf("ParseLogFile() returned %d entries, want %d", len(entries), expectedEntries)
	}

	// Check the content of the first and last valid entries
	if entries[0].Content != "First log entry" {
		t.Errorf("First entry content = %s, want %s", entries[0].Content, "First log entry")
	}
	if entries[2].Content != "Third log entry" {
		t.Errorf("Last entry content = %s, want %s", entries[2].Content, "Third log entry")
	}
}

func TestParseLogFile_NonexistentFile(t *testing.T) {
	s := NewParser()

	_, err := s.ParseLogFile("nonexistent.log")
	if err == nil {
		t.Error("ParseLogFile() error = nil, want error for nonexistent file")
	}
}
