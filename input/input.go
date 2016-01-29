package input

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/cweill/gotests/models"
)

var NoFilesFound = errors.New("no files found")

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
			src := models.Path(p)
			if !src.IsTestPath() {
				srcPaths = append(srcPaths, src)
			}
		}
		return srcPaths, nil
	}
	if filepath.Ext(srcPath) == ".go" {
		src := models.Path(srcPath)
		if !src.IsTestPath() {
			srcPaths = append(srcPaths, src)
		}
		return srcPaths, nil
	}
	return nil, NoFilesFound
}
