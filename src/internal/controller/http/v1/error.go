package v1

import (
	"errors"
	"fmt"

	"github.com/labstack/echo"
)

var (
	ErrCannotParseToken  = fmt.Errorf("cannot parse API key")
	ErrInvalidAuthHeader = fmt.Errorf("invalid auth token")
)

func newErrorResponse(c echo.Context, errStatus int, message string) {
	err := errors.New(message)
	var HTTPError *echo.HTTPError
	ok := errors.As(err, &HTTPError)
	if !ok {
		report := echo.NewHTTPError(errStatus, err.Error())
		_ = c.JSON(errStatus, report)
	}
	c.Error(errors.New("internal server error"))
}