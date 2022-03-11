package main

import (
	"fmt"
	"reflect"
)

func newInvalidDataTypeError(t reflect.Type) error {
	return fmt.Errorf("invalid type %s", t.String())
}
