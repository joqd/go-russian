package wrapper

import "github.com/joqd/ruskee/internal/adapter/delivery/http/response"

// RetrievedWord
type RetrievedWordWrapper struct {
	Ok     bool                   `json:"ok" example:"true"`
	Result response.RetrievedWord `json:"result"`
}
