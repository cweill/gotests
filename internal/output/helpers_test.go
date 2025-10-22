package output

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsFileExist(t *testing.T) {
	tests := []struct {
		name     string
		setupFn  func(t *testing.T) string
		want     bool
	}{
		{
			name: "existing file",
			setupFn: func(t *testing.T) string {
				tmpDir := t.TempDir()
				filePath := filepath.Join(tmpDir, "test.go")
				if err := os.WriteFile(filePath, []byte("package test"), 0644); err != nil {
					t.Fatal(err)
				}
				return filePath
			},
			want: true,
		},
		{
			name: "non-existing file",
			setupFn: func(t *testing.T) string {
				tmpDir := t.TempDir()
				return filepath.Join(tmpDir, "nonexistent.go")
			},
			want: false,
		},
		{
			name: "existing directory",
			setupFn: func(t *testing.T) string {
				tmpDir := t.TempDir()
				dirPath := filepath.Join(tmpDir, "mydir")
				if err := os.Mkdir(dirPath, 0755); err != nil {
					t.Fatal(err)
				}
				return dirPath
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setupFn(t)
			got := IsFileExist(path)
			if got != tt.want {
				t.Errorf("IsFileExist(%q) = %v, want %v", path, got, tt.want)
			}
		})
	}
}
