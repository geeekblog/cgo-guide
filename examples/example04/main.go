package main

// #cgo CFLAGS: -I./include01/header
// #cgo LDFLAGS: -L/Users/logan/Developer/GOMOD/cgo-guide/examples/example04/include01/source
// #include "check.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.check())
}
