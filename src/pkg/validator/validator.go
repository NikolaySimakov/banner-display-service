package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	v *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	cv := &CustomValidator{v: v}

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return cv
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.v.Struct(i)
	if err != nil {
		fieldErr := err.(validator.ValidationErrors)[0]
		return cv.newValidationError(fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
	}
	return nil
}

func (cv *CustomValidator) newValidationError(field string, tag string, param string) error {
	switch tag {
	case "required":
		return fmt.Errorf("field %s is required", field)
	case "max":
		return fmt.Errorf("field %s must be at most %s characters", field, param)
	default:
		return fmt.Errorf("field %s is invalid", field)
	}
}