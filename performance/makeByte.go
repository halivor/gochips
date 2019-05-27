package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeByte() {
	fixedMake()
	randMake()
}

func fixedMake() {
	b := make([]byte, 1)
	p := time.Now().UnixNano()
	for i := 0; i < 10*1000*1000; i++ {
		b = make([]byte, 2048)
	}
	fmt.Println((time.Now().UnixNano() - p) / 1000 / 1000)
	fmt.Println("fixed", b[:0])
}

func randMake() {
	b := make([]byte, 1)
	p := time.Now().UnixNano()
	for i := 0; i < 10*1000*1000; i++ {
		b = make([]byte, rand.Uint32()%1000)
	}
	fmt.Println((time.Now().UnixNano() - p) / 1000 / 1000)
	fmt.Println("random", b[:0])
}
