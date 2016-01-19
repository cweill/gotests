package main

import (
	"fmt"
	"tester/code"
	"tester/render"
)

type LogWriter struct {
	log []byte
}

func (l *LogWriter) Write(p []byte) (n int, err error) {
	l.log = append(l.log, p...)
	return len(p), nil
}

func main() {
	for _, path := range []string{"examples/ex1.go"} {
		w := &LogWriter{}
		info := code.Parse(path)
		render.Header(w, info)
		render.TestCases(w, info)
		fmt.Println(string(w.log))
	}
}
