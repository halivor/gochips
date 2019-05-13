package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	tmr := (now/86400+1)*86400 - 8*3600
	tm := (1550073600/86400+1)*86400 - 8*3600
	fmt.Println(now, tmr, tm)
	fmt.Println(tmr-now, tm-1550073600)
}
