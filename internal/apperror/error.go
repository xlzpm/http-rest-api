package apperror

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNotFound = NewAppError(nil, "not found", "", "RAT-003000")
)

type ErrorFields map[string]string
type ErrorParams map[string]string

type AppError struct {
	Err              error       `json:"-"`
	Message          string      `json:"message,omitempty"`
	DeveloperMessage string      `json:"developer_message,omitempty"`
	Code             string      `json:"code,omitempty"`
	Fields           ErrorFields `jspn:"fields,omitempty"`
	Params           ErrorParams `json:"params,omitempty"`
}

func (e *AppError) WithFields(fields ErrorFields) {
	e.Fields = fields
}

func (e *AppError) WithParams(params ErrorParams) {
	e.Params = params
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

func NewAppError(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func systemError(err error) *AppError {
	return NewAppError(err, "internal system error", err.Error(), "RAT-000000")
}

func BadRequest(message, devMessage string) *AppError {
	return NewAppError(fmt.Errorf(message), message, devMessage, "RAT-000001")
}
