package source

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cweill/gotests/models"
)

func Files(srcPath string) []*models.FileInfo {
	var srcPaths []*models.FileInfo
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
			if !isTestPath(p) {
				srcPaths = append(srcPaths, &models.FileInfo{
					SourcePath: p,
					TestPath:   testPath(p),
				})
			}
		}
	} else if filepath.Ext(srcPath) == ".go" {
		if !isTestPath(srcPath) {
			srcPaths = append(srcPaths, &models.FileInfo{
				SourcePath: srcPath,
				TestPath:   testPath(srcPath),
			})
		}
	}
	return srcPaths
}

func testPath(srcPath string) string {
	return strings.TrimSuffix(srcPath, ".go") + "_test.go"
}

func isTestPath(path string) bool {
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
