package core

import (
	"reflect"
	"strings"

	"github.com/gofrs/uuid"
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

func Error(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func NewUUID() *uuid.UUID {
	uid, _ := uuid.NewV4()
	return &uid
}

func Hidden(err error, elemStyles ...string) string {
	if err == nil {
		elemStyles = append(elemStyles, "hidden")
	}
	return strings.Join(elemStyles, " ")
}
