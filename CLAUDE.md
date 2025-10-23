# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

gotests is a Go commandline tool that automatically generates table-driven tests based on function and method signatures. It parses Go source files and generates test files with proper test scaffolding, including automatic import resolution.

## Build and Test Commands

```bash
# Install the tool
go get -u github.com/cweill/gotests/...

# Build the binary
go build -o gotests ./gotests

# Run all tests
go test -v ./...

# Run tests for a specific package
go test -v ./gotests/process

# Run tests with coverage (excluding certain packages)
export PKGS=$(go list ./... | grep -vE "(gotests/gotests|.*data|templates)" | tr -s '\n' ',' | sed 's/.{1}$//')
go test -v -covermode=count -coverpkg=$PKGS -coverprofile=coverage.cov

# Run a single test
go test -v -run TestSpecificTest ./path/to/package
```

Note: Tests may fail when run with the `-race` flag (see comments in codebase).

## Architecture

### Core Flow

1. **Input Processing** (`internal/input`): Resolves file paths from command-line arguments
2. **Parsing** (`internal/goparser`): Parses Go source files into AST and extracts function signatures
3. **Model Generation** (`internal/models`): Converts parsed data into domain models (Function, Receiver, Field, etc.)
4. **Template Rendering** (`internal/render`): Uses Go templates to generate test code
5. **Output Generation** (`internal/output`): Combines templates with parsed data and manages imports

### Key Components

**gotests.go**: Main library entry point
- `GenerateTests()`: Orchestrates the entire test generation pipeline
- `parallelize()`: Generates tests for multiple files concurrently using goroutines
- `testableFuncs()`: Filters functions based on options (exported, regex patterns, existing tests)

**internal/models/models.go**: Core data structures
- `Function`: Represents a function/method with parameters, results, and receiver
- `Receiver`: Represents method receivers with their fields
- `Field`: Represents function parameters, results, or receiver fields
- `Expression`: Type information including whether it's a pointer, variadic, or writer

**internal/goparser/goparser.go**: AST parsing
- Parses source files into `ast.File` objects
- Uses `go/types` package for type checking and resolution
- Extracts function signatures, receivers, parameters, and return types

**internal/render/render.go**: Template system
- Default templates in `internal/render/templates/`
- Custom templates supported via `-template` (named sets like "testify") or `-template_dir` (custom directory)
- Templates can be embedded as bindata or loaded from filesystem
- Template functions: `Field`, `Receiver`, `Param`, `Want`, `Got` for generating variable names

**gotests/process/process.go**: Command-line processing
- Handles file I/O and error reporting
- Converts CLI options to library options
- Manages writing output to stdout or files with `-w` flag

### Template System

The tool uses Go text templates to generate test code. There are two built-in template sets:
- `test`: Default table-driven tests
- `testify`: Tests using the testify assertion library

Templates can be customized via:
- `-template_dir`: Directory with custom templates (takes precedence)
- `-template`: Named template set from `templates/` directory
- Environment variables: `GOTESTS_TEMPLATE_DIR`, `GOTESTS_TEMPLATE`

### Function Filtering

Functions are filtered in this precedence order:
1. `-excl`: Exclude functions matching regex (highest precedence)
2. `-exported`: Only exported functions
3. `-only`: Only functions matching regex
4. `-all`: All functions (default is none)

Existing test functions are automatically excluded to avoid duplication.

## Development Notes

- The project uses Go modules (`go.mod`)
- Minimum supported Go version: 1.6
- The codebase uses concurrent test generation via goroutines for performance
- Tests are in `testdata/` directories with golden file comparisons in `testdata/goldens/`
- The `templates/` directory contains built-in template sets
- Bindata is used to embed templates in the binary (via `internal/render/bindata/`)
- Always use scripts/regenerate-goldens.sh to generate the goldens for tests.