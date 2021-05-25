package hexlink

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"

	http_transport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {
	// Spawn a Gorilla Mux Router
	r := mux.NewRouter()
	// Add Middlewares
	// Add Methods
	// GET /{code}
	r.Methods("GET").Path("/{code}").Handler(
		http_transport.NewServer(
			endpoints.GetRedirect,
			decodeGetRedirectRequest,
			encodeGetRedirectResponse))
	// POST /api/shorten {url: {input}}}
	r.Methods("POST").Path("/api/shorten").Handler(
		http_transport.NewServer(
			endpoints.CreateRedirect,
			decodeCreateRedirectRequest,
			encodeJsonResponse,
		),
	)
	//// Serve static files
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	return r
}
