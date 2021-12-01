package main

import (
	"os"
	"reflect"
)

type Function struct{}

func main() {
	puzzle := os.Args[1]
	var function Function
	reflect.ValueOf(&function).MethodByName(puzzle).Call([]reflect.Value{})
}
