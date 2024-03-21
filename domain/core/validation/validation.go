package validation

import (
	"encoding/json"
	"errors"

	validator "github.com/go-ozzo/ozzo-validation/v4"
	rule "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/spf13/cast"
)

type ValidateRule func(val interface{}) error

const (
	requiredErr = "Must be Required"
	urlErr      = "Must be a valid URL"
)

var (
	Required ValidateRule = func(val interface{}) error {
		return validator.Validate(val, validator.Required.Error(requiredErr))
	}
	Selected ValidateRule = func(val interface{}) error {
		return validator.Validate(val, validator.Min(0))
	}
	URL ValidateRule = func(val interface{}) error {
		return validator.Validate(val, rule.URL.Error(urlErr))
	}
	Number ValidateRule = func(val interface{}) error {
		return validator.Validate(val, validator.By(func(value interface{}) error {
			_, err := cast.ToIntE(value)
			return err
		}))
	}
	Bool ValidateRule = func(val interface{}) error {
		return validator.Validate(val, validator.By(func(value interface{}) error {
			_, err := cast.ToBoolE(value)
			return err
		}))
	}
	Json ValidateRule = func(val interface{}) error {
		return validator.Validate(val, validator.By(func(value interface{}) error {
			v := cast.ToString(value)
			if v == "" {
				return nil
			}
			m := map[string]interface{}{}
			if err := json.Unmarshal([]byte(v), &m); err != nil {
				return err
			}
			if len(m) == 0 && v != "" {
				return errors.New("must be json string")
			}
			return nil
		}))
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
