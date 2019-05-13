package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

func main() {
	var num int64 = 0
	addr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:8864")
	l, _ := net.ListenUDP("udp", addr)
	go func() {
		for {
			n := atomic.LoadInt64(&num)
			time.Sleep(time.Second * 1)
			fmt.Println("recv", num-n, "packets")
		}
	}()
	buff := make([]byte, 1500)
	for {
		l.Read(buff)
		atomic.AddInt64(&num, 1)
	}
}
