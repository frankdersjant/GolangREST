package Domain

import "net/http"

type AppError struct {
	Code    int    `json:"omitempty"`
	Message string `json:"message"`
}

func (e *AppError) AsMessage() string {
	return e.Message
}

func NotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

func UnexpectedError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}
