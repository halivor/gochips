package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	A int16
	B int8
	C int32
	E int8
	D int64
}

func main() {
	fmt.Println(unsafe.Sizeof(T{}))
}
