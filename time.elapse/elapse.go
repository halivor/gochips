package main

import (
	"fmt"
	"time"
)

func main() {
	var i int64
	m := make(map[int64]int64, 1000*1000)
	fmt.Println("begin", time.Now().Format("2006-01-02 15:04:05.000000000"))
	for i = 0; i < 1000*1000; i++ {
		m[i] = time.Now().Unix()
	}
	fmt.Println("end  ", time.Now().Format("2006-01-02 15:04:05.000000000"))
}
