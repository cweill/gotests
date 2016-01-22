package output

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cweill/gotests/models"
	"github.com/cweill/gotests/render"
	"golang.org/x/tools/imports"
)

const newFilePerm os.FileMode = 0644

func Write(srcPath, destPath string, head *models.Header, funcs []*models.Function) ([]string, error) {
	tf, err := ioutil.TempFile("", "gotests_")
	if err != nil {
		return nil, fmt.Errorf("ioutil.TempFile: %v", err)
	}
	defer os.Remove(tf.Name())
	tests, err := writeTestsToTemp(tf, head, funcs)
	if err != nil {
		return nil, err
	}
	var isDestNew bool
	if IsFileExist(destPath) {
		df, err := os.Create(destPath)
		if err != nil {
			return nil, fmt.Errorf("os.Create: %v", err)
		}
		defer df.Close()
		destPath = df.Name()
		isDestNew = true
	}
	if err := copyTempToDest(tf.Name(), destPath); err != nil {
		if isDestNew {
			os.Remove(destPath)
		}
		return nil, err
	}
	return tests, nil
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func writeTestsToTemp(temp *os.File, head *models.Header, funcs []*models.Function) ([]string, error) {
	w := bufio.NewWriter(temp)
	if err := render.Header(w, head); err != nil {
		return nil, fmt.Errorf("render.Header: %v", err)
	}
	var tests []string
	for _, fun := range funcs {
		if err := render.TestFunction(w, fun); err != nil {
			return nil, fmt.Errorf("render.TestFunction: %v", err)
		}
		tests = append(tests, fmt.Sprintf("%v.%v", head.Package, fun.TestName()))
	}
	if err := w.Flush(); err != nil {
		return nil, fmt.Errorf("bufio.Flush: %v", err)
	}
	if err := processImports(temp.Name()); err != nil {
		return nil, err
	}
	return tests, nil
}

func processImports(path string) error {
	v, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	b, err := imports.Process(path, v, nil)
	if err != nil {
		return fmt.Errorf("imports.Process: %v", err)
	}
	return ioutil.WriteFile(path, b, 0)
}

func copyTempToDest(tempPath, destPath string) error {
	b, err := ioutil.ReadFile(tempPath)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	return ioutil.WriteFile(destPath, b, newFilePerm)
}
