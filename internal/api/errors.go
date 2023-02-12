package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ErrNotFound = NewAppError("not found", http.StatusNotFound)
)

type AppError struct {
	Err      error  `json:"-"`
	HttpCode int    `json:"code"`
	Message  string `json:"message,omitempty"`
}

func NewAppError(message string, code int) *AppError {
	return &AppError{
		Err:      fmt.Errorf(message),
		Message:  message,
		HttpCode: code,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return bytes
}

func UnauthorizedError(message string) *AppError {
	return NewAppError(message, http.StatusUnauthorized)
}

func BadRequestError(message string) *AppError {
	return NewAppError(message, http.StatusBadRequest)
}

func SystemError(message string) *AppError {
	return NewAppError(message, http.StatusInternalServerError)
}

func APIError(message string, code int) *AppError {
	return NewAppError(message, code)
}
