package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randTime() {
	p := time.Now().UnixNano()
	for i := 0; i < 100*1000*1000; i++ {
		rand.Uint32()
	}
	fmt.Println((time.Now().UnixNano() - p) / 1000 / 1000)
}
