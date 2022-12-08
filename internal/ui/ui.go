package ui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

//go:embed dist
var ui embed.FS

func Serve(r *mux.Router, log *zap.Logger) {
	sub, err := fs.Sub(ui, "dist")
	if err != nil {
		log.Error("failed to get subtree for admin pages", zap.Error(err))
		return
	}

	fs := http.FileServer(http.FS(sub))
	r.PathPrefix("/admin").Handler(http.StripPrefix("/admin", fs))
	r.PathPrefix("/admin/*").Handler(http.StripPrefix("/admin/*", fs))
}
