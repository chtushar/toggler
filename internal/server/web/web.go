package web

import (
	"github.com/chtushar/toggler/internal/ui"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func Routes(r *mux.Router, log *zap.Logger) {
	ui.Serve(r, log)
}
