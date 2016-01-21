package main

import (
	"errors"
	"fmt"

	"github.com/cweill/gotests/code"
	"github.com/cweill/gotests/output"
)

var NoTestsError = errors.New("no tests generated")

func generateTests(srcPath, destPath string, onlyFuncs, exclFuncs []string) ([]string, error) {
	info, err := code.Parse(srcPath)
	if err != nil {
		return nil, fmt.Errorf("code.Parse: %v", err)
	}
	funcs := info.TestableFuncs(onlyFuncs, exclFuncs)
	if len(funcs) == 0 {
		return nil, NoTestsError
	}
	return output.Write(srcPath, destPath, info.Header, funcs)
}
