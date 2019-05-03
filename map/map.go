package main

import "fmt"

func main() {
	m := map[int]string{
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
