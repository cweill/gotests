package checker

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"

	"golang.org/x/tools/imports"
)

type Import string

const MessageDiff Import = `package check

func F(i1, i2 interface{}) (string, bool) {
	return messagediff.DeepDiff(i1, i2)
}
`

func CheckImport(imp Import) (bool, error) {
	tf, err := ioutil.TempFile("", "gotests_checker_")
	if err != nil {
		return false, fmt.Errorf("ioutil.TempFile: %v", err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())
	b := []byte(imp)
	if _, err = tf.Write(b); err != nil {
		return false, fmt.Errorf("f.Write: %v", err)
	}
	tf.Sync()
	out, err := imports.Process(tf.Name(), b, nil)
	if err != nil {
		return false, fmt.Errorf("imports.Process: %v", err)
	}
	if err = ioutil.WriteFile(tf.Name(), out, 0644); err != nil {
		return false, fmt.Errorf("ioutil.WriteFile: %v", err)
	}
	f, err := parser.ParseFile(token.NewFileSet(), tf.Name(), nil, 0)
	if err != nil {
		return false, fmt.Errorf("parser.ParseFile: %v", err)
	}
	return len(f.Imports) == 1, nil
}
