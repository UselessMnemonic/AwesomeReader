package main

import (
	"fmt"
	"io"
	"reflect"
)

func readStringAwesomely(r io.Reader, dst reflect.Value, count uint64, ctx AwesomeContext) error {
	bytes := make([]byte, count)
	n, e := r.Read(bytes)
	if e != nil {
		return e
	}
	if uint64(n) != count {
		return fmt.Errorf("expected %d bytes, got %d", count, n)
	}
	dst.SetString(string(bytes))
	return nil
}
