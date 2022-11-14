package web

import (
	"github.com/chtushar/toggler/internal/admin"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func Routes(r *mux.Router, log *zap.Logger) {
	admin.Serve(r, log)
}
