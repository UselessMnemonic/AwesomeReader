package main

import (
	"io"
	"reflect"
)

func readSliceAwesomely(r io.Reader, dst reflect.Value, count int, ctx AwesomeContext) error {
	for i := 0; i < count; i++ {
		if err := readAwesomely(r, dst.Index(i).Addr(), ctx); err != nil {
			return err
		}
	}
	return nil
}
