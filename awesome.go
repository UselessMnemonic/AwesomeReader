package main

import (
	"encoding/binary"
	"io"
	"reflect"
)

// An AwesomeReader can perform
type AwesomeReader func(io.Reader, interface{}, AwesomeContext) error

type AwesomeReading interface {
	ReadAwesomely(io.Reader, AwesomeContext) error
}

type AwesomeContext struct {
	DefaultOrder binary.ByteOrder
	Readers      map[reflect.Type]AwesomeReader
}

// ReadAwesomely reads in binary data from a reader and deserializes the data into the format
// specified by dst. Supported types may fall into either the "scalar" category, or the "reference" category.
//
// Scalar types include:
//
// - Signed and unsigned fixed-width integer types
//
// - Floating point and complex types
//
// - Structs containing supported data types
//
// Reference types include:
//
// - Types that implement the AwesomeReading interface
//
// - Arrays of supported data types
//
// - Slices of supported data types
//
// - Pointers to supported data types
//
// Nil references are unsupported.
func ReadAwesomely(r io.Reader, dst interface{}, ctx AwesomeContext) error {

	// data type implements the AwesomeReading interface
	// let it control its own deserialization
	if ar, ok := dst.(AwesomeReading); ok {
		return ar.ReadAwesomely(r, ctx)
	}

	// we support slices of and ptrs to supported data types
	switch value := reflect.ValueOf(dst); value.Kind() {
	case reflect.Slice:
		return readSliceAwesomely(r, value, value.Len(), ctx)
	case reflect.Ptr:
		return readAwesomely(r, value, ctx)
	default:
		return newInvalidDataTypeError(value.Type())
	}
}

// Root of deserialization tree, which is heavily reflective.
//
// dst must represent a pointer to a: pointer, struct, array, or slice
func readAwesomely(r io.Reader, dst reflect.Value, ctx AwesomeContext) error {

	// decoding a type that implements Awesome
	if ar, ok := dst.Interface().(AwesomeReading); ok {
		return ar.ReadAwesomely(r, ctx)
	}

	// decoding a ptr to a supported data type
	switch elem := dst.Elem(); elem.Kind() {
	case reflect.Ptr:
		return readAwesomely(r, elem, ctx)
	case reflect.Struct:
		return readStructAwesomely(r, elem, ctx)
	case reflect.Array:
		return readArrayAwesomely(r, elem, elem.Len(), ctx)
	case reflect.Slice:
		return readSliceAwesomely(r, elem, elem.Len(), ctx)
	default:
		return binary.Read(r, ctx.DefaultOrder, dst.Interface())
	}
}
