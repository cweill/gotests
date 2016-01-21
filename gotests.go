package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cweill/gotests/code"
	"github.com/cweill/gotests/models"
	"github.com/cweill/gotests/render"
	"golang.org/x/tools/imports"
)

func generateTests(fi *models.FileInfo, onlyFuncs, exclFuncs []string) {
	info := code.Parse(fi.SourcePath)
	tfs := info.TestableFuncs(onlyFuncs, exclFuncs)
	if len(tfs) == 0 {
		fmt.Println("No tests generated")
		return
	}
	f, err := os.Create(fi.TestPath)
	if err != nil {
		fmt.Printf("oc.Create: %v\n", err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	if err := render.Header(w, info); err != nil {
		fmt.Printf("render.Header: %v\n", err)
		os.Remove(f.Name())
		return
	}
	var count int
	for _, fun := range tfs {
		if err := render.TestCases(w, fun); err != nil {
			fmt.Printf("render.TestCases: %v\n", err)
			continue
		}
		fmt.Printf("Generated %v.%v\n", info.Package, fun.TestName())
		count++
	}
	if err := w.Flush(); err != nil {
		fmt.Printf("bufio.Flush: %v\n", err)
		os.Remove(f.Name())
		return
	}
	if err := processImports(f); err != nil {
		fmt.Printf("processImports: %v\n", err)
	}
	if count == 0 {
		fmt.Println("No tests generated")
		os.Remove(f.Name())
	}
}

func processImports(f *os.File) error {
	v, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	b, err := imports.Process(f.Name(), v, nil)
	if err != nil {
		return fmt.Errorf("imports.Process: %v\n", err)
	}
	n, err := f.WriteAt(b, 0)
	if err != nil {
		return fmt.Errorf("file.Write: %v\n", err)
	}
	if err := f.Truncate(int64(n)); err != nil {
		return fmt.Errorf("file.Truncate: %v\n", err)
	}
	return nil
}
