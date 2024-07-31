package utils

import "errors"

type Error struct {
	InternalError error         `json:"internalError"`
	ExternalError ExternalError `json:"externalError"`
}

type ExternalError struct {
	Id            string `json:"id"`
	Message       string `json:"message"`
	Code          int    `json:"code"`
	TransactionId string `json:"transactionId"`
}

const (
	IncidentNotFound string = "incident not found"
)

var (
	ErrIncidentNotFound = errors.New(IncidentNotFound)
)

// GenerateError
// Generates an error message that is used when processing API requests.
func GenerateError(id string, message string, code int, transactionId string, err error) *Error {
	return &Error{
		InternalError: err,
		ExternalError: ExternalError{
			Id:            id,
			Message:       message,
			Code:          code,
			TransactionId: transactionId,
		},
	}
}
