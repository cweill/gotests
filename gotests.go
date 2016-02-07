package gotests

import (
	"fmt"
	"go/importer"
	"go/types"
	"path"
	"regexp"

	"github.com/cweill/gotests/input"
	"github.com/cweill/gotests/internal/goparser"
	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/output"
)

type Options struct {
	Only        *regexp.Regexp
	Excl        *regexp.Regexp
	PrintInputs bool
	Write       bool
	Importer    types.Importer
}

func GenerateTests(srcPath, testPath, destPath string, opt *Options) ([]*models.Function, []byte, error) {
	files, err := input.Files(path.Dir(srcPath))
	if err != nil {
		return nil, nil, fmt.Errorf("input.Files: %v", err)
	}
	if opt.Importer == nil {
		opt.Importer = importer.Default()
	}
	p := goparser.Parser{Importer: opt.Importer}
	srcInfo, err := p.Parse(srcPath, files)
	if err != nil {
		return nil, nil, fmt.Errorf("Parser.Parse: %v", err)
	}
	header := srcInfo.Header
	var testFuncs []string
	if models.Path(testPath).IsTestPath() && output.IsFileExist(testPath) {
		testInfo, err := p.Parse(testPath, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("Parser.Parse: %v", err)
		}
		for _, fun := range testInfo.Funcs {
			testFuncs = append(testFuncs, fun.Name)
		}
		h, err := goparser.ParseHeader(srcPath, testPath)
		if err != nil {
			return nil, nil, fmt.Errorf("goparser.ParseHeader: %v", err)
		}
		header = h
	}
	funcs := srcInfo.TestableFuncs(opt.Only, opt.Excl, testFuncs)
	if len(funcs) == 0 {
		return nil, nil, nil
	}
	b, err := output.Process(header, funcs, &output.Options{
		PrintInputs: opt.PrintInputs,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("output.Process: %v", err)
	}
	if opt.Write {
		if err := output.Write(destPath, b); err != nil {
			return nil, nil, err
		}
	}
	return funcs, b, nil
}
