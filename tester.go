package main

import "tester/code"

func main() {
	for _, path := range []string{"examples/ex1.go"} {
		code.Read(path)
	}
}
