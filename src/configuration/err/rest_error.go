package rest_error

import "net/http"

type Err struct {
	Message   string   `json: "message"`
	HttpDescr string   `json: "httpDescr"`
	HttpCode  int      `json: "httpCode"`
	Causes    []Causes `json: "causes,omitempty"`
}

type Causes struct {
	Field   string `json: "field"`
	Message string `json: "message"`
}

func (err *Err) Error() string {
	return err.Message
}

func NewError(message, httpDescr string, httpCode int, causes []Causes) *Err {
	return &Err{
		Message:   message,
		HttpDescr: httpDescr,
		HttpCode:  httpCode,
		Causes:    causes,
	}
}

func NewBadRequestError(message string) *Err {
	return &Err{
		Message:   message,
		HttpDescr: "Bad Request",
		HttpCode:  http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *Err {
	return &Err{
		Message:   message,
		HttpDescr: "Bad Request",
		HttpCode:  http.StatusBadRequest,
		Causes:    causes,
	}
}

func NewInternalServerError(message string) *Err {
	return &Err{
		Message:   message,
		HttpDescr: "Internal Server Error",
		HttpCode:  http.StatusInternalServerError,
	}
}

func NewForbiddenError(message string) *Err {
	return &Err{
		Message:   message,
		HttpDescr: "Forbidden",
		HttpCode:  http.StatusForbidden,
	}
}
