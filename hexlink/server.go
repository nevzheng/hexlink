package hexlink

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {
	// Spawn a Gorilla Mux Router
	r := mux.NewRouter()
	// Add Middlewares
	// Add Methods
	// GET /{code}
	r.Methods("GET").Path("/{code}").Handler(
		httptransport.NewServer(
			endpoints.GetRedirect,
			decodeGetRedirectRequest,
			encodeGetRedirectResponse))
	// POST /api/shorten {url: {input}}}
	r.Methods("POST").Path("/api/shorten").Handler(
		httptransport.NewServer(
			endpoints.CreateRedirect,
			decodeCreateRedirectRequest,
			encodeJsonResponse,
		),
	)
	return r
}
