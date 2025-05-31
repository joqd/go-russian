package domain

type Word struct {
	ID       string
	Bare     string
	Accented string
	Type     *string
	Level    *string
	Disable  bool
}
