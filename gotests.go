// Package gotests contains the core logic for generating table-driven tests.
package gotests

import (
	"fmt"
	"go/importer"
	"go/types"
	"path"
	"regexp"
	"sort"
	"strconv"
	"sync"

	"github.com/cweill/gotests/internal/gomod"
	"github.com/cweill/gotests/internal/goparser"
	"github.com/cweill/gotests/internal/input"
	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/output"
)

// Options provides custom filters and parameters for generating tests.
type Options struct {
	Only           *regexp.Regexp         // Includes only functions that match.
	Exclude        *regexp.Regexp         // Excludes functions that match.
	Exported       bool                   // Include only exported methods
	PackageTest    bool                   // Generate black-box tests in the <pkg>_test package. Implies Exported.
	PrintInputs    bool                   // Print function parameters in error messages
	Subtests       bool                   // Print tests using Go 1.7 subtests
	Parallel       bool                   // Print tests that runs the subtests in parallel.
	Named          bool                   // Create Map instead of slice
	Importer       func() types.Importer  // A custom importer.
	Template       string                 // Name of custom template set
	TemplateDir    string                 // Path to custom template set
	TemplateParams map[string]interface{} // Custom external parameters
	TemplateData   [][]byte               // Data slice for templates
	UseGoCmp       bool                   // Use cmp.Equal (google/go-cmp) instead of reflect.DeepEqual
	UseAI          bool                   // Generate test cases using AI
	AIModel        string                 // AI model to use
	AIEndpoint     string                 // AI API endpoint
	AIMinCases     int                    // Minimum number of test cases to generate
	AIMaxCases     int                    // Maximum number of test cases to generate
}

// A GeneratedTest contains information about a test file with generated tests.
type GeneratedTest struct {
	Path      string             // The test file's absolute path.
	Functions []*models.Function // The functions with new test methods.
	Output    []byte             // The contents of the test file. Nil when every function was skipped.
	Skipped   []SkippedFunction  // Selected functions that no test could be generated for.
}

// SkippedFunction is a function selected for testing that gotests could not
// generate a test for, such as one whose signature uses an unexported type
// when generating tests in an external _test package.
type SkippedFunction struct {
	Name   string // The function's name, qualified by its receiver type for methods.
	Reason string // Why no test was generated.
}

// GenerateTests generates table-driven tests for the function and method
// signatures defined in the target source path file(s). The source path
// parameter can be either a Go source file or directory containing Go files.
func GenerateTests(srcPath string, opt *Options) ([]*GeneratedTest, error) {
	if opt == nil {
		opt = &Options{}
	}
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
	return parallelize(srcFiles, files, opt)
}

// result stores a generateTest result.
type result struct {
	gt  *GeneratedTest
	err error
}

// parallelize generates tests for the given source files concurrently.
func parallelize(srcFiles, files []models.Path, opt *Options) ([]*GeneratedTest, error) {
	var wg sync.WaitGroup
	rs := make(chan *result, len(srcFiles))
	for _, src := range srcFiles {
		wg.Add(1)
		// Worker
		go func(src models.Path) {
			defer wg.Done()
			r := &result{}
			r.gt, r.err = generateTest(src, files, opt)
			rs <- r
		}(src)
	}
	// Closer.
	go func() {
		wg.Wait()
		close(rs)
	}()
	return readResults(rs)
}

