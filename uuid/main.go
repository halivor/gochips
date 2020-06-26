package main

import (
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	for _, c := range id {
		fmt.Printf("%02x\n", c)
	}
	fmt.Println(hex.EncodeToString(id[:]))
	xid, _ := uuid.Parse(id.String())
	fmt.Println(hex.EncodeToString(xid[:]))
}
