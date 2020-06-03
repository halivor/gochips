package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := New()

	mm := map[string]interface{}{"method": m}

	mi := mm["method"]
	ts := reflect.ValueOf(mi)

	ms := ts.MethodByName("Set")
	vals := ms.Call([]reflect.Value{reflect.ValueOf("name")})
	for _, val := range vals {
		fmt.Println(val)
	}
	mg := ts.MethodByName("Get")
	fmt.Println(mg.Call([]reflect.Value{reflect.ValueOf("name")})[0])
}
