package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	cnf [1]byte
}

type buffer struct {
	Len uint16
	Cap uint16
	ref uint32
	bp  *int32
	ptr uintptr
}

func main() {
	t := &T{}
	fmt.Printf("%p\n", t)
	fmt.Println(&t.cnf[0])
	fmt.Println(unsafe.Sizeof(buffer{}))
	b := &buffer{}
	fmt.Println(unsafe.Sizeof(b.ptr))
}
