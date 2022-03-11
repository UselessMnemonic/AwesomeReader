package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

func readStructAwesomely(r io.Reader, dst reflect.Value, ctx AwesomeContext) error {
	for _, meta := range reflect.VisibleFields(dst.Type()) {
		err := readFieldAwesomely(r, dst, dst.FieldByIndex(meta.Index).Addr(), meta, ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func readFieldAwesomely(r io.Reader, parent reflect.Value, dst reflect.Value, meta reflect.StructField, ctx AwesomeContext) error {
	if order, hasOrder := meta.Tag.Lookup("order"); hasOrder {
		switch order {
		case "be":
			ctx.DefaultOrder = binary.BigEndian
		case "le":
			ctx.DefaultOrder = binary.LittleEndian
		default:
			return fmt.Errorf("invalid byte ordering (%s)", order)
		}
	}
	return readAwesomely(r, dst, ctx)
}

func parseSizeTag(parent reflect.Value, field reflect.StructField) (result uint64, err error) {
	size, hasSize := field.Tag.Lookup("size")
	sizedBy, hasSizedBy := field.Tag.Lookup("sizedBy")

	if hasSize && hasSizedBy {
		err = fmt.Errorf("both size and sizedBy specified")
		return
	}

	if hasSize {
		result, err = strconv.ParseUint(size, 10, 64)
		return
	}

	if hasSizedBy {
		defer Catch(err)
		sizeValue := parent.FieldByName(sizedBy)

		switch sizeValue.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result = sizeValue.Uint()
			return

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			signed := sizeValue.Int()
			if signed < 0 {
				err = fmt.Errorf("invalid length specified %d", signed)
				return
			}
			result = uint64(signed)
			return

		default:
			err = fmt.Errorf("length specifier %s cannot be type %s", field.Name, sizeValue.Type().Name())
			return
		}
	}

	err = fmt.Errorf("no size specified")
	return
}
