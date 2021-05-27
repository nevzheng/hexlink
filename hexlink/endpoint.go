package hexlink

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/nevzheng/hexlink/hexlink/shortener"
)

type Endpoints struct {
	CreateRedirect endpoint.Endpoint
	GetRedirect    endpoint.Endpoint
}

func MakeEndpoints(s shortener.RedirectService) Endpoints {
	return Endpoints{
		CreateRedirect: makeCreateRedirectEndpoint(s),
		GetRedirect:    makeGetRedirectEndpoint(s),
	}
}

func makeCreateRedirectEndpoint(s shortener.RedirectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRedirectRequest)
		redirect := &shortener.Redirect{URL: req.Url}
		code, err := s.Store(redirect)
		return CreateRedirectResponse{Code: code}, err
	}
}

func makeGetRedirectEndpoint(s shortener.RedirectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRedirectRequest)
		if redirect, err := s.Find(req.Code); err != nil {
			return nil, shortener.ErrRedirectNotFound
		} else {
			return GetRedirectResponse{
				Url: redirect.URL,
			}, nil
		}
	}
}
