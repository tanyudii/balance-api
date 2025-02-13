package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	errutil2 "github.com/tanyudii/balance-api/internal/pkg/errutil"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Remark string      `json:"remark"`
	Errors interface{} `json:"errors,omitempty"`
}

func EchoResponse(c echo.Context, code int, data interface{}) error {
	return EchoResponseWithoutData(c, code, &Response{Data: data})
}

func EchoResponseWithoutData(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, data)
}

func EchoErrorResponse(c echo.Context, err error) error {
	code := http.StatusInternalServerError
	errResp := &ErrorResponse{}

	var customErr errutil2.CustomError
	if errors.As(err, &customErr) {
		code = customErr.GetHTTPCode()
		errResp.Remark = customErr.Error()
		if errutil2.IsBadRequestError(err) {
			var badRequest *errutil2.BadRequestError
			errors.As(err, &badRequest)
			fields := badRequest.GetFields()
			if len(fields) != 0 {
				errResp.Errors = fields
			}
		}
	} else {
		errResp.Remark = http.StatusText(code)
	}

	return c.JSON(code, errResp)
}
