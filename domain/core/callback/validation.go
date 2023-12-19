package callback

import (
	"reflect"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func setValueError(field *error, err error) {
	if err != nil {
		reflect.ValueOf(field).Elem().Set(reflect.ValueOf(err))
		return
	}
	*field = nil
}

func OnValidateCallback(elem app.Composer, field *error) func(err error) {
	return func(err error) {
		setValueError(field, err)
		elem.Update()
	}
}
