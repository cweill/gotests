package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/cweill/gotests/goparser"
	"github.com/cweill/gotests/input"
	"github.com/cweill/gotests/models"
	"github.com/cweill/gotests/output"
)

type funcs []string

func (f *funcs) String() string {
	return fmt.Sprint(*f)
}

func (f *funcs) Set(value string) error {
	if len(*f) > 0 {
		return errors.New("flag already set")
	}
	for _, fun := range strings.Split(value, ",") {
		*f = append(*f, fun)
	}
	return nil
}

var (
	onlyFuncs, exclFuncs funcs
	allFuncs             = flag.Bool("all", false, "generate tests for all functions in specified files or directories")
	writeOutput          = flag.Bool("w", false, "write result to (test) file instead of stdout")
)

func main() {
	flag.Var(&onlyFuncs, "only", "comma-separated list of case-sensitive function names for which tests will be generating exclusively. Takes precedence over -all")
	flag.Var(&exclFuncs, "excl", "comma-separated list of case-sensitive function names to exclude when generating tests. Take precedence over -only and -all")
	flag.Parse()
	if len(onlyFuncs) == 0 && len(exclFuncs) == 0 && !*allFuncs {
		fmt.Println("Please specify either the -only, -excl, or -all flag")
		return
	}
	if len(flag.Args()) == 0 {
		fmt.Println("Please specify a file or directory containing the source")
		return
	}
	var count int
	for _, path := range flag.Args() {
		ps, err := input.Files(path)
		if err != nil {
			if err == input.NoFilesFound {
				fmt.Printf("No source files found at %v\n", path)
			} else {
				fmt.Println(err.Error())
			}
			continue
		}
		for _, src := range ps {
			tests, b, err := generateTests(string(src), src.TestPath(), src.TestPath(), &options{
				only:  onlyFuncs,
				excl:  exclFuncs,
				write: *writeOutput,
			})
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			if len(tests) == 0 {
				continue
			}
			for _, test := range tests {
				fmt.Printf("Generated %v\n", test.TestName())
				count++
			}
			if !*writeOutput {
				fmt.Println(string(b))
			}
		}
	}
	if count == 0 {
		fmt.Println("No tests generated")
	}
}

type options struct {
	only  []string
	excl  []string
	write bool
}

func generateTests(srcPath, testPath, destPath string, opt *options) ([]*models.Function, []byte, error) {
	srcInfo, err := goparser.Parse(srcPath)
	if err != nil {
		return nil, nil, fmt.Errorf("goparser.Parse: %v", err)
	}
	header := srcInfo.Header
	if models.Path(testPath).IsTestPath() && output.IsFileExist(testPath) {
		testInfo, err := goparser.Parse(testPath)
		if err != nil {
			return nil, nil, fmt.Errorf("goparser.Parse: %v", err)
		}
		for _, fun := range testInfo.Funcs {
			opt.excl = append(opt.excl, fun.Name)
		}
		h, err := goparser.ParseHeader(srcPath, testPath)
		if err != nil {
			return nil, nil, fmt.Errorf("goparser.ParseHeader: %v", err)
		}
		header = h
	}
	funcs := srcInfo.TestableFuncs(opt.only, opt.excl)
	if len(funcs) == 0 {
		return nil, nil, nil
	}
	b, err := output.Process(header, funcs)
	if err != nil {
		return nil, nil, fmt.Errorf("output.Process: %v", err)
	}
	if opt.write {
		if err := output.Write(destPath, b); err != nil {
			return nil, nil, err
		}
	}
	return funcs, b, nil
}
