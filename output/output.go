package output

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/tools/imports"

	"github.com/cweill/gotests/models"
	"github.com/cweill/gotests/render"
)

const newFilePerm os.FileMode = 0644

type Options struct {
	Messagediff bool
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

func Write(dest string, b []byte) error {
	var isNewFile bool
	if IsFileExist(dest) {
		df, err := os.Create(dest)
		if err != nil {
			return fmt.Errorf("os.Create: %v", err)
		}
		defer df.Close()
		isNewFile = true
	}
	if err := ioutil.WriteFile(dest, b, newFilePerm); err != nil {
		if isNewFile {
			os.Remove(dest)
		}
		return err
	}
	return nil
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
		if err := render.TestFunction(w, fun, opt.Messagediff); err != nil {
			return fmt.Errorf("render.TestFunction: %v", err)
		}
	}
	if err := w.Flush(); err != nil {
		return fmt.Errorf("bufio.Flush: %v", err)
	}
	return nil
}
