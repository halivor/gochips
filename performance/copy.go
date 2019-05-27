package main

import (
	"fmt"
	"time"
)

func copyByte() {
	a := make([]byte, 0)
	for i := 0; i < 2048; i++ {
		a = append(a, '0')
	}
	b := make([]byte, 2048)

	p := time.Now().UnixNano()
	for i := 0; i < 100*1000*1000; i++ {
		copy(b, a)
	}
	fmt.Println((time.Now().UnixNano() - p) / 1000 / 1000)

}
