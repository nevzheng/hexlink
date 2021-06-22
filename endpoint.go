package hexlink

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nevzheng/hexlink/shortener"
	t "github.com/nevzheng/hexlink/types"
)

type Endpoints struct {
	// Create a Batch of Redirects
	CreateRedirects endpoint.Endpoint
	// Query a Batch of Codes
	QueryRedirects endpoint.Endpoint
	// Resolve a Single Code
	GetRedirect endpoint.Endpoint
	// Shorten a Single URL
	CreateRedirect endpoint.Endpoint
}

func MakeEndpoints(s shortener.RedirectService) Endpoints {
	return Endpoints{
		CreateRedirects: makeCreateRedirectsEndpoint(s),
		QueryRedirects:  makeQueryRedirectsEndpoint(s),
		GetRedirect:     makeGetRedirectEndpoint(s),
		CreateRedirect:  makeCreateRedirectEndpoint(s),
	}
}

func makeCreateRedirectsEndpoint(s shortener.RedirectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRedirectsRequest)
		resp := new(CreateRedirectsResponse)
		for _, url := range req.Urls {
			redirect := &t.Redirect{Url: t.URL(url)}
			if code, err := s.Store(redirect); err != nil {
				be := BatchError{
					Code:   code,
					Reason: err.Error(),
				}
				resp.Failed = append(resp.Failed, be)
			} else {
				resp.Successful = append(resp.Successful, *redirect)
			}
		}
		// Any Failures are encoded in the failure array\
		// Larger system failures are out of scope
		return resp, nil
	}
}

func makeQueryRedirectsEndpoint(s shortener.RedirectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(QueryRedirectsRequest)
		resp := new(QueryRedirectsResponse)
		for _, code := range req.Codes {
			if redirect, err := s.Find(string(code)); err != nil {
				be := BatchError{
					Code:   string(code),
					Reason: shortener.ErrRedirectNotFound.Error(),
				}
				resp.Failed = append(resp.Failed, be)
			} else {
				resp.Successful = append(resp.Successful, *redirect)
			}
		}
		// Any Failures are encoded in the failure array\
		// Larger system failures are out of scope
		return resp, nil
	}
}

func makeGetRedirectEndpoint(s shortener.RedirectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRedirectRequest)
		if redirect, err := s.Find(req.Code); err != nil {
			return nil, shortener.ErrRedirectNotFound
		} else {
			return GetRedirectResponse{
				Url: string(redirect.Url),
			}, nil
		}
	}
}

func makeCreateRedirectEndpoint(s shortener.RedirectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRedirectRequest)
		redirect := &t.Redirect{Url: t.URL(req.Url)}
		code, err := s.Store(redirect)
		return CreateRedirectResponse{Code: code}, err
	}
}
