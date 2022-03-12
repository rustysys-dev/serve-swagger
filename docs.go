package main

import (
	"embed"
	"io/fs"
	"log"

	"net/http"
)

//go:embed "swagger"
//go:embed "swagger.json"
var swaggerRoot embed.FS

//go:embed "swagger-dev/*"
//go:embed "swagger.json"
var swaggerDevRoot embed.FS

// SwaggerAPI serves swaggerapi
func SwaggerAPI() http.Handler {
	var ffs = make(flattenedFS, 0)
	_ = ffs.FlattenEmbeddedFS(swaggerRoot, ".")

	if Development() {
		_ = ffs.FlattenEmbeddedFS(swaggerDevRoot, ".")
	}

	return http.FileServer(http.FS(ffs))
}

type flattenedFS map[string]fs.File

func (f flattenedFS) FlattenEmbeddedFS(efs embed.FS, dir string) error {
	del, e := efs.ReadDir(dir)
	if e != nil {
		log.Fatalf("failed to ReadDir: %v", e)
	}
	for _, v := range del {
		if v.IsDir() {
			if dir == "." {
				dir = ""
			}
			f.FlattenEmbeddedFS(efs, dir+v.Name())
			continue
		}
		dirLen := len(dir)
		if dirLen != 0 && string(dir[dirLen-1]) != "/" {
			dir = dir + "/"
		}
		file, e := efs.Open(dir + v.Name())
		if e != nil {
			log.Fatalf("failed to get open file: %v", e)
		}
		f[v.Name()] = file
	}
	return nil
}

func (f flattenedFS) Open(name string) (fs.File, error) {
	if name == "." {
		name = "index.html"
	}
	return f[name], nil
}
