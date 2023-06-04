package dashboard

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist
var BuildFs embed.FS

// Get the subtree of the embedded files with `build` directory as a root.
func BuildHTTPFS() http.FileSystem {
	build, err := fs.Sub(BuildFs, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(build)
}

type Embed struct {
	http.FileSystem
}

func (e Embed) Open(name string) (http.File, error) {
	if f, err := e.FileSystem.Open(name); err == nil {
		return f, err
	} else {
		return e.FileSystem.Open("index.html")
	}
}
