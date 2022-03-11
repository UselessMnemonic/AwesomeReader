package main

import "fmt"

func Catch(err error) {
	p := recover()
	if p != nil {
		err = fmt.Errorf("%v", p)
	}
}
