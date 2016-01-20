package main

import (
	"fmt"
	"gentests/code"
	"gentests/render"
	"io/ioutil"
	"os"
	"os/exec"
)

func generateTestCases(f *os.File, path string) {
	info := code.Parse(path)
	if err := render.Header(f, info); err != nil {
		fmt.Printf("error %v", err)
		return
	}
	if err := render.TestCases(f, info); err != nil {
		fmt.Printf("error %v", err)
		return
	}
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
		defer f.Close()
		generateTestCases(f, path)
	}
}
