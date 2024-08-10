package handlers

import (
	"net/http"

	"github.com/brequet/loggy/frontend"
)

func FrontendHandler() http.Handler {
	return http.FileServer(http.FS(frontend.GetFrontendFS()))
	// fs := http.FileServer(http.FS(frontend.GetFrontendFS()))

	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// Serve index.html for requests that don't match any static file
	// 	if _, err := frontend.GetFrontendFS().Open(strings.TrimPrefix(r.URL.Path, "/")); err != nil {
	// 		r.URL.Path = "/index.html"
	// 	}
	// 	fs.ServeHTTP(w, r)
	// })
}
