package main

import (
	"fmt"
	"net"
	"os"
	_ "runtime"
	"strconv"
	"time"
)

func main() {
	fmt.Println("remote address ", os.Args[1], ", coroutine", os.Args[2])
	num, _ := strconv.Atoi(os.Args[2])
	length, _ := strconv.Atoi(os.Args[3])
	data := make([]byte, 0)
	for i := 0; i < length; i++ {
		data = append(data, 'o')
	}
	for i := 1; i < num; i++ {
		go send(i, os.Args[1], data)
	}
	send(num, os.Args[1], data)
}

func send(no int, ip string, data []byte) {
	num := 0
	addr, _ := net.ResolveUDPAddr("udp", ip)
	fd, _ := net.DialUDP("udp", nil, addr)
	go func() {
		for {
			n := num
			time.Sleep(time.Second * 1)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), no, num-n)
		}
	}()
	for {
		fd.Write(data)
		num++
	}
}
