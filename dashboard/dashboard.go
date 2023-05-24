package dashboard

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:out
var BuildFs embed.FS

// Get the subtree of the embedded files with `build` directory as a root.
func BuildHTTPFS() http.FileSystem {
	build, err := fs.Sub(BuildFs, "out")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(build)
}
