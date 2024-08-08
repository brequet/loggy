package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/brequet/loggy/api"
	"github.com/brequet/loggy/database"
)

type Server struct {
	port   int
	db     *database.SQLiteDB
	logger *slog.Logger
}

func NewServer(port int, logger *slog.Logger) (*Server, error) {
	// TODO: check if db exists
	db, err := database.NewSQLiteDB("loggy.db")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	return &Server{
		port:   port,
		db:     db,
		logger: logger,
	}, nil
}

func (s *Server) Start() error {
	router := api.SetupRoutes(s.db)

	s.logger.Info("Starting server", "port", s.port)
	return http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", s.port), router)
}
