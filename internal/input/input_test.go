package input

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cweill/gotests/internal/models"
)

func TestFiles_RecursivePattern(t *testing.T) {
	// Create temporary directory structure for testing
	tmpDir := t.TempDir()

	// Create subdirectories
	pkgDir := filepath.Join(tmpDir, "pkg")
	subDir := filepath.Join(pkgDir, "subpkg")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create test files
	files := map[string]string{
		filepath.Join(pkgDir, "main.go"):        "package pkg\nfunc Foo() {}",
		filepath.Join(pkgDir, "main_test.go"):   "package pkg\nfunc TestFoo(t *testing.T) {}",
		filepath.Join(subDir, "sub.go"):         "package subpkg\nfunc Bar() {}",
		filepath.Join(subDir, "sub_test.go"):    "package subpkg\nfunc TestBar(t *testing.T) {}",
		filepath.Join(subDir, ".hidden.go"):     "package subpkg\nfunc Hidden() {}",
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Test recursive pattern
	pattern := filepath.Join(pkgDir, "...")
	got, err := Files(pattern)
	if err != nil {
		t.Fatalf("Files() error = %v", err)
	}

	// Should find main.go and sub.go, but not test files or hidden files
	want := 2
	if len(got) != want {
		t.Errorf("Files() returned %d files, want %d", len(got), want)
		for i, f := range got {
			t.Logf("  [%d] %s", i, f)
		}
	}

	// Verify we got the right files
	foundMain := false
	foundSub := false
	for _, file := range got {
		base := filepath.Base(string(file))
		if base == "main.go" {
			foundMain = true
		}
		if base == "sub.go" {
			foundSub = true
		}
		// Make sure we didn't get test files or hidden files
		if base == "main_test.go" || base == "sub_test.go" || base == ".hidden.go" {
			t.Errorf("Files() should not return %s", base)
		}
	}

	if !foundMain {
		t.Error("Files() did not find main.go")
	}
	if !foundSub {
		t.Error("Files() did not find sub.go")
	}
}

func TestFiles_SingleFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a single Go file
	goFile := filepath.Join(tmpDir, "test.go")
	if err := os.WriteFile(goFile, []byte("package test\nfunc Test() {}"), 0644); err != nil {
		t.Fatal(err)
	}

	got, err := Files(goFile)
	if err != nil {
		t.Fatalf("Files() error = %v", err)
	}

	if len(got) != 1 {
		t.Errorf("Files() returned %d files, want 1", len(got))
	}

	if len(got) > 0 && string(got[0]) != goFile {
		t.Errorf("Files() = %v, want %v", got[0], goFile)
	}
}

func TestFiles_Directory(t *testing.T) {
	tmpDir := t.TempDir()

	// Create files in directory
	files := map[string]string{
		filepath.Join(tmpDir, "a.go"):        "package test",
		filepath.Join(tmpDir, "b.go"):        "package test",
		filepath.Join(tmpDir, "a_test.go"):   "package test",
		filepath.Join(tmpDir, ".hidden.go"):  "package test",
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	got, err := Files(tmpDir)
	if err != nil {
		t.Fatalf("Files() error = %v", err)
	}

	// Should find a.go and b.go, but not test file or hidden file
	want := 2
	if len(got) != want {
		t.Errorf("Files() returned %d files, want %d", len(got), want)
	}
}

func TestFiles_NonGoFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a non-Go file
	txtFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(txtFile, []byte("not go"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Files(txtFile)
	if err == nil {
		t.Error("Files() should return error for non-Go file")
	}
}

func Test_recursiveDirFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create nested directory structure
	level1 := filepath.Join(tmpDir, "level1")
	level2 := filepath.Join(level1, "level2")
	hiddenDir := filepath.Join(level1, ".hidden")

	if err := os.MkdirAll(level2, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(hiddenDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create test files
	files := map[string]string{
		filepath.Join(tmpDir, "root.go"):         "package root",
		filepath.Join(level1, "l1.go"):           "package level1",
		filepath.Join(level1, "l1_test.go"):      "package level1",
		filepath.Join(level2, "l2.go"):           "package level2",
		filepath.Join(hiddenDir, "hidden.go"):    "package hidden",
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	got, err := recursiveDirFiles(tmpDir)
	if err != nil {
		t.Fatalf("recursiveDirFiles() error = %v", err)
	}

	// Should find root.go, l1.go, and l2.go (3 files)
	// Should NOT find l1_test.go (test file) or hidden.go (in hidden dir)
	want := 3
	if len(got) != want {
		t.Errorf("recursiveDirFiles() returned %d files, want %d", len(got), want)
		for i, f := range got {
			t.Logf("  [%d] %s", i, f)
		}
	}

	// Verify we didn't get test files or files in hidden directories
	for _, file := range got {
		if models.Path(string(file)).IsTestPath() {
			t.Errorf("recursiveDirFiles() should not return test file: %s", file)
		}
		if filepath.Base(filepath.Dir(string(file))) == ".hidden" {
			t.Errorf("recursiveDirFiles() should not return files in hidden directory: %s", file)
		}
	}
}

func Test_isHiddenFile(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "hidden file",
			path: ".hidden.go",
			want: true,
		},
		{
			name: "hidden file with path",
			path: "/path/to/.hidden.go",
			want: true,
		},
		{
			name: "normal file",
			path: "normal.go",
			want: false,
		},
		{
			name: "normal file with path",
			path: "/path/to/normal.go",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHiddenFile(tt.path); got != tt.want {
				t.Errorf("isHiddenFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
