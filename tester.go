package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"tester/code"
	"tester/render"
)

func generateTestCases(f *os.File, path string) {
	info := code.Parse(path)
	render.Header(f, info)
	render.TestCases(f, info)
	if err := exec.Command("gofmt", "-w", f.Name()).Run(); err != nil {
		fmt.Printf("error %v", err)
	}
}

func main() {
	for _, path := range []string{"testfiles/test1.go"} {
		f, err := ioutil.TempFile("", "")
		if err != nil {
			fmt.Printf("error %v", err)
			continue
		}
		generateTestCases(f, path)
	}
}
