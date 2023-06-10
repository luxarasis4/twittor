package dtos

type ErrorResponseDTO struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   any    `json:"data,omitempty"`
	Error  error  `json:"error,omitempty"`
}
