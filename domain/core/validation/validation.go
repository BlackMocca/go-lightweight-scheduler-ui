package validation

import (
	validator "github.com/go-ozzo/ozzo-validation/v4"
	rule "github.com/go-ozzo/ozzo-validation/v4/is"
)

type ValidateRule func(val interface{}) error

const (
	requiredErr = "Must be Required"
)

var (
	Required ValidateRule = func(val interface{}) error {
		return validator.Validate(val, validator.Required.Error(requiredErr))
	}
	URL ValidateRule = func(val interface{}) error {
		return validator.Validate(val, rule.URL)
	}
)

func Validate(val interface{}, rules ...ValidateRule) error {
	if len(rules) > 0 {
		for _, fn := range rules {
			if err := fn(val); err != nil {
				return err
			}
		}
	}

	return nil
}
