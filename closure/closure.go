package main

import (
	"fmt"
	"unsafe"
)

func main() {
	t := test()
	fmt.Println("main test", t, unsafe.Pointer(&t[0]))
	t1 := test1()
	fmt.Println("main test1", t1, unsafe.Pointer(&t1[0]))
}

func test() (t [3]int) {
	i := 0
	t = [3]int{0, 1, 2}
	fmt.Println("top  ", t)

	defer func() {
		t[0] = 10
		fmt.Println("defer", unsafe.Pointer(&i), unsafe.Pointer(&t[0]), t)
	}()
	fmt.Println("out  ", unsafe.Pointer(&i), unsafe.Pointer(&t[0]), t)
	return t
}

func test1() [3]int {
	i := 0
	t := [3]int{0, 1, 2}
	fmt.Println("top  ", t)

	defer func() {
		t[0] = 10
		fmt.Println("defer", unsafe.Pointer(&i), unsafe.Pointer(&t[0]), t)
	}()
	fmt.Println("out  ", unsafe.Pointer(&i), unsafe.Pointer(&t[0]), t)
	return t
}
