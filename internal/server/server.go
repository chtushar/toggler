package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chtushar/toggler/internal/logger"
	"github.com/chtushar/toggler/internal/server/web"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Server struct {
	server    http.Server
	router    *mux.Router
	connClose chan int
	logger    *zap.Logger
}

type Config struct {
	Port   int
	Logger *zap.Logger
}

func NewServer(cfg *Config, db *sqlx.DB) *Server {
	r := mux.NewRouter().StrictSlash(true)
	return &Server{
		server: http.Server{
			Addr:         fmt.Sprintf("%s:%d", "", cfg.Port),
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
	defer s.graceFullShutdown()
	web.Routes(s.router, s.logger)

	s.server.Handler = s.router

	// handlers queue
	s.server.Handler = logger.NewHandler(s.logger)(s.server.Handler)
}

func (s *Server) WaitForShutdown() {
	<-s.connClose
}

func (s *Server) graceFullShutdown() {
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)

		sig := <-sigint
		s.logger.Info("OS terminate signal received", zap.String("signal", sig.String()))

		s.logger.Debug("Shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := s.server.Shutdown(ctx)
		if err != nil {
			s.logger.Error("Error shutting down server", zap.Error(err))
		}

		close(s.connClose)
	}()
}
