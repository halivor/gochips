package main

import (
	"log"
	"runtime"
	"sync"
	//"sync/atomic"
	"time"
)

const (
	cn = 100
)

var cnts uint64 = 0
var wg sync.WaitGroup

func main() {
	wg.Add(cn)

	log.SetFlags(log.Lmicroseconds)
	for i := 0; i < cn; i++ {
		go work()
	}

	time.Sleep(time.Second)
	log.Println("begin")
	wg.Wait()
	log.Println("done")

	// Wait to see the global run queue deplete.
	time.Sleep(3 * time.Second)
}

func work() {
	time.Sleep(time.Second)

	for i := 0; i < 10*1000*1000; i++ {
		//atomic.AddUint64(&cnts, 1)
		runtime.Gosched()
	}

	wg.Done()
}
