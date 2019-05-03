package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		switch {
		case i < 10:
			fmt.Println(10)
		case i < 100:
			fmt.Println(100)
		case i < 1000:
			fmt.Println(1000)
		default:
			fmt.Println("default")
		}
	}
}
