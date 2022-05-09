package validator

import (
	"fmt"

	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

var _ validatoradapter.Adapter = Validator{}

func New() Validator {
	return Validator{
		validate: validator.New(),
	}
}

func (v Validator) Struct(s interface{}) error {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	reason := ErrorReason(validationErr)

	return fmt.Errorf("%w on %s, %s", adapter.ErrInvalidInput, validationErr.Field(), reason)
}

func (v Validator) Var(field validatoradapter.Field) error {
	err := v.validate.Var(field.Value, field.Tag.String())
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	reason := ErrorReason(validationErr)

	return fmt.Errorf("%w on %s, %s", adapter.ErrInvalidInput, field.Name, reason)
}

func ErrorReason(err validator.FieldError) string {
	var reason string
	switch err.Tag() {
	case "required":
		reason = "this field is required"
	case "numeric":
		reason = "this field should only contains numeric value"
	case "alpha":
		reason = "this field should only contains alphabet value"
	case "alphanum":
		reason = "this field should only contains alphanumeric value"
	case "email":
		reason = "this field should be a valid email address"
	case "url":
		reason = "this field should be a valid URL"
	case "max":
		reason = fmt.Sprintf("this field should not be longer than %s character(s)", err.Param())
	case "min":
		reason = fmt.Sprintf("this field should not be shorter than %s character(s)", err.Param())
	case "oneof":
		reason = fmt.Sprintf("this field should be one of: %s", err.Param())
	case "gt":
		reason = fmt.Sprintf("this field should be greater than %s", err.Param())
	case "gte":
		reason = fmt.Sprintf("this field should be greater than or equal %s", err.Param())
	case "lt":
		reason = fmt.Sprintf("this field should be less than %s", err.Param())
	case "lte":
		reason = fmt.Sprintf("this field should be less than or equal %s", err.Param())
	}

	return reason
}
