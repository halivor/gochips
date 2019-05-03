package main

import "fmt"

const STD_LEN = 4096

func main() {
	length := 80244
	fmt.Printf("length % 40b\t %d\n", length, length)
	fmt.Printf("%6d % 40b\n", 4095, 4095)
	fmt.Printf("length % 40b\t %d\n", (^length&(STD_LEN-1)+1)&(STD_LEN-1), (^length&(STD_LEN-1)+1)&(STD_LEN-1))
	fmt.Printf("length % 40b\t %d\n", length+(^length&(STD_LEN-1)+1)&(STD_LEN-1), length+(^length&(STD_LEN-1)+1)&(STD_LEN-1))
}
