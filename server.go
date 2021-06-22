package hexlink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {
	// Spawn a Gorilla Mux Router
	r := mux.NewRouter()
	// Add Middlewares
	// Add Methods

	// Post a collection of URLS returning the redirects
	// POST /api/redirects/createRedirects
	r.Methods("POST").Path("/api/redirects/createRedirects").Handler(
		httptransport.NewServer(
			endpoints.CreateRedirects,
			decodeCreateRedirectsRequest,
			encodeJsonResponse,
		),
	)

	// Post a collection of Codes eturning the redirects
	// POST /api/redirects/queryRedirects
	r.Methods("POST").Path("/api/redirects/queryRedirects").Handler(
		httptransport.NewServer(
			endpoints.QueryRedirects,
			decodeQueryRedirectsRequest,
			encodeJsonResponse,
		),
	)

	// Shorten a URL, returning a redirect Code
	// POST /api/redirects {url: {input}}}
	r.Methods("POST").Path("/api/redirects").Handler(
		httptransport.NewServer(
			endpoints.CreateRedirect,
			decodeCreateRedirectRequest,
			encodeJsonResponse,
		),
	)

	// Follow a Redirect
	// GET /{code}
	r.Methods("GET").Path("/{code}").Handler(
		httptransport.NewServer(
			endpoints.GetRedirect,
			decodeGetRedirectRequest,
			encodeGetRedirectResponse,
		),
	)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})
	return r
}
