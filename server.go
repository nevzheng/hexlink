package hexlink

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stderr, next)
}

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {
	// Spawn a Gorilla Mux Router
	r := mux.NewRouter()

	// Add Middlewares
	r.Use(loggingMiddleware)

	// Create Handlers
	// Post a collection of URLS returning the redirects
	createRedirectsHandler := httptransport.NewServer(
		endpoints.CreateRedirects,
		decodeCreateRedirectsRequest,
		encodeJsonResponse,
	)

	// Post a collection of Codes eturning the redirects
	queryRedirectsHandler := httptransport.NewServer(
		endpoints.QueryRedirects,
		decodeQueryRedirectsRequest,
		encodeJsonResponse,
	)

	// Shorten a URL, returning a redirect Code
	postRedirectHandler := httptransport.NewServer(
		endpoints.CreateRedirect,
		decodeCreateRedirectRequest,
		encodeJsonResponse,
	)

	// Follow a Redirect
	followRedirectHandler := httptransport.NewServer(
		endpoints.GetRedirect,
		decodeGetRedirectRequest,
		encodeGetRedirectResponse,
	)

	// Attach Methods
	// POST /api/redirects/createRedirects
	r.Methods("POST").Path("/redirects/createRedirects").Handler(createRedirectsHandler)
	// POST /api/redirects/queryRedirects
	r.Methods("POST").Path("/redirects/queryRedirects").Handler(queryRedirectsHandler)
	// POSredirects {url: {input}}}
	r.Methods("POST").Path("/redirects").Handler(postRedirectHandler)
	// GET /{code}
	r.Methods("GET").Path("/{code}").Handler(followRedirectHandler)

	return r
}
