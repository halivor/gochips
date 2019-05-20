package main

import (
	"fmt"
	"time"
	"unsafe"
)

type p struct {
	len int
	cap int
	ptr uintptr
}

func main() {
	performance()
}

func performance() {
	m := map[int]int{1: 1}
	a := [1]int{0}

	b := time.Now().UnixNano()
	for i := 0; i < 1000*1000*1000; i++ {
		m[1] = 1
	}
	e := time.Now().UnixNano()
	fmt.Println(e, b, e-b, (e-b)/(1000*1000)%(1000*1000))

	b = e
	for i := 0; i < 1000*1000*1000; i++ {
		a[0] = 0
	}
	e = time.Now().UnixNano()
	fmt.Println(e, b, e-b, (e-b)/(1000*1000)%(1000*1000))
}

func pointer() {
	bp := make([]byte, 2048)
	b := (*p)(unsafe.Pointer(&bp[0]))
	fmt.Println(b)
	ptr := (*[32]byte)(unsafe.Pointer(&b.ptr))[:]
	copy(ptr, "abcdefg")
	fmt.Printf("%v %v", ptr, bp[8:16])
}
