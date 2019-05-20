package main

import (
	"fmt"
	"time"
)

func main() {
	ticker()
}

func ticker() {
	tn := time.NewTicker(time.Second)
	for {
		select {
		case <-tn.C:
			fmt.Println("ticker")
		}
	}
}

func tomorrow() {
	now := time.Now().Unix()
	tmr := (now/86400+1)*86400 - 8*3600
	tm := (1550073600/86400+1)*86400 - 8*3600
	fmt.Println(now, tmr, tm)
	fmt.Println(tmr-now, tm-1550073600)
}
