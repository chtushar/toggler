package server

import (
	"net/http"
	"time"

	"github.com/chtushar/toggler.in/internal/logger"
	"github.com/chtushar/toggler.in/internal/server/web"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	server    http.Server
	router    *mux.Router
	connClose chan int
	logger    *zap.Logger
}

type Config struct {
	Logger *zap.Logger
}

func NewServer(cfg *Config) *Server {
	r := mux.NewRouter().StrictSlash(true)
	return &Server{
		server: http.Server{
			Addr:         ":8080",
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
		router:    r,
		connClose: make(chan int, 1),
		logger:    cfg.Logger,
	}
}

func (s *Server) Listen() {
	s.setup()

	s.logger.Info("Starting server", zap.String("addr", s.server.Addr))
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.Fatal("HTTP server error", zap.Error(err))
	}
}

func (s *Server) setup() {
	web.Routes(s.router, s.logger)

	s.server.Handler = s.router

	// handlers queue
	s.server.Handler = logger.NewHandler(s.logger)(s.server.Handler)
}

func (s *Server) WaitForShutdown() {
	<-s.connClose
}
