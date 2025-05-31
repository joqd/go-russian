package request

type CreateWord struct {
	Bare     string  `json:"bare" validate:"required,min=1" example:"весь"`
	Accented string  `json:"accented" validate:"required,min=1" example:"весь"`
	Type     *string `json:"type,omitempty" validate:"omitempty,oneof=adjective noun verb adverb other" example:"noun"`
	Level    *string `json:"level,omitempty" validate:"omitempty,oneof=A1 A2 B1 B2 C1 C2" example:"B1"`
	Disable  *bool   `json:"disable,omitempty" validate:"omitempty" example:"false"`
}
