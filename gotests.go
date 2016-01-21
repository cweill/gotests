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

func GenerateTests(srcPath, destPath string, onlyFuncs, exclFuncs []string) ([]string, error) {
	info, err := code.Parse(srcPath)
	if err != nil {
		return nil, fmt.Errorf("code.Parse: %v", err)
	}
	tfs := info.TestableFuncs(onlyFuncs, exclFuncs)
	if len(tfs) == 0 {
		return nil, NoTestsError
	}
	tf, err := ioutil.TempFile("", "gotests_")
	if err != nil {
		return nil, fmt.Errorf("ioutil.TempFile: %v", err)
	}
	defer os.Remove(tf.Name())
	tests, err := writeTestsToTemp(tf, info, tfs)
	if err != nil {
		return nil, err
	}
	df, err := os.Create(destPath)
	if err != nil {
		return nil, fmt.Errorf("os.Create: %v", err)
	}
	defer df.Close()
	if err := copyTempToDest(tf, df); err != nil {
		os.Remove(df.Name())
		return nil, err
	}
	return tests, nil
}

func writeTestsToTemp(temp *os.File, info *models.SourceInfo, funcs []*models.Function) ([]string, error) {
	w := bufio.NewWriter(temp)
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
	if err := processImports(temp); err != nil {
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
	return overwriteFile(f, b)
}

func copyTempToDest(tempf, destf *os.File) error {
	b, err := ioutil.ReadFile(tempf.Name())
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	return overwriteFile(destf, b)
}

func overwriteFile(f *os.File, b []byte) error {
	n, err := f.WriteAt(b, 0)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}
	if err := f.Truncate(int64(n)); err != nil {
		return fmt.Errorf("file.Truncate: %v", err)
	}
	return f.Sync()
}
