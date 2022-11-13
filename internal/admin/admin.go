package admin

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

//go:embed ui/dist
var adminUI embed.FS

func Serve(r *mux.Router, log *zap.Logger) {
	sub, err := fs.Sub(adminUI, "ui/dist")
	if err != nil {
		log.Error("failed to get subtree for admin pages", zap.Error(err))
		return
	}

	adminFS := http.FileServer(http.FS(sub))

	r.PathPrefix("/admin").Handler(http.StripPrefix("/admin", adminFS))
}