// readResults reads the result channel.
func readResults(rs <-chan *result) ([]*GeneratedTest, error) {
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
	srcPkg := h.Package
	testPath := src.TestPath()
	h, tf, err := parseTestFile(p, testPath, h)
	if err != nil {
		return nil, err
	}
	// parseTestFile returns the source header when there is no test file to extend.
	extendsTestFile := h != sr.Header
	packageTest := opt.PackageTest
	switch {
	case h.Package == srcPkg+"_test":
		// Tests added to an existing external test file must be external too.
		packageTest = true
	case packageTest && extendsTestFile:
		return nil, fmt.Errorf("%s is in package %s; generating tests in package %s_test would break its existing tests", testPath, h.Package, srcPkg)
	}
	// An external test package can only reach exported identifiers.
	funcs := testableFuncs(sr.Funcs, opt.Only, opt.Exclude, opt.Exported || packageTest, tf)
	if len(funcs) == 0 {
		return nil, nil
	}
	qualifier := srcPkg
	var skipped []SkippedFunction
	if packageTest {
		importPath, err := gomod.GetFullImportPath(string(src))
		if err != nil {
			return nil, fmt.Errorf("resolving import path for package %s_test: %v", srcPkg, err)
		}
		h.Package = srcPkg + "_test"
		qualifier = importPackage(h, srcPkg, importPath)
		var testable []*models.Function
		for _, fn := range funcs {
			if reason := goparser.QualifyTypes(fn, qualifier, sr.TypeNames); reason != "" {
				skipped = append(skipped, SkippedFunction{Name: displayName(fn), Reason: reason})
				continue
			}
			testable = append(testable, fn)
		}
		funcs = testable
		if len(funcs) == 0 {
			return &GeneratedTest{Path: testPath, Skipped: skipped}, nil
		}
	}

	options := output.Options{
		PrintInputs:    opt.PrintInputs,
		Subtests:       opt.Subtests,
		PackageTest:    packageTest,
		PackageName:    qualifier,
		Parallel:       opt.Parallel,
		Named:          opt.Named,
		UseGoCmp:       opt.UseGoCmp,
		Template:       opt.Template,
		TemplateDir:    opt.TemplateDir,
		TemplateParams: opt.TemplateParams,
		TemplateData:   opt.TemplateData,
		UseAI:          opt.UseAI,
		AIModel:        opt.AIModel,
		AIEndpoint:     opt.AIEndpoint,
		AIMinCases:     opt.AIMinCases,
		AIMaxCases:     opt.AIMaxCases,
	}

	b, err := options.Process(h, funcs)
	if err != nil {
		return nil, fmt.Errorf("output.Process: %v", err)
	}
	return &GeneratedTest{
		Path:      testPath,
		Functions: funcs,
		Output:    b,
		Skipped:   skipped,
	}, nil
}

// displayName names fn as it appears in source, e.g. "Configure" or "server.Start".
func displayName(fn *models.Function) string {
	if fn.Receiver != nil {
		return fn.Receiver.Type.Value + "." + fn.Name
	}
	return fn.Name
}

// importPackage makes sure h imports the package under test and returns the
// name to qualify its identifiers with. An existing test file may already
// import it, possibly under an alias; importing it again would not compile.
func importPackage(h *models.Header, pkgName, importPath string) string {
	quoted := strconv.Quote(importPath)
	for _, imp := range h.Imports {
		if imp.Path != quoted {
			continue
		}
		switch imp.Name {
		case "":
			return pkgName
		case "_", ".":
			// Not usable as a qualifier; add a named import alongside it.
		default:
			return imp.Name
		}
	}
	h.Imports = append(h.Imports, &models.Import{Name: pkgName, Path: quoted})
	return pkgName
}

func parseTestFile(p *goparser.Parser, testPath string, h *models.Header) (*models.Header, []string, error) {
	if !output.IsFileExist(testPath) {
		return h, nil, nil
	}
	tr, err := p.Parse(testPath, nil)
	if err != nil {
		if err == goparser.ErrEmptyFile {
			// Overwrite empty test files.
			return h, nil, nil
		}
		return nil, nil, fmt.Errorf("Parser.Parse test file: %v", err)
	}
	var testFuncs []string
	for _, fun := range tr.Funcs {
		testFuncs = append(testFuncs, fun.Name)
	}
	tr.Header.Imports = append(tr.Header.Imports, h.Imports...)
	h = tr.Header
	return h, testFuncs, nil
}

func testableFuncs(funcs []*models.Function, only, excl *regexp.Regexp, exp bool, testFuncs []string) []*models.Function {
	sort.Strings(testFuncs)
	var fs []*models.Function
	for _, f := range funcs {
		if isTestFunction(f, testFuncs) || isExcluded(f, excl) || isUnexported(f, exp) || !isIncluded(f, only) || isInvalid(f) {
			continue
		}
		fs = append(fs, f)
	}
	return fs
}

func isInvalid(f *models.Function) bool {
	if f.Name == "init" && f.IsNaked() {
		return true
	}
	// Benchmark functions are not test candidates. Detect them by their
	// signature rather than their name so that helper functions taking a
	// *testing.B are also skipped.
	if f.IsBenchmark {
		return true
	}
	return false
}

func isTestFunction(f *models.Function, testFuncs []string) bool {
	return len(testFuncs) > 0 && contains(testFuncs, f.TestName())
}

func isExcluded(f *models.Function, excl *regexp.Regexp) bool {
	return excl != nil && (excl.MatchString(f.Name) || excl.MatchString(f.FullName()))
}

func isUnexported(f *models.Function, exp bool) bool {
	return exp && !f.IsExported
}

func isIncluded(f *models.Function, only *regexp.Regexp) bool {
	return only == nil || only.MatchString(f.Name) || only.MatchString(f.FullName())
}

func contains(ss []string, s string) bool {
	if i := sort.SearchStrings(ss, s); i < len(ss) && ss[i] == s {
		return true
	}
	return false
}
