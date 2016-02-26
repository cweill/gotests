package gotests

import (
	"fmt"
	"go/importer"
	"go/types"
	"path"
	"regexp"
	"sync"

	"github.com/cweill/gotests/internal/goparser"
	"github.com/cweill/gotests/internal/input"
	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/output"
)

type Options struct {
	Only        *regexp.Regexp
	Exclude     *regexp.Regexp
	Exported    bool
	PrintInputs bool
	Importer    func() types.Importer
}

type GeneratedTest struct {
	Path      string             // The test file's absolute path.
	Functions []*models.Function // The functions with new test methods.
	Output    []byte             // The contents of the test file.
}

func GenerateTests(srcPath string, opt *Options) ([]*GeneratedTest, error) {
	srcFiles, err := input.Files(srcPath)
	if err != nil {
		return nil, fmt.Errorf("input.Files: %v", err)
	}
	files, err := input.Files(path.Dir(srcPath))
	if err != nil {
		return nil, fmt.Errorf("input.Files: %v", err)
	}
	if opt.Importer == nil || opt.Importer() == nil {
		opt.Importer = importer.Default
	}
	type result struct {
		gt  *GeneratedTest
		err error
	}
	var wg sync.WaitGroup
	rs := make(chan *result, len(srcFiles))
	for _, src := range srcFiles {
		wg.Add(1)
		// Worker.
		go func(s models.Path) {
			defer wg.Done()
			r := &result{}
			r.gt, r.err = generateTest(s, files, opt)
			rs <- r
		}(src)
	}
	// Closer.
	go func() {
		wg.Wait()
		close(rs)
	}()
	var gts []*GeneratedTest
	for r := range rs {
		if r.err != nil {
			return nil, r.err
		}
		if r.gt != nil {
			gts = append(gts, r.gt)
		}
	}
	return gts, nil
}

func generateTest(src models.Path, files []models.Path, opt *Options) (*GeneratedTest, error) {
	p := &goparser.Parser{Importer: opt.Importer()}
	srcInfo, err := p.Parse(string(src), files)
	if err != nil {
		return nil, fmt.Errorf("Parser.Parse: %v", err)
	}
	header := srcInfo.Header
	var testFuncs []string
	testPath := models.Path(src).TestPath()
	if output.IsFileExist(testPath) {
		testInfo, err := p.Parse(testPath, nil)
		if err != nil {
			return nil, fmt.Errorf("Parser.Parse: %v", err)
		}
		for _, fun := range testInfo.Funcs {
			testFuncs = append(testFuncs, fun.Name)
		}
		h, err := goparser.ParseHeader(string(src), testPath)
		if err != nil {
			return nil, fmt.Errorf("goparser.ParseHeader: %v", err)
		}
		header = h
	}
	funcs := srcInfo.TestableFuncs(opt.Only, opt.Exclude, opt.Exported, testFuncs)
	if len(funcs) == 0 {
		return nil, nil
	}
	b, err := output.Process(header, funcs, &output.Options{
		PrintInputs: opt.PrintInputs,
	})
	if err != nil {
		return nil, fmt.Errorf("output.Process: %v", err)
	}
	return &GeneratedTest{
		Path:      testPath,
		Functions: funcs,
		Output:    b,
	}, nil
}
