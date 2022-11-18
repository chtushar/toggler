package router

import (
	"github.com/chtushar/toggler/internal/db"
	v1 "github.com/chtushar/toggler/internal/router/v1"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Config struct {
	R      *mux.Router
	DB     *db.DB
	Logger *zap.Logger
}

func Routes(cfg *Config) {
	v1.Routes(&v1.Config{
		R:      cfg.R,
		DB:     cfg.DB,
		Logger: cfg.Logger,
	})
}
