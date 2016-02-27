package gotests

import (
	"fmt"
	"go/importer"
	"go/types"
	"path"
	"regexp"
	"sort"
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
	sr, err := p.Parse(string(src), files)
	if err != nil {
		return nil, fmt.Errorf("Parser.Parse source file: %v", err)
	}
	h := sr.Header
	h.Code = nil // Code is only needed from parsed test files.
	var testFuncs []string
	testPath := models.Path(src).TestPath()
	if output.IsFileExist(testPath) {
		tr, err := p.Parse(testPath, nil)
		if err != nil {
			return nil, fmt.Errorf("Parser.Parse test file: %v", err)
		}
		for _, fun := range tr.Funcs {
			testFuncs = append(testFuncs, fun.Name)
		}
		tr.Header.Imports = append(tr.Header.Imports, h.Imports...)
		h = tr.Header
	}
	funcs := testableFuncs(sr.Funcs, opt.Only, opt.Exclude, opt.Exported, testFuncs)
	if len(funcs) == 0 {
		return nil, nil
	}
	b, err := output.Process(h, funcs, &output.Options{
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

func testableFuncs(funcs []*models.Function, only, excl *regexp.Regexp, exp bool, testFuncs []string) []*models.Function {
	sort.Strings(testFuncs)
	var fs []*models.Function
	for _, f := range funcs {
		if f.Receiver == nil && len(f.Parameters) == 0 && len(f.Results) == 0 {
			continue
		}
		if len(testFuncs) > 0 && contains(testFuncs, f.TestName()) {
			continue
		}
		if excl != nil && (excl.MatchString(f.Name) || excl.MatchString(f.FullName())) {
			continue
		}
		if exp && !f.IsExported {
			continue
		}
		if only != nil && !only.MatchString(f.Name) && !only.MatchString(f.FullName()) {
			continue
		}
		fs = append(fs, f)
	}
	return fs
}

func contains(ss []string, s string) bool {
	if i := sort.SearchStrings(ss, s); i < len(ss) && ss[i] == s {
		return true
	}
	return false
}
