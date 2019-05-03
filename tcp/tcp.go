package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	td, _ := net.ResolveTCPAddr("tcp", "62.234.130.115:10301")
	fmt.Println(td.IP.To4())
	addr := syscall.SockaddrInet4{Port: td.Port}
	copy(addr.Addr[:], td.IP)
	fmt.Println(addr)
}
