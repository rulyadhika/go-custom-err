package errs

import "net/http"

type Error interface {
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

func NewInternalServerError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusInternalServerError,
		ErrStatus:     http.StatusText(http.StatusInternalServerError),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewNotFoundError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusNotFound,
		ErrStatus:     http.StatusText(http.StatusNotFound),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewBadRequestError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusBadRequest,
		ErrStatus:     http.StatusText(http.StatusBadRequest),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewConflictError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusConflict,
		ErrStatus:     http.StatusText(http.StatusConflict),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewUnprocessableEntityError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusUnprocessableEntity,
		ErrStatus:     http.StatusText(http.StatusUnprocessableEntity),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewUnauthorizedError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusUnauthorized,
		ErrStatus:     http.StatusText(http.StatusUnauthorized),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewForbiddenError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusForbidden,
		ErrStatus:     http.StatusText(http.StatusForbidden),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewServiceUnavailableError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusServiceUnavailable,
		ErrStatus:     http.StatusText(http.StatusServiceUnavailable),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewGatewayTimeoutError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusGatewayTimeout,
		ErrStatus:     http.StatusText(http.StatusGatewayTimeout),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewNotImplementedError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusNotImplemented,
		ErrStatus:     http.StatusText(http.StatusNotImplemented),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewBadGatewayError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusBadGateway,
		ErrStatus:     http.StatusText(http.StatusBadGateway),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewRequestTimeoutError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusRequestTimeout,
		ErrStatus:     http.StatusText(http.StatusRequestTimeout),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewTooManyRequestsError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusTooManyRequests,
		ErrStatus:     http.StatusText(http.StatusTooManyRequests),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewUnsupportedMediaTypeError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusUnsupportedMediaType,
		ErrStatus:     http.StatusText(http.StatusUnsupportedMediaType),
		ErrMessage:    msg,
		Data:          nil,
	}
}

func NewMethodNotAllowedError(msg string) Error {
	return &generalError{
		ErrStatusCode: http.StatusMethodNotAllowed,
		ErrStatus:     http.StatusText(http.StatusMethodNotAllowed),
		ErrMessage:    msg,
		Data:          nil,
	}
}
