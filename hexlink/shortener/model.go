package shortener

type Redirect struct {
	Code      string `json:"code,omitempty"`
	URL       string `json:"url,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
