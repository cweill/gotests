package input

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/cweill/gotests/internal/models"
)

var ErrNoFilesFound = errors.New("no files found")

// Returns all the Golang files for the given path. Ignores hidden files.
func Files(srcPath string) ([]models.Path, error) {
	var srcPaths []models.Path
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
			if isHiddenFile(p) {
				continue
			}
			src := models.Path(p)
			if !src.IsTestPath() {
				srcPaths = append(srcPaths, src)
			}
		}
		return srcPaths, nil
	}
	if filepath.Ext(srcPath) == ".go" && !isHiddenFile(srcPath) {
		src := models.Path(srcPath)
		if !src.IsTestPath() {
			srcPaths = append(srcPaths, src)
		}
		return srcPaths, nil
	}
	return nil, fmt.Errorf("no files found at %v", srcPath)
}

func isHiddenFile(path string) bool {
	return []rune(filepath.Base(path))[0] == '.'
}
