package errutil

import (
	"errors"
	"net/http"
)

type InternalServerError struct {
	remark   string
	httpCode int
}

func (i *InternalServerError) Error() string {
	return i.remark
}

func (i *InternalServerError) GetHTTPCode() int {
	return i.httpCode
}

func NewInternalServerError(msg string) error {
	return &InternalServerError{
		remark:   msg,
		httpCode: http.StatusInternalServerError,
	}
}

func IsInternalServerError(err error) bool {
	var expectedErr *InternalServerError
	return errors.As(err, &expectedErr)
}
