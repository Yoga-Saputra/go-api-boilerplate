package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator defines the validator config
type Validator struct {
	engine *validator.Validate
}

// NewValidator creates new validator instances
func NewValidator() *Validator {
	return &Validator{
		engine: validator.New(),
	}
}

// Validate validates given struct and return error
// Otherwise return nil if pass the validate
func (v *Validator) Validate(i interface{}) error {
	var errs []string
	err := v.engine.Struct(i)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errTag := ""
			val := fmt.Sprintf("%v", err.Value())
			kind := fmt.Sprintf("%v", err.Type())

			switch true {
			case err.ActualTag() == "required":
				errTag = "is required"
			case err.ActualTag() == "min" && kind == "string":
				errTag = fmt.Sprintf("is minimum %v character", err.Param())
			case err.ActualTag() == "max" && kind == "string":
				errTag = fmt.Sprintf("is maximum %v character", err.Param())
			case err.ActualTag() == "min":
				errTag = fmt.Sprintf("must be greater then %v", val)
			case err.ActualTag() == "max":
				errTag = fmt.Sprintf("must be less then %v", val)
			default:
				errTag = "not match with any rules"
			}

			payloadError := fmt.Sprintf("Value of payload %s %s", err.Field(), errTag)
			errs = append(errs, payloadError)

			if err.ActualTag() == "required" && val == "0" {
				break
			}
		}

		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}
