package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	ints()
}

func ints() {
	pb := []byte(`{"file_id" : 100}`)
	st := map[string]interface{}{}
	json.Unmarshal(pb, &st)
	fmt.Println(st)
	to := reflect.TypeOf(st["file_id"])
	fmt.Println(to.Name(), int64(st["file_id"].(float64)))
}

func pack() {
	m := map[string]interface{}{
		"string": "hello",
		"struct": struct{}{},
		"data":   map[string]string{"hello": "world"},
	}

	n := map[string]interface{}{"key": struct{}{}, "val": struct{}{}}
	pb, _ := json.Marshal(m)
	fmt.Println(string(pb))
	pb, _ = json.Marshal(n)
	fmt.Println(string(pb))
}

func del() {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
	}

	for k, v := range m {
		if k == 3 {
			delete(m, 3)
		}
		fmt.Println(k, v)
	}

	for k, v := range m {
		if k == 2 {
			delete(m, 2)
		}
		fmt.Println(k, v)
	}
	fmt.Println(m)
}

func writer() {
	m := make(map[io.Writer]int)
	m[nil] = 1000
	m[os.Stdout] = 1
	m[os.Stderr] = 2
	fmt.Println(m[nil])
	fmt.Println(m[os.Stdout])
	fmt.Println(m[os.Stderr])
}

func pointer() {
	m := map[int]*int{
		1: new(int),
		2: new(int),
		3: new(int),
		4: new(int),
		5: new(int),
		6: new(int),
		7: new(int),
	}
	fmt.Println(*m[1])
	fmt.Println(m[2])
	fmt.Println(m[3])

}

func order() {
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
