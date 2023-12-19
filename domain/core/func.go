package core

import (
	"reflect"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

/* ptr = struct | methodTagName is get function for private property*/
func CallMethod(ptr interface{}, methodTagName string) interface{} {
	val := reflect.ValueOf(ptr)
	typ := val.Type()

	if _, ok := typ.MethodByName(methodTagName); ok {
		result := val.MethodByName(methodTagName).Call([]reflect.Value{})
		return result[0].Interface()
	}

	app.Log("struct name:", typ.Name(), " method not found with:", methodTagName)

	return nil
}
