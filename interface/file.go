package main

import (
	"io"
	"os"
)

func file() {
	fp, _ := os.OpenFile("./main.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	var w io.Writer = fp
	w.Write([]byte("before close\n"))
	if fp, ok := w.(*os.File); ok {
		fp.Close()
	}
	fpn, _ := os.OpenFile("./main.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	*fp = *fpn
	w.Write([]byte("after close"))
	fp.Close()
}
