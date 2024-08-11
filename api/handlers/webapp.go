package handlers

import (
	"net/http"

	"github.com/brequet/loggy/frontend"
)

func FrontendHandler() http.Handler {
	return http.FileServer(http.FS(frontend.GetFrontendFS()))
}
