package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cweill/gotests/code"
	"github.com/cweill/gotests/models"
	"github.com/cweill/gotests/render"
	"golang.org/x/tools/imports"
)

var NoTestsError = errors.New("no tests generated")

func GenerateTests(fi *models.FileInfo, onlyFuncs, exclFuncs []string) ([]string, error) {
	info := code.Parse(fi.SourcePath)
	tfs := info.TestableFuncs(onlyFuncs, exclFuncs)
	if len(tfs) == 0 {
		return nil, NoTestsError
	}
	f, err := os.Create(fi.TestPath)
	if err != nil {
		return nil, fmt.Errorf("oc.Create: %v", err)
	}
	defer f.Close()
	tests, err := writeTests(f, info, tfs)
	if err != nil {
		os.Remove(f.Name())
		return nil, err
	}
	f.Sync()
	return tests, nil
}

func writeTests(f *os.File, info *models.SourceInfo, funcs []*models.Function) ([]string, error) {
	w := bufio.NewWriter(f)
	if err := render.Header(w, info); err != nil {
		return nil, fmt.Errorf("render.Header: %v", err)
	}
	var tests []string
	for _, fun := range funcs {
		if err := render.TestFunction(w, fun); err != nil {
			return nil, fmt.Errorf("render.TestFunction: %v", err)
		}
		tests = append(tests, fmt.Sprintf("%v.%v", info.Package, fun.TestName()))
	}
	if err := w.Flush(); err != nil {
		return nil, fmt.Errorf("bufio.Flush: %v", err)
	}
	if err := processImports(f); err != nil {
		return nil, fmt.Errorf("processImports: %v", err)
	}
	return tests, nil
}

func processImports(f *os.File) error {
	v, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	b, err := imports.Process(f.Name(), v, nil)
	if err != nil {
		return fmt.Errorf("imports.Process: %v", err)
	}
	n, err := f.WriteAt(b, 0)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}
	if err := f.Truncate(int64(n)); err != nil {
		return fmt.Errorf("file.Truncate: %v", err)
	}
	return nil
}
