package errutil

import (
	"errors"
	"net/http"
)

type NotFoundError struct {
	remark   string
	httpCode int
}

func (i *NotFoundError) Error() string {
	return i.remark
}

func (i *NotFoundError) GetHTTPCode() int {
	return i.httpCode
}

func NewNotFoundError(msg string) error {
	return &NotFoundError{
		remark:   msg,
		httpCode: http.StatusNotFound,
	}
}

func IsNotFoundError(err error) bool {
	var expectedErr *NotFoundError
	return errors.As(err, &expectedErr)
}
