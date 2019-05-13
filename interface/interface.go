package main

import "fmt"

type T struct {
	a int
}

var m map[int][]interface{}

func main() {
	i := 1
	m = make(map[int][]interface{})
	si := make([]interface{}, 0)
	t := &T{}
	m[i] = append(si, t)
	equal(t)
}

func equal(t *T) {
	si := m[1]
	for _, k := range si {
		if t == k {
			fmt.Println(true)
		}
	}
}

func equal_nil() {
	if test() == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("non nil")
	}
}

func test() interface{} {
	return 1
}

// ---------------------------------
type IT interface {
	Nil()
}

type T1 struct {
}

func (t *T1) Nil() {
}

func t1() {
	i := New()
	if i != nil {
		fmt.Println("i != nil")
	}
	n := NewNil()
	if n == nil {
		fmt.Println("n = nil")
	}
}

func New() IT {
	return &T1{}
}

func NewNil() IT {
	return nil
}
