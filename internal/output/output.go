package output

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/tools/imports"

	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/render"
)

type Options struct {
	PrintInputs bool
	Subtests    bool
	TemplateDir string
	Benchmark   bool
}

func Process(head *models.Header, funcs []*models.Function, opt *Options) ([]byte, error) {
	if opt != nil && opt.TemplateDir != "" {
		err := render.LoadCustomTemplates(opt.TemplateDir)
		if err != nil {
			return nil, fmt.Errorf("loading custom templates: %v", err)
		}
	}

	tf, err := ioutil.TempFile("", "gotests_")
	if err != nil {
		return nil, fmt.Errorf("ioutil.TempFile: %v", err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())
	b := &bytes.Buffer{}
	if err := writeAll(b, head, funcs, opt); err != nil {
		return nil, err
	}
	/*	if err := writeTests(b, head, funcs, opt); err != nil {
			return nil, err
		}
		if err := writeBenchmarks(b, head, funcs, opt); err != nil {
			return nil, err
		}*/
	out, err := imports.Process(tf.Name(), b.Bytes(), nil)
	fmt.Println("###", string(b.Bytes()))
	if err != nil {
		return nil, fmt.Errorf("imports.Process: %v", err)
	}
	return out, nil
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// writeAll write tests, benchmarks, examples to output
// but now support tests and benchmarks
func writeAll(w io.Writer, head *models.Header, funcs []*models.Function, opt *Options) error {
	b := bufio.NewWriter(w)
	if err := render.Header(b, head); err != nil {
		return fmt.Errorf("render.Header: %v", err)
	}
	for _, fun := range funcs {
		if err := render.TestFunction(b, fun, opt.PrintInputs, opt.Subtests); err != nil {
			return fmt.Errorf("render.TestFunction: %v", err)
		}
		if opt.Benchmark {
			if err := render.BenchmarkFunction(b, fun, opt.PrintInputs, opt.Subtests); err != nil {
				return fmt.Errorf("render.BenchmarkFunction: %v", err)
			}
		}
	}
	return b.Flush()
}

func writeTests(w io.Writer, head *models.Header, funcs []*models.Function, opt *Options) error {
	b := bufio.NewWriter(w)
	if err := render.Header(b, head); err != nil {
		return fmt.Errorf("render.Header: %v", err)
	}
	for _, fun := range funcs {
		if err := render.TestFunction(b, fun, opt.PrintInputs, opt.Subtests); err != nil {
			return fmt.Errorf("render.TestFunction: %v", err)
		}
	}
	return b.Flush()
}

func writeBenchmarks(w io.Writer, head *models.Header, funcs []*models.Function, opt *Options) error {
	b := bufio.NewWriter(w)
	fmt.Println("start render header")
	/*	if err := render.Header(b, head); err != nil {
		return fmt.Errorf("render.Header: %v", err)
	}*/
	fmt.Println("start render func")
	for _, fun := range funcs {
		if err := render.BenchmarkFunction(b, fun, opt.PrintInputs, opt.Subtests); err != nil {
			return fmt.Errorf("render.BenchmarkFunction: %v", err)
		}
	}
	return b.Flush()
}
