package main

import "fmt"

const (
	ONE = 1 << iota
	TWO
	THREE
	FOUR

	A = 1 + iota
	B
	C
	D
)

func main() {
	fmt.Println(ONE, TWO, THREE, FOUR)
	fmt.Println(A, B, C, D)
}
