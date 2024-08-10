package frontend

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var embeddedFiles embed.FS

func GetFrontendFS() fs.FS {
	webAssets, err := fs.Sub(embeddedFiles, "dist")
	if err != nil {
		panic(err) // TODO: handle error instead of panicking
	}
	return webAssets
}
