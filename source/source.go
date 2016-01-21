package source

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cweill/gotests/models"
)

var NoFilesFound = errors.New("no files found")

func Files(srcPath string) ([]*models.FileInfo, error) {
	var srcPaths []*models.FileInfo
	srcPath, err := filepath.Abs(srcPath)
	if err != nil {
		return nil, fmt.Errorf("filepath.Abs: %v\n", err)
	}
	if filepath.Ext(srcPath) == "" {
		ps, err := filepath.Glob(srcPath + "/*.go")
		if err != nil {
			return nil, fmt.Errorf("filepath.Glob: %v\n", err)
		}
		for _, p := range ps {
			if !isTestPath(p) {
				srcPaths = append(srcPaths, &models.FileInfo{
					SourcePath: p,
					TestPath:   testPath(p),
				})
			}
		}
		return srcPaths, nil
	}
	if filepath.Ext(srcPath) == ".go" {
		if !isTestPath(srcPath) {
			srcPaths = append(srcPaths, &models.FileInfo{
				SourcePath: srcPath,
				TestPath:   testPath(srcPath),
			})
		}
		return srcPaths, nil
	}
	return nil, NoFilesFound
}

func testPath(srcPath string) string {
	return strings.TrimSuffix(srcPath, ".go") + "_test.go"
}

func isTestPath(path string) bool {
	return strings.HasSuffix(path, "_test.go")
}
