package output

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/tools/imports"

	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/render"
)

type Options struct {
	PrintInputs bool
}

func Process(head *models.Header, funcs []*models.Function, opt *Options) ([]byte, error) {
	tf, err := ioutil.TempFile("", "gotests_")
	if err != nil {
		return nil, fmt.Errorf("ioutil.TempFile: %v", err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())
	if err := writeTestsToTemp(tf, head, funcs, opt); err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(tf.Name())
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	out, err := imports.Process(tf.Name(), b, nil)
	if err != nil {
		return nil, fmt.Errorf("imports.Process: %v", err)
	}
	return out, nil
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func writeTestsToTemp(temp *os.File, head *models.Header, funcs []*models.Function, opt *Options) error {
	w := bufio.NewWriter(temp)
	if err := render.Header(w, head); err != nil {
		return fmt.Errorf("render.Header: %v", err)
	}
	for _, fun := range funcs {
		if err := render.TestFunction(w, fun, opt.PrintInputs); err != nil {
			return fmt.Errorf("render.TestFunction: %v", err)
		}
	}
	return w.Flush()
}
