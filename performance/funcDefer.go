package main

import (
	"fmt"
	"time"
)

func callFunc() {
	p := time.Now().UnixNano()
	for i := 0; i < 100*1000*1000; i++ {
		function()
	}
	fmt.Println("func ", (time.Now().UnixNano()-p)/1000/1000)

	p = time.Now().UnixNano()
	for i := 0; i < 100*1000*1000; i++ {
		deferFunc()
	}
	fmt.Println("defer", (time.Now().UnixNano()-p)/1000/1000)
}

func f() {
}

func function() {
	f()
}

func deferFunc() {
	defer f()
}
