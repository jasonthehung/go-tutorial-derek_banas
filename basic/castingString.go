package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	cV3 := "50000"

	// ASCII to Integer
	cV4, err := strconv.Atoi(cV3)

	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 50000

	// Integer to ASCII
	cV6 := strconv.Itoa(cV5)

	pl(cV6, err, reflect.TypeOf(cV6))

	cV7 := "3.14"

	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8, reflect.TypeOf(cV8))
	}
}
