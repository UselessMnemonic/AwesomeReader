package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var err error

	data := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	reader := bytes.NewReader(data)
	val := []*uint8{new(uint8), new(uint8)}

	ctx := AwesomeContext{
		DefaultOrder: binary.BigEndian,
	}
	err = ReadAwesomely(reader, &val, ctx)
	fmt.Printf("Result: %d\n", val)
	fmt.Printf("Error: %v", err)
}
