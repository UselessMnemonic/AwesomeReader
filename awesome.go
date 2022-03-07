package main

import (
	"errors"
	"io"
	"reflect"
)

type AwesomeReader interface {
	ReadAwesomely(io.Reader) error
}

func ReadAwesomely(r io.Reader, dst interface{}) error {
	if dst == nil {
		return errors.New("destination cannot not be nil")
	}
	if v, ok := dst.(AwesomeReader); ok {
		return v.ReadAwesomely(r)
	}

	v := reflect.ValueOf(dst)
	if v.Kind() == reflect.Struct
	if v.Kind() != reflect.Ptr {
		return errors.New("destination must be a pointer")
	}

	v = v.Elem()
	if v.Kind() == reflect.Struct {
		return readStructAwesomely(r, v)
	}
	if v.Kind() == reflect.Array {
		return readArrayAwesomely(r, v)
	}
	if v.Kind() == reflect.Slice {
		return readSliceAwesomely(r, v)
	}
}

func main() {
	var u struct{ p int }
	Read(nil, u)
}
