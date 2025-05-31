package wrapper

type ErrorNotFoundWrapper struct {
	Ok          bool   `json:"ok" example:"false"`
	ErrorCode   int    `json:"error_code" example:"404"`
	Description string `json:"description" example:"data not found"`
}

type ErrorInvalidObjectIdWrapper struct {
	Ok          bool   `json:"ok" example:"false"`
	ErrorCode   int    `json:"error_code" example:"400"`
	Description string `json:"description" example:"invalid object id"`
}

type ErrorBadRequestWrapper struct {
	Ok          bool   `json:"ok" example:"false"`
	ErrorCode   int    `json:"error_code" example:"400"`
	Description string `json:"description" example:"bad request"`
}

type ErrorUnprocessableEntityWrapper struct {
	Ok          bool   `json:"ok" example:"false"`
	ErrorCode   int    `json:"error_code" example:"422"`
	Description string `json:"description" example:"unprocessable entity"`
}

type ErrorInternalServerWrapper struct {
	Ok          bool   `json:"ok" example:"false"`
	ErrorCode   int    `json:"error_code" example:"500"`
	Description string `json:"description" example:"internal server error"`
}
