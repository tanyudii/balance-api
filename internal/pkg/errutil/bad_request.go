package errutil

import (
	"errors"
	"fmt"
	"github.com/jinzhu/inflection"
	"net/http"
)

type BadRequestError struct {
	remark   string
	httpCode int
	fields   ErrorField
}

func (i *BadRequestError) Error() string {
	return i.remark
}

func (i *BadRequestError) GetHTTPCode() int {
	return i.httpCode
}

func (i *BadRequestError) GetFields() ErrorField {
	return i.fields
}

func NewBadRequestError(msg string) error {
	return &BadRequestError{
		remark:   msg,
		httpCode: http.StatusBadRequest,
	}
}

func NewBadRequestErrorWithFields(msg string, fields ErrorField) error {
	return &BadRequestError{
		remark:   msg,
		fields:   fields,
		httpCode: http.StatusBadRequest,
	}
}

func BadRequestOrNil(fields ErrorField) error {
	if len(fields) != 0 {
		firstErr, otherErr := fields.GetFirstErrorAndOtherTotal()
		wordErr := "error"
		if otherErr == 0 {
			return NewBadRequestErrorWithFields(firstErr, fields)
		}
		if otherErr > 1 {
			wordErr = inflection.Plural("error")
		}
		return NewBadRequestErrorWithFields(fmt.Sprintf("%s. and there are %d %s", firstErr, otherErr, wordErr), fields)
	}
	return nil
}

func IsBadRequestError(err error) bool {
	var expectedErr *BadRequestError
	return errors.As(err, &expectedErr)
}
