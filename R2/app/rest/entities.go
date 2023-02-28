package rest

type Header struct {
	User string `json:"user" validate:"required"`
	Pass string `json:"pass" validate:"required"`
}