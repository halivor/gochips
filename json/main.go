package main

import (
	"encoding/json"
	"fmt"
)

type S struct {
	I []json.RawMessage
}

type Role struct {
	RoleId    int    `json:"role_id"`
	RoleName  string `json:"role_name"`
	Style     string `json:"style"`
	FontSize  int    `json:"font_size"`
	FontColor string `json:"font_color"`
	Remark    string `json:"remark"`
}

func main() {
	pb := []byte(`[{"role_id":1,"role_name":"用户","style":"","font_size":6,"font_color":"000","remark":""}]`)
	rs := make([]*Role, 0)
	e := json.Unmarshal(pb, &rs)
	fmt.Println(e, rs[0])
}

func slice() {
	i := make([]int, 0)
	b := []byte("[1, 2, 3, 4, 5, 6, 7]")
	json.Unmarshal(b, &i)
	fmt.Println(i)
}

func row() {
	a := []byte(`[123, 4002, 2003, {"key": "value"}]`)
	var i []json.RawMessage
	json.Unmarshal(a, &i)
	for _, v := range i {
		fmt.Println(string(v))
		fmt.Println(v)
	}
}
