package main

import (
	"fmt"
	"unsafe"
)

func main() {
	sliceBegin()
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
