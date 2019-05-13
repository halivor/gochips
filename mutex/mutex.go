package main

import (
	"log"
	"sync"
)

type ts struct {
	sync.Mutex
}

var s *ts = nil

func main() {
	log.Println("")
	s = &ts{}
	s.Lock()
	defer s.Unlock()
	func() {
		s = &ts{}
	}()
	s.Lock()
	defer s.Unlock()
}
