package models

type ErrorResponse struct {
	ErrorMessage string `json:"message"`
	Field        string `json:"field"`
}
