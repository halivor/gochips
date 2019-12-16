package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	sub()
}

func sub() {
	fmt.Println(time.Now().Sub(time.Now().Add(time.Minute)) < 0)
	now := time.Now()
	fmt.Println(now)
	after := now.Sub(time.Now().Add(-time.Minute))
	fmt.Println(now, uint64(time.Minute), uint64(after))
}

func timer() {
	tm := time.NewTimer(time.Second * 5)
	tk := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-tm.C:
			log.Println("tmo")
			tm = time.NewTimer(time.Second * 5)
		case <-tk.C:
			log.Println("ticker")
		}
	}
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
