package gotests

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"testing"
)

// TestPackageTestOutputCompiles checks that tests generated for an external
// _test package build against the package under test, which golden files
// alone cannot show.
func TestPackageTestOutputCompiles(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the go command")
	}
	goBin, err := exec.LookPath("go")
	if err != nil {
		t.Skip("go command not found")
	}
	const src = "testdata/packagetest/shapes/shapes.go"
	gts, err := GenerateTests(src, &Options{PackageTest: true, Subtests: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(gts) != 1 {
		t.Fatalf("GenerateTests(%s) returned %d tests, want 1", src, len(gts))
	}

	// Recreate the package at the import path the generated test expects:
	// github.com/cweill/gotests/testdata/packagetest/shapes.
	mod := t.TempDir()
	pkg := filepath.Join(mod, "shapes")
	b, err := os.ReadFile(src)
	if err != nil {
		t.Fatal(err)
	}
	files := map[string][]byte{
		filepath.Join(mod, "go.mod"):         []byte("module github.com/cweill/gotests/testdata/packagetest\n\ngo 1.24\n"),
		filepath.Join(pkg, "shapes.go"):      b,
		filepath.Join(pkg, "shapes_test.go"): gts[0].Output,
	}
	if err := os.Mkdir(pkg, 0o755); err != nil {
		t.Fatal(err)
	}
	for name, data := range files {
		if err := os.WriteFile(name, data, 0o644); err != nil {
			t.Fatal(err)
		}
	}

	cmd := exec.Command(goBin, "test", "-count=1", "-run", "^$", "./shapes")
	cmd.Dir = mod
	cmd.Env = append(os.Environ(), "GOWORK=off", "GOFLAGS=-mod=mod", "GOPROXY=off")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated _test package does not compile: %v\n%s\n--- shapes_test.go:\n%s", err, out, gts[0].Output)
	}
}

func TestPackageTestReportsSkipped(t *testing.T) {
	const src = "testdata/packagetest/shapes/shapes.go"
	var mu sync.Mutex
	var skipped []SkippedFunction
	onSkip := func(s SkippedFunction) {
		mu.Lock()
		defer mu.Unlock()
		skipped = append(skipped, s)
	}

	gts, err := GenerateTests(src, &Options{PackageTest: true, OnSkip: onSkip})
	if err != nil {
		t.Fatal(err)
	}
	if len(gts) != 1 {
		t.Fatalf("GenerateTests() returned %d tests, want 1", len(gts))
	}
	abs, err := filepath.Abs(src)
	if err != nil {
		t.Fatal(err)
	}
	want := []SkippedFunction{
		{Path: abs, Name: "Configure", Reason: "uses unexported type config"},
		{Path: abs, Name: "server.Start", Reason: "receiver type server is unexported"},
	}
	if !reflect.DeepEqual(skipped, want) {
		t.Errorf("skipped = %+v, want %+v", skipped, want)
	}

	// When every selected function is skipped, no result is returned, so
	// callers that write each result's Output never empty a test file.
	skipped = nil
	gts, err = GenerateTests(src, &Options{PackageTest: true, Only: regexp.MustCompile(`^Configure$`), OnSkip: onSkip})
	if err != nil {
		t.Fatal(err)
	}
	if len(gts) != 0 {
		t.Errorf("all functions skipped: GenerateTests() returned %d results, want 0", len(gts))
	}
	if len(skipped) != 1 {
		t.Errorf("all functions skipped: OnSkip called %d times, want 1", len(skipped))
	}
}
