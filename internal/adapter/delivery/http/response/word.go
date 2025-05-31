package response

type RetrievedWord struct {
	ID       string  `json:"_id" example:"6835a2db5a859aff5197007a"`
	Bare     string  `json:"bare" example:"весь"`
	Accented string  `json:"accented" example:"весь"`
	Type     *string `json:"type,omitempty" example:"noun"`
	Level    *string `json:"level,omitempty" example:"B1"`
}
