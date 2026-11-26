package gomod_test

import (
	"fmt"
	"log"

	"github.com/cweill/gotests/internal/gomod"
)

func ExampleGetFullImportPath_file() {
	// Get import path for a specific Go file
	importPath, err := gomod.GetFullImportPath("gomod.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(importPath)
	// Output: github.com/cweill/gotests/internal/gomod
}

func ExampleGetFullImportPath_directory() {
	// Get import path for a package directory
	importPath, err := gomod.GetFullImportPath(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(importPath)
	// Output: github.com/cweill/gotests/internal/gomod
}

func ExampleGetFullImportPath_moduleRoot() {
	// Get import path for the module root directory
	importPath, err := gomod.GetFullImportPath("../..")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(importPath)
	// Output: github.com/cweill/gotests
}
