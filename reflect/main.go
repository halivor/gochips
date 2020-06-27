package main

import (
	"fmt"
	"reflect"
)

type img struct {
	Image string
}

func main() {
	is := []*img{&img{"a"}, &img{"b"}, &img{"c"}}
	getImage(is)
}

func getImage(is interface{}) {
	if reflect.ValueOf(is).Kind() != reflect.Slice {
		return
	}
	slen := reflect.ValueOf(is).Len()
	for i := 0; i < slen; i++ {
		isd, _ := reflect.Indirect(reflect.ValueOf(is).Index(i)).
			FieldByName("Image").Interface().(string)
		fmt.Println(reflect.ValueOf(is).Index(i), isd)
	}
	return
}

func pValue() {
	i := &img{"d"}
	T := reflect.TypeOf(i)
	V := reflect.ValueOf(i)

	fmt.Println(T, V, reflect.Indirect(V).FieldByName("Image").Interface())
}

func call() {
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
