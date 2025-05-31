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

type ErrorInternalServerWrapper struct {
	Ok          bool   `json:"ok" example:"false"`
	ErrorCode   int    `json:"error_code" example:"500"`
	Description string `json:"description" example:"internal server error"`
}
