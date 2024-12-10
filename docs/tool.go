package docs

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=cfg.yaml api.yaml

import (
	"fmt"
	"reflect"

	"github.com/labstack/echo/v4"
)

func BindPath(ctx echo.Context, paramName string, dest any) error {
	// Everything comes in by pointer, dereference it
	v := reflect.Indirect(reflect.ValueOf(dest))

	// This is the basic type of the destination object.
	t := v.Type()

	if t.Kind() == reflect.String {
		return echo.PathParamsBinder(ctx).String(paramName, dest.(*string)).BindError()
	} else if bu, ok := dest.(echo.BindUnmarshaler); ok {
		return echo.PathParamsBinder(ctx).BindUnmarshaler(paramName, bu).BindError()
	}

	return fmt.Errorf("unsupported type %T", t.Kind())
}
