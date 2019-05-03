package main

import "fmt"

func main() {
	b := []byte("abcdefghijklmnopqrstuvwxyz")
	fmt.Println(string(b))
	copy(b, b[7:])
	fmt.Println(string(b))
}
