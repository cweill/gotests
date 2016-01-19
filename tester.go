package main

import (
	"fmt"
	"io"
	"tester/code"
	"tester/render"
)

type logWriter struct {
	log []byte
}

func (l *logWriter) Write(p []byte) (n int, err error) {
	l.log = append(l.log, p...)
	return len(p), nil
}

func generateTestCases(w io.Writer, path string) {
	info := code.Parse(path)
	render.Header(w, info)
	render.TestCases(w, info)
}

func main() {
	for _, path := range []string{"examples/ex1.go"} {
		w := &logWriter{}
		generateTestCases(w, path)
		fmt.Println(string(w.log))
	}
}
