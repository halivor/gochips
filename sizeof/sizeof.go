package main

import (
	"fmt"
	"unsafe"
)

type TT struct {
	A int16
	B int8
	C int32
	E int8
	D int64
}

type B struct {
	i uint32
	u uint64
}

type T struct {
	Buf []byte
	ori []byte
	bp  *B
	Pos uint16
	ref uint16
	arr [988]byte
}

func main() {
	var ch chan int
	fmt.Println("ch =>", unsafe.Sizeof(ch))
	u := uint16(0)
	fmt.Println("uint16 =>", unsafe.Sizeof(u))
	fmt.Println("uint16 =>", unsafe.Sizeof(u))
	fmt.Println("pointer =>", unsafe.Sizeof(&B{}))
	fmt.Println("[]byte =>", unsafe.Sizeof([]byte{}))
	fmt.Println("array =>", unsafe.Sizeof([988]byte{}))
	fmt.Println("T{} => ", unsafe.Sizeof(T{}))
}

func t() {
	fmt.Println(unsafe.Sizeof(TT{}))
}
