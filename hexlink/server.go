package hexlink

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer() http.Handler {
	// Spawn a Gorilla Mux Router
	r := mux.NewRouter()
	// Add Middlewares
	// Add Methods
	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	return r
}
