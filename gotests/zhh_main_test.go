package main

import (
	"os"
	"testing"

	"github.com/cweill/gotests/gotests/process"
)

func Test_main(t *testing.T) {
	process.Run(os.Stdout, []string{
		"./mock_need_test.go",
	}, &process.Options{
		AllFuncs: true,
	})
}


