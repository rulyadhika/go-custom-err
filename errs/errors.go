package errs

import "net/http"

type CustomError interface {
	StatusCode() int
	Status() string
	Message() string
}

type generalError struct {
	ErrStatusCode int    `json:"status_code"`
	ErrStatus     string `json:"status"`
	ErrMessage    string `json:"message"`
	Data          any    `json:"data"`
}

func (c *generalError) StatusCode() int {
	return c.ErrStatusCode
}

func (c *generalError) Status() string {
	return c.ErrStatus
}

func (c *generalError) Message() string {
	return c.ErrMessage
}

func NewInternalServerError(msg string) CustomError {
	return &generalError{
		ErrStatusCode: http.StatusInternalServerError,
		ErrStatus:     http.StatusText(http.StatusInternalServerError),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewNotFoundError(msg string) CustomError {
	return &generalError{
		ErrStatusCode: http.StatusNotFound,
		ErrStatus:     http.StatusText(http.StatusNotFound),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewBadRequestError(msg string) CustomError {
	return &generalError{
		ErrStatusCode: http.StatusBadRequest,
		ErrStatus:     http.StatusText(http.StatusBadRequest),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewConflictError(msg string) CustomError {
	return &generalError{
		ErrStatusCode: http.StatusConflict,
		ErrStatus:     http.StatusText(http.StatusConflict),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewUnprocessableEntityError(msg string) CustomError {
	return &generalError{
		ErrStatusCode: http.StatusUnprocessableEntity,
		ErrStatus:     http.StatusText(http.StatusUnprocessableEntity),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewUnauthorizedError(msg string) CustomError {
	return &generalError{
		ErrStatusCode: http.StatusUnauthorized,
		ErrStatus:     http.StatusText(http.StatusUnauthorized),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewForbiddenError(msg string) CustomError {
	return &generalError{
		ErrStatusCode: http.StatusForbidden,
		ErrStatus:     http.StatusText(http.StatusForbidden),
		ErrMessage:    msg,
		Data:          nil,
	}
}
