package main

// #include "03.h"
import "C"
import "fmt"

func main() {
	a := C.int(5)
	b := C.int(7)
	fmt.Println(C.Add(a, b))
}
