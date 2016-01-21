package source

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Files(srcPath string) []string {
	var srcPaths []string
	srcPath, err := filepath.Abs(srcPath)
	if err != nil {
		fmt.Printf("filepath.Abs: %v\n", err)
		return nil
	}
	if filepath.Ext(srcPath) == "" {
		ps, err := filepath.Glob(srcPath + "/*.go")
		if err != nil {
			fmt.Printf("filepath.Glob: %v\n", err)
			return nil
		}
		for _, p := range ps {
			if !isTestFile(p) {
				srcPaths = append(srcPaths, p)
			}
		}
	} else if filepath.Ext(srcPath) == ".go" {
		if !isTestFile(srcPath) {
			srcPaths = append(srcPaths, srcPath)
		}
	}
	return srcPaths
}

func TestPath(srcPath string) string {
	srcPath = strings.TrimSuffix(srcPath, ".go")
	return srcPath + "_test.go"
}

func isTestFile(path string) bool {
	ok, err := filepath.Match("*_test.go", path)
	if err != nil {
		fmt.Printf("filepath.Match: %v\n", err)
		return false
	}
	if ok {
		return true
	}
	return false
}
