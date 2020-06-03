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

	var s *TS
	s.Print()
}

// -----------------------------
type TT struct {
	Ver uint16
	Nid uint16
	Uid uint32
	Cid uint32
	Len uint16
	Res [34]byte
}

type TS struct {
	Ver int8
	Nid uint16
	Pid uint16
	Len uint16
	Uid uint32
	Cid uint32
	Res []byte
}

func ts() {
	fmt.Println(unsafe.Sizeof(TS{}))
	a := TS{}
	fmt.Println(unsafe.Offsetof(a.Ver))
	fmt.Println(unsafe.Offsetof(a.Nid))
	fmt.Println(unsafe.Offsetof(a.Pid))
	fmt.Println(unsafe.Offsetof(a.Len))
	fmt.Println(unsafe.Offsetof(a.Uid))
	fmt.Println(unsafe.Offsetof(a.Cid))
	fmt.Println(unsafe.Sizeof(TT{}))
}

func (ts *TS) Print() {
	if ts == nil {
		fmt.Println("no panic and ts is nil")
	}
}
