package main

import (
	"fmt"
	"time"
)

type p struct {
	len int
	cap int
	ptr uintptr
}

func arrayMap() {
	m := map[int]int{1: 1}
	a := [1]int{0}

	b := time.Now().UnixNano()
	for i := 0; i < 1000*1000*1000; i++ {
		m[i/100] = 1
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
