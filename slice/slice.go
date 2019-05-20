package main

import (
	"fmt"
	"unsafe"
)

func main() {
	sliceLen()
}

func sliceLen() {
	b := make([]byte, 1024, 2048)
	fmt.Println("b := make([]byte, 1024, 2048), len", len(b), ", cap", cap(b))
	b1 := b[0:0:1024]
	fmt.Println("lb := b[0:0:1024],             len", len(b1), ", cap", cap(b1))
	b2 := b[0:1024:1024]
	fmt.Println("lb := b[0:1024:1024],          len", len(b2), ", cap", cap(b2))
}

func sliceBegin() {
	var arr [1024]unsafe.Pointer
	fmt.Println(unsafe.Pointer(&arr[0]))
	as := arr[:]
	as = as[856:]
	fmt.Println(cap(arr), cap(as))
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(&as[0])) - uintptr(cap(arr)-cap(as))*8))
	as = nil
	fmt.Println(len(as))
}

func sliceEqual() {
	var arr [128]byte
	fmt.Println(arr[1])
	as := arr[:]
	fmt.Println("before", unsafe.Pointer(&as))
	as = as[1:]
	fmt.Println("after ", unsafe.Pointer(&as))
}
