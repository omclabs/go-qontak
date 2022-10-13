package helpers

import (
	"errors"
	"net/http"
)

var errBadRequest = errors.New("bad request")
var errUnauthorized = errors.New("unauthorized request")
var errForbidden = errors.New("forbidden request")
var errNotFound = errors.New("not found")
var errMethodNotAllowed = errors.New("method not allowed")
var errUnprocessable = errors.New("unprocessable entity")
var errInternalServer = errors.New("internal server error")
var errMethodNotImplemented = errors.New("method not implemented")
var errBadGateway = errors.New("bad gateway")
var errServiceUnavailable = errors.New("service unavailable")
var errGatewayTimeout = errors.New("gateway timeout")

func GetErrBadRequest() error {
	return errBadRequest
}

func GetErrUnauthorized() error {
	return errUnauthorized
}

func GetErrForbidden() error {
	return errForbidden
}

func GetErrNotFound() error {
	return errNotFound
}

func GetErrMethodNotAllowed() error {
	return errMethodNotAllowed
}

func GetErrUnprocessable() error {
	return errUnprocessable
}

func GetErrInternalServer() error {
	return errInternalServer
}

func GetErrMethodNotImplemented() error {
	return errMethodNotImplemented
}

func GetErrBadGateway() error {
	return errBadGateway
}

func GetErrServiceUnavailable() error {
	return errServiceUnavailable
}

func GetErrGatewayTimeout() error {
	return errGatewayTimeout
}

func MapHttpStatusCode(err error) int {
	statusCode := http.StatusInternalServerError

	switch err {
	case errBadRequest:
		statusCode = http.StatusBadRequest
	case errUnauthorized:
		statusCode = http.StatusUnauthorized
	case errForbidden:
		statusCode = http.StatusForbidden
	case errNotFound:
		statusCode = http.StatusNotFound
	case errMethodNotAllowed:
		statusCode = http.StatusMethodNotAllowed
	case errUnprocessable:
		statusCode = http.StatusUnprocessableEntity
	case errMethodNotImplemented:
		statusCode = http.StatusNotImplemented
	case errBadGateway:
		statusCode = http.StatusBadGateway
	case errServiceUnavailable:
		statusCode = http.StatusServiceUnavailable
	case errGatewayTimeout:
		statusCode = http.StatusGatewayTimeout
	}

	return statusCode
}

func MapHttpErrorCode(code int) error {
	errorByCode := GetErrInternalServer()

	switch code {
	case 400:
		errorByCode = GetErrBadRequest()
	case 401:
		errorByCode = GetErrUnauthorized()
	case 403:
		errorByCode = GetErrForbidden()
	case 404:
		errorByCode = GetErrNotFound()
	case 405:
		errorByCode = GetErrMethodNotAllowed()
	case 422:
		errorByCode = GetErrUnprocessable()
	case 501:
		errorByCode = GetErrMethodNotImplemented()
	case 502:
		errorByCode = GetErrBadGateway()
	case 503:
		errorByCode = GetErrServiceUnavailable()
	case 504:
		errorByCode = GetErrGatewayTimeout()
	}
	return errorByCode
}
