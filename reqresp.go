package hexlink

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	t "github.com/nevzheng/hexlink/types"
)

type (
	BatchError struct {
		Code   string `json:"code"`
		Reason string `json:"reason"`
	}

	CreateRedirectsRequest struct {
		Urls []string `json:"urls"`
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
		Codes []string `json:"codes"`
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
	var urls []string
	err := json.NewDecoder(r.Body).Decode(&urls)
	if err != nil {
		return nil, err
	}
	return CreateRedirectsRequest{Urls: urls}, nil
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
	// Ensure there is a query parameter
	var qUrl string
	if qUrl = r.URL.Query().Get("url"); qUrl == "" {
		return nil, ServerError{
			reason:      "Invalid request: Missing 'url' query Param",
			status_code: http.StatusBadRequest,
		}
	}
	// ensure the url is valid
	if _, err := url.ParseRequestURI(qUrl); err != nil {
		return nil, ServerError{
			reason:      "Invalid request: Invalid 'url'",
			status_code: http.StatusBadRequest,
		}
	}
	return CreateRedirectRequest{
		Url: qUrl,
	}, nil
}

type ServerError struct {
	reason      string
	status_code int
}

func (err ServerError) Error() string {
	return err.reason
}

func (err ServerError) StatusCode() int {
	return err.status_code
}
