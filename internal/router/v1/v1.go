package v1

import (
	"github.com/chtushar/toggler/internal/db"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Config struct {
	R      *mux.Router
	DB     *db.DB
	Logger *zap.Logger
}

func Routes(cfg *Config) {

}
