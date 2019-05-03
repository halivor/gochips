package main

import (
	"encoding/json"
	"fmt"
)

type S struct {
	I []json.RawMessage
}

func main() {
	a := []byte(`[123, 4002, 2003, {"key": "value"}]`)
	var i []json.RawMessage
	json.Unmarshal(a, &i)
	for _, v := range i {
		fmt.Println(string(v))
		fmt.Println(v)
	}
}
