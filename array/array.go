package main

import (
	"fmt"
	"unsafe"
)

type p struct {
	len int
	cap int
	ptr uintptr
}

func main() {
	bp := make([]byte, 2048)
	b := (*p)(unsafe.Pointer(&bp[0]))
	fmt.Println(b)
	ptr := (*[32]byte)(unsafe.Pointer(&b.ptr))[:]
	copy(ptr, "abcdefg")
	fmt.Printf("%v %v", ptr, bp[8:16])
}
