package hexlink

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	t "github.com/nevzheng/hexlink/types"
)

type (
	BatchError struct {
		Code   string `json:"code"`
		Reason string `json:"reason"`
	}

	CreateRedirectsRequest struct {
		Urls []t.URL `json:"urls"`
	}

	CreateRedirectsResponse struct {
		Successful []t.Redirect `json:"successful"`
		Failed     []BatchError `json:"failed"`
	}

	QueryRedirectsResponse struct {
		Successful []t.Redirect `json:"successful"`
		Failed     []BatchError `json:"failed"`
	}

	QueryRedirectsRequest struct {
		Codes []t.Code `json:"codes"`
	}

	CreateRedirectRequest struct {
		Url string `json:"url,omitempty"`
	}

	CreateRedirectResponse struct {
		Code string `json:"code,omitempty"`
	}

	GetRedirectRequest struct {
		Code string `json:"code,omitempty"`
	}

	GetRedirectResponse struct {
		Url string `json:"url,omitempty"`
	}
)

func decodeCreateRedirectsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateRedirectsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeQueryRedirectsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req QueryRedirectsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil

}

func encodeGetRedirectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	r := response.(GetRedirectResponse)
	w.Header().Set("Location", r.Url)
	w.WriteHeader(http.StatusSeeOther)
	return nil
}

func decodeGetRedirectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	code := params["code"]
	return GetRedirectRequest{Code: code}, nil
}

func encodeJsonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateRedirectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateRedirectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
