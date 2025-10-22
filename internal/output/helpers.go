package output

import "os"

// IsFileExist checks whether a file exists at the given path.
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
