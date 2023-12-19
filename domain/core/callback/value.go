package callback

import (
	"reflect"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func OnCallbackValueString(elem app.Composer, field *string) func(val app.Value) {
	return func(val app.Value) {
		str := val.String()
		reflect.ValueOf(field).Elem().Set(reflect.ValueOf(str))
		elem.Update()
	}
}