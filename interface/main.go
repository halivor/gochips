package main

import "fmt"

type T struct {
	a int
}

func main() {
	file()
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
