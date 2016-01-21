package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cweill/gotests/code"
	"github.com/cweill/gotests/render"
)

func generateTestCases(f *os.File, path string) {
	info := code.Parse(path)
	if err := render.Header(f, info); err != nil {
		fmt.Printf("render.Header: %v\n", err)
		return
	}
	if err := render.TestCases(f, info); err != nil {
		fmt.Printf("render.TestCases: %v\n", err)
		return
	}
	if err := exec.Command("gofmt", "-w", f.Name()).Run(); err != nil {
		fmt.Printf("exec.Command: %v\n", err)
	}
}

func main() {
	for _, path := range os.Args[1:] {
		for _, src := range sourceFiles(path) {
			testPath := strings.Replace(src, ".go", "_test.go", -1)
			f, err := os.Create(testPath)
			if err != nil {
				fmt.Printf("oc.Create: %v\n", err)
				continue
			}
			defer f.Close()
			generateTestCases(f, src)
		}
	}
}

func sourceFiles(path string) []string {
	var srcs []string
	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("filepath.Abs: %v\n", err)
		return nil
	}
	if filepath.Ext(path) == "" {
		ps, err := filepath.Glob(path + "/*.go")
		if err != nil {
			fmt.Printf("filepath.Glob: %v\n", err)
			return nil
		}
		for _, p := range ps {
			if !isTestFile(p) {
				srcs = append(srcs, p)
			}
		}
	} else if filepath.Ext(path) == ".go" {
		if !isTestFile(path) {
			srcs = append(srcs, path)
		}
	}
	return srcs
}

func isTestFile(path string) bool {
	ok, err := filepath.Match("*_test.go", path)
	if err != nil {
		fmt.Printf("filepath.Match: %v\n", err)
		return false
	}
	if ok {
		return true
	}
	return false
}
