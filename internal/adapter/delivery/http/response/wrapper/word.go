package wrapper

import "github.com/joqd/go-russian/internal/adapter/delivery/http/response"

type RetrievedWordWrapper struct {
	Ok     bool                   `json:"ok" example:"true"`
	Result response.RetrievedWord `json:"result"`
}

type DeletedWordWrapper struct {
	Ok     bool                 `json:"ok" example:"true"`
	Result response.DeletedWord `json:"result"`
}
